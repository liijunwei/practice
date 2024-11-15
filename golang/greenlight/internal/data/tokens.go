package data

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"database/sql"
	"encoding/base32"
	"greenlight/internal/common"
	"greenlight/internal/validator"
	"time"
)

const ScopeActivation = "activation"

type Token struct {
	Plaintext string
	Hash      []byte
	UserId    int64
	Scope     string
	ExpireAt  time.Time
}

type TokenModel struct {
	DB *sql.DB
}

func generateToken(userID int64, ttl time.Duration, scope string) (*Token, error) {
	randomBytes := make([]byte, 16)
	if _, err := rand.Read(randomBytes); err != nil {
		return nil, err
	}

	plaintext := base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(randomBytes)
	hash := sha256.Sum256([]byte(plaintext))

	token := &Token{
		Plaintext: plaintext,
		Hash:      hash[:],
		UserId:    userID,
		Scope:     scope,
		ExpireAt:  time.Now().Add(ttl),
	}

	return token, nil
}

func ValidateTokenPlaintext(v *validator.Validator, plaintext string) {
	common.Assert(v != nil)

	v.Check(plaintext != "", "token", "must present")
	v.Check(len(plaintext) == 26, "token", "length must be 26")
}

func (m TokenModel) New(userID int64, ttl time.Duration, scope string) (*Token, error) {
	token, err := generateToken(userID, ttl, scope)
	if err != nil {
		return nil, err
	}

	if err := m.Insert(token); err != nil {
		return nil, err
	}

	return token, nil
}

func (m TokenModel) Insert(token *Token) error {
	query := `insert into tokens(hash,user_id,expire_at,scope)
	values($1,$2,$3,$4)
	`

	args := []any{token.Hash, token.UserId, token.ExpireAt, token.Scope}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if _, err := m.DB.ExecContext(ctx, query, args...); err != nil {
		return err
	}

	return nil
}

func (m TokenModel) DeleteAllForUser(scope string, userID int64) error {
	query := `delete from tokens where scope = $1 and user_id = $2`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if _, err := m.DB.ExecContext(ctx, query, scope, userID); err != nil {
		return err
	}

	return nil
}
