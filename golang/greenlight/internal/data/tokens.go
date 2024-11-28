package data

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base32"
	"greenlight/internal/assert"
	"greenlight/internal/sqlcdb"
	"greenlight/internal/validator"
	"time"
)

const ScopeActivation = "activation"
const ScopeAuthentication = "authentication"

type Token struct {
	Plaintext string    `json:"token"`
	Hash      []byte    `json:"-"`
	UserId    int64     `json:"-"`
	Scope     string    `json:"-"`
	ExpireAt  time.Time `json:"expire_at"` // expect format: RFC3339
}

type TokenModel struct {
	queries *sqlcdb.Queries
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
	assert.Assert(v != nil)

	v.Check(plaintext != "", "token", "must present")
	v.Check(len(plaintext) == 26, "token", "length must be 26")
}

func (m TokenModel) New(ctx context.Context, userID int64, ttl time.Duration, scope string) (*Token, error) {
	token, err := generateToken(userID, ttl, scope)
	if err != nil {
		return nil, err
	}

	if err := m.Create(ctx, token); err != nil {
		return nil, err
	}

	return token, nil
}

func (m TokenModel) Create(ctx context.Context, token *Token) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err := m.queries.CreateScopedUserToken(ctx, &sqlcdb.CreateScopedUserTokenParams{
		Hash:     token.Hash,
		UserID:   token.UserId,
		ExpireAt: token.ExpireAt,
		Scope:    token.Scope,
	}); err != nil {
		return err
	}

	return nil
}

func (m TokenModel) DeleteAllForUser(ctx context.Context, scope string, userID int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err := m.queries.DeleteScopedUserTokens(ctx, &sqlcdb.DeleteScopedUserTokensParams{
		Scope:  scope,
		UserID: userID,
	}); err != nil {
		return err
	}

	return nil
}
