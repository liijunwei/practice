package data

import (
	"context"
	"database/sql"
	"errors"
	"greenlight/internal/sqlcdb"
	"time"
)

var ErrRecordNotFound = errors.New("record not found")
var ErrStaleObject = errors.New("trying to update stale object")

type Models struct {
	Movies      Movies
	Users       Users
	Tokens      Tokens
	Permissions PermissionsDB
}

func NewModels(db *sql.DB, queries *sqlcdb.Queries) Models {
	return Models{
		Movies:      MovieModel{queries: queries},
		Users:       UserModel{DB: db, queries: queries},
		Tokens:      TokenModel{queries: queries},
		Permissions: PermissinModel{queries: queries},
	}
}

type Movies interface {
	Create(ctx context.Context, movie *Movie) error
	Get(ctx context.Context, id int64) (*Movie, error)
	GetAll(ctx context.Context, title string, genres []string, filters Filters) ([]*Movie, Metadata, error)
	Update(ctx context.Context, movie *Movie) error
	Delete(ctx context.Context, id int64) error
}

type Users interface {
	Create(ctx context.Context, user *User) error
	GetByEmail(ctx context.Context, email string) (*User, error)
	Update(ctx context.Context, user *User) error
	GetByToken(ctx context.Context, tokenScope, plaintextToken string) (*User, error)
}

type Tokens interface {
	New(ctx context.Context, userID int64, ttl time.Duration, scope string) (*Token, error)
	Create(ctx context.Context, token *Token) error
	DeleteAllForUser(ctx context.Context, scope string, userID int64) error
}

type PermissionsDB interface {
	GetAllForUser(ctx context.Context, userID int64) (Permissions, error)
	AddForUser(ctx context.Context, userID int64, codes ...string) error
}
