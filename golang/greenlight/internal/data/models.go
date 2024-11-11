package data

import (
	"database/sql"
	"errors"
)

var ErrRecordNotFound = errors.New("record not found")
var ErrStaleObject = errors.New("trying to update stale object")

type Movies interface {
	Insert(movie *Movie) error
	Get(id int64) (*Movie, error)
	GetAll(title string, genres []string, filters Filters) ([]*Movie, Metadata, error)
	Update(movie *Movie) error
	Delete(id int64) error
}

type Models struct {
	Movies Movies
}

func NewModels(db *sql.DB) Models {
	return Models{
		Movies: MovieModel{DB: db},
	}
}
