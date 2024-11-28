package data

import (
	"context"
	"database/sql"
	"errors"
	"greenlight/internal/sqlcdb"
	"greenlight/internal/validator"
	"time"
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
	v.Check(validator.Unique(movie.Genres), "genres", "must not contain duplicate values")
}

var _ Movies = &MovieModel{}

type MovieModel struct {
	queries *sqlcdb.Queries
}

// TODO return new instance instead of mutating input variable
func (m MovieModel) Create(ctx context.Context, movie *Movie) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	row, err := m.queries.CreateMovie(ctx, &sqlcdb.CreateMovieParams{
		movie.Title,
		movie.Year,
		int32(movie.Runtime),
		movie.Genres,
	})

	if err != nil {
		return err
	}

	movie.ID = row.ID
	movie.CreatedAt = row.CreatedAt
	movie.Version = row.Version

	return nil
}

func (m MovieModel) GetAll(ctx context.Context, title string, genres []string, filters Filters) ([]*Movie, Metadata, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := m.queries.GetAllMovies(ctx)
	if err != nil {
		return nil, emptyMetadata(), err
	}

	movies := make([]*Movie, 0, len(rows))

	for _, row := range rows {
		movies = append(movies, &Movie{
			ID:        row.ID,
			CreatedAt: row.CreatedAt,
			Title:     row.Title,
			Year:      row.Year,
			Runtime:   Runtime(row.Runtime),
			Genres:    row.Genres,
			Version:   row.Version,
		})
	}

	metadata := getMetadata(len(rows), filters.Page, filters.limit())

	return movies, metadata, nil
}

func (m MovieModel) Get(ctx context.Context, id int64) (*Movie, error) {
	if id < 1 {
		return nil, ErrRecordNotFound
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	row, err := m.queries.GetMovieByID(ctx, id)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	movie := Movie{
		ID:        row.ID,
		CreatedAt: row.CreatedAt,
		Title:     row.Title,
		Year:      row.Year,
		Runtime:   Runtime(row.Runtime),
		Genres:    row.Genres,
		Version:   row.Version,
	}

	return &movie, nil
}

func (m MovieModel) Update(ctx context.Context, movie *Movie) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.queries.UpdateMovie(ctx, &sqlcdb.UpdateMovieParams{
		Title:   movie.Title,
		Year:    movie.Year,
		Runtime: int32(movie.Runtime),
		Genres:  movie.Genres,
		ID:      movie.ID,
		Version: movie.Version,
	})

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return ErrStaleObject
		default:
			return err
		}
	}

	return nil
}

func (m MovieModel) Delete(ctx context.Context, id int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// note: DELETE from movies where id = 9; does not return error even record is not found
	// TODO maybe optimze by `DELETE from movies where id = 9 returning id;`
	if err := m.queries.DeleteMovie(ctx, id); err != nil {
		return err
	}

	return nil
}
