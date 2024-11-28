package data

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"errors"
	"fmt"
	"greenlight/internal/assert"
	"greenlight/internal/sqlcdb"
	"greenlight/internal/validator"
	"time"

	"golang.org/x/crypto/bcrypt"
)

var ErrDuplicatedEmail = errors.New("duplicated email")
var duplicatedEmailMessage = `pq: duplicate key value violates unique constraint "users_email_key"`

var AnonymousUser = &User{}

type UserModel struct {
	queries *sqlcdb.Queries
}

type User struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  password  `json:"-"`
	Status    string    `json:"status"`
	Version   int64     `json:"version"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type password struct {
	plaintext string
	hash      []byte
}

func (p *password) Set(plaintextPassword string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(plaintextPassword), 12)
	if err != nil {
		return err
	}

	p.plaintext = plaintextPassword
	p.hash = hash

	return nil
}

func (p *password) Matches(plaintextPassword string) (bool, error) {
	if err := bcrypt.CompareHashAndPassword(p.hash, []byte(plaintextPassword)); err != nil {
		switch {
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
			return false, nil
		default:
			return false, err
		}
	}

	return true, nil
}

func ValidateUser(v *validator.Validator, user *User) {
	assert.Assert(v != nil)
	assert.Assert(user != nil)

	v.Check(user.Name != "", "name", "must present")
	v.Check(len(user.Name) < 500, "name", "length must <500")
	ValidateEmail(v, user.Email)

	if user.Password.plaintext != "" {
		ValidatePasswordPlaintext(v, user.Password.plaintext)
	}

	assert.Assert(user.Password.hash != nil)
}

func ValidateEmail(v *validator.Validator, email string) {
	assert.Assert(v != nil)

	v.Check(email != "", "email", "must present")
	v.Check(validator.Match(email, validator.EmailRegexp), "email", "must be a valid email address")
}

func ValidatePasswordPlaintext(v *validator.Validator, password string) {
	assert.Assert(v != nil)

	v.Check(password != "", "password", "must present")
	v.Check(len(password) >= 8, "password", "length must >= 8")
	v.Check(len(password) <= 72, "password", "length must <= 72")
}

func (m UserModel) Create(ctx context.Context, user *User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	row, err := m.queries.CreateUser(ctx, &sqlcdb.CreateUserParams{
		Name:         user.Name,
		Email:        user.Email,
		PasswordHash: user.Password.hash,
		Status:       user.Status,
	})

	if err != nil {
		switch {
		case err.Error() == duplicatedEmailMessage:
			return ErrDuplicatedEmail
		default:
			return fmt.Errorf("insert user error: %w", err)
		}
	}

	user.ID = row.ID
	user.CreatedAt = row.CreatedAt
	user.UpdatedAt = row.UpdatedAt
	user.Version = int64(row.Version)

	return nil
}

func (m UserModel) GetByEmail(ctx context.Context, email string) (*User, error) {
	assert.Assert(email != "")

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	row, err := m.queries.GetUserByEmail(ctx, email)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	user := User{
		ID:    row.ID,
		Name:  row.Name,
		Email: row.Email,
		Password: password{
			hash: row.PasswordHash,
		},
		Status:    row.Status,
		Version:   int64(row.Version),
		CreatedAt: row.CreatedAt,
		UpdatedAt: row.UpdatedAt,
	}

	return &user, nil
}

func (m UserModel) Update(ctx context.Context, user *User) error {
	assert.Assert(user != nil)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	newVersion, err := m.queries.UpdateUser(ctx, &sqlcdb.UpdateUserParams{
		Name:         user.Name,
		Email:        user.Email,
		PasswordHash: user.Password.hash,
		Status:       user.Status,
		ID:           user.ID,
		Version:      int32(user.Version),
	})
	if err != nil {
		switch {
		case err.Error() == duplicatedEmailMessage:
			return ErrDuplicatedEmail
		case errors.Is(err, sql.ErrNoRows):
			return ErrStaleObject
		default:
			return fmt.Errorf("insert user error: %w", err)
		}
	}

	user.Version = int64(newVersion)

	return nil
}

func (m UserModel) GetByToken(ctx context.Context, tokenScope, plaintextToken string) (*User, error) {
	tokenHash := sha256.Sum256([]byte(plaintextToken))

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	row, err := m.queries.GetUserByToken(ctx, &sqlcdb.GetUserByTokenParams{
		Hash:  tokenHash[:],
		Scope: tokenScope,
	})
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	user := User{
		ID:    row.ID,
		Name:  row.Name,
		Email: row.Email,
		Password: password{
			hash: row.PasswordHash,
		},
		Version:   int64(row.Version),
		CreatedAt: row.CreatedAt,
		Status:    row.Status,
	}

	return &user, nil
}

func (u *User) IsAnonymous() bool {
	return u == AnonymousUser
}

func (u *User) Activated() bool {
	return u.Status == "activated"
}
