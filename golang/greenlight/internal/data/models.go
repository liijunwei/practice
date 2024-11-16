package data

import (
	"database/sql"
	"errors"
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

func NewModels(db *sql.DB) Models {
	return Models{
		Movies:      MovieModel{DB: db},
		Users:       UserModel{DB: db},
		Tokens:      TokenModel{DB: db},
		Permissions: PermissinModel{DB: db},
	}
}

type Movies interface {
	Create(movie *Movie) error
	Get(id int64) (*Movie, error)
	GetAll(title string, genres []string, filters Filters) ([]*Movie, Metadata, error)
	Update(movie *Movie) error
	Delete(id int64) error
}

type Users interface {
	Create(user *User) error
	GetByEmail(email string) (*User, error)
	Update(user *User) error
	GetByToken(tokenScope, plaintextToken string) (*User, error)
}

type Tokens interface {
	New(userID int64, ttl time.Duration, scope string) (*Token, error)
	Create(token *Token) error
	DeleteAllForUser(scope string, userID int64) error
}

type PermissionsDB interface {
	GetAllForUser(userID int64) (Permissions, error)
	AddForUser(userID int64, codes ...string) error
}
