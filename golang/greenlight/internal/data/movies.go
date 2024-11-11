package data

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"greenlight/internal/validator"
	"time"

	"github.com/lib/pq"
)

type Movie struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"`
	Title     string    `json:"title"`
	Year      int32     `json:"year,omitempty,string"`    // release year
	Runtime   Runtime   `json:"runtime,omitempty,string"` // runtime in minutes
	Genres    []string  `json:"genres,omitempty"`         // e.g. romance, comedy...
	Version   int32     `json:"version"`
}

func ValidateMovie(v *validator.Validator, movie *Movie) {
	v.Check(movie.Title != "", "title", "must be provided")
	v.Check(len(movie.Title) <= 500, "title", "must not be more than 500 bytes long")

	v.Check(movie.Year != 0, "year", "must be provided")
	v.Check(movie.Year >= 1888, "year", "must be greater than 1888")
	v.Check(movie.Year <= int32(time.Now().Year()), "year", "must not be in the future")

	v.Check(movie.Runtime != 0, "runtime", "must be provided")
	v.Check(movie.Runtime > 0, "runtime", "must be a positive integer")

	v.Check(movie.Genres != nil, "genres", "must be provided")
	v.Check(len(movie.Genres) >= 1, "genres", "must contain at least 1 genre")
	v.Check(len(movie.Genres) <= 5, "genres", "must not contain more than 5 genres")
	v.Check(v.Unique(movie.Genres), "genres", "must not contain duplicate values")
}

var _ Movies = &MovieModel{}

type MovieModel struct {
	DB *sql.DB
}

func (m MovieModel) Insert(movie *Movie) error {
	query := `insert into movies(title, year, runtime, genres)
	values($1, $2, $3, $4)
	returning id, created_at, version;`

	args := []any{movie.Title, movie.Year, movie.Runtime, pq.Array(movie.Genres)}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return m.DB.QueryRowContext(ctx, query, args...).Scan(&movie.ID, &movie.CreatedAt, &movie.Version)
}

func (m MovieModel) GetAll(title string, genres []string, filters Filters) ([]*Movie, Metadata, error) {
	// window function in pg
	query := fmt.Sprintf(`select count(*) over(),id,created_at,title,year,runtime,genres,version
	from movies
	where (to_tsvector('english', title) @@ plainto_tsquery('english', $1) OR $1 = '')
	AND (genres @> $2 OR $2 = '{}')
	order by %s %s, id asc
	limit $3 offset $4`, filters.sortColumn(), filters.sortDirection())

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := m.DB.QueryContext(ctx, query, title, pq.Array(genres), filters.limit(), filters.offset())
	if err != nil {
		return nil, emptyMetadata(), err
	}

	defer rows.Close()

	var totalRecords int
	movies := make([]*Movie, 0, 20)

	for rows.Next() {
		var movie Movie

		if err := rows.Scan(&totalRecords, &movie.ID, &movie.CreatedAt, &movie.Title, &movie.Year, &movie.Runtime, pq.Array(&movie.Genres), &movie.Version); err != nil {
			return nil, emptyMetadata(), err
		}
		movies = append(movies, &movie)
	}

	if err := rows.Err(); err != nil {
		return nil, emptyMetadata(), err
	}

	metadata := getMetadata(totalRecords, filters.Page, filters.limit())

	return movies, metadata, nil
}

func (m MovieModel) Get(id int64) (*Movie, error) {
	if id < 1 {
		return nil, ErrRecordNotFound
	}

	query := `select id,created_at,title,year,runtime,genres,version from movies where id = $1`

	var movie Movie

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err := m.DB.QueryRowContext(ctx, query, id).Scan(&movie.ID, &movie.CreatedAt, &movie.Title, &movie.Year, &movie.Runtime, pq.Array(&movie.Genres), &movie.Version); err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &movie, nil
}

func (m MovieModel) Update(movie *Movie) error {
	query := `update movies set title = $1, year = $2, runtime = $3, genres = $4, version = version+1, updated_at = now()
	where id = $5 and version = $6
	returning version;`

	args := []any{movie.Title, movie.Year, movie.Runtime, pq.Array(movie.Genres), movie.ID, movie.Version}

	// fmt.Println("version:", movie.Version)
	// time.Sleep(3 * time.Second) // to mock concurrent update

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err := m.DB.QueryRowContext(ctx, query, args...).Scan(&movie.Version); err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return ErrStaleObject
		default:
			return err
		}
	}

	return nil
}

func (m MovieModel) Delete(id int64) error {
	if id < 1 {
		return ErrRecordNotFound
	}

	query := `DELETE from movies where id = $1;`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	result, err := m.DB.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrRecordNotFound
	}

	return nil
}