package data

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"errors"
	"fmt"
	"greenlight/internal/assert"
	"greenlight/internal/validator"
	"time"

	"golang.org/x/crypto/bcrypt"
)

var ErrDuplicatedEmail = errors.New("duplicated email")
var duplicatedEmailMessage = `pq: duplicate key value violates unique constraint "users_email_key"`

var AnonymousUser = &User{}

type UserModel struct {
	DB *sql.DB
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

func (m UserModel) Insert(user *User) error {
	query := `insert into users(name,email,password_hash,status)
	values($1,$2,$3,$4)
	returning id,created_at,updated_at,version`

	args := []any{user.Name, user.Email, user.Password.hash, user.Status}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err := m.DB.QueryRowContext(ctx, query, args...).Scan(
		&user.ID,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.Version,
	); err != nil {
		switch {
		case err.Error() == duplicatedEmailMessage:
			return ErrDuplicatedEmail
		default:
			return fmt.Errorf("insert user error: %w", err)
		}
	}

	return nil
}

func (m UserModel) GetByEmail(email string) (*User, error) {
	assert.Assert(email != "")

	query := `select id,name,email,password_hash,status,version,created_at,updated_at
	from users
	where email = $1`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var user User

	if err := m.DB.QueryRowContext(ctx, query, email).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password.hash,
		&user.Status,
		&user.Version,
		&user.CreatedAt,
		&user.UpdatedAt,
	); err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &user, nil
}

func (m UserModel) Update(user *User) error {
	assert.Assert(user != nil)

	query := `update users
	set name=$1, email=$2,password_hash=$3,status=$4,version=version+1,updated_at=now()
	where id=$5 and version=$6
	returning version`

	args := []any{
		user.Name,
		user.Email,
		user.Password.hash,
		user.Status,
		user.ID,
		user.Version,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err := m.DB.QueryRowContext(ctx, query, args...).Scan(&user.Version); err != nil {
		switch {
		case err.Error() == duplicatedEmailMessage:
			return ErrDuplicatedEmail
		case errors.Is(err, sql.ErrNoRows):
			return ErrStaleObject
		default:
			return fmt.Errorf("insert user error: %w", err)
		}
	}

	return nil
}

func (m UserModel) GetByToken(tokenScope, plaintextToken string) (*User, error) {
	tokenHash := sha256.Sum256([]byte(plaintextToken))
	query := `select u.id,u.created_at,u.name,u.email,u.password_hash,u.status,u.version
	from users u
	inner join tokens t
	on u.id = t.user_id
	where t.hash = $1 -- <- this is vulnerable to a timing attack
	and t.scope = $2
	and t.expire_at > $3` // q: what's difference between using sql now() and go time.Now()?

	args := []any{tokenHash[:], tokenScope, time.Now()}

	var user User

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err := m.DB.QueryRowContext(ctx, query, args...).Scan(
		&user.ID,
		&user.CreatedAt,
		&user.Name,
		&user.Email,
		&user.Password.hash,
		&user.Status,
		&user.Version,
	); err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &user, nil
}

func (u *User) IsAnonymous() bool {
	return u == AnonymousUser
}
