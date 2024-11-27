// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: movies.sql

package sqlcdb

import (
	"context"
	"time"

	"github.com/lib/pq"
)

const createMovie = `-- name: CreateMovie :one
insert into
  movies(title, year, runtime, genres)
values
  ($1, $2, $3, $4) returning id,
  created_at,
  version
`

type CreateMovieParams struct {
	Title   string
	Year    int32
	Runtime int32
	Genres  []string
}

type CreateMovieRow struct {
	ID        int64
	CreatedAt time.Time
	Version   int32
}

func (q *Queries) CreateMovie(ctx context.Context, arg CreateMovieParams) (CreateMovieRow, error) {
	row := q.db.QueryRowContext(ctx, createMovie,
		arg.Title,
		arg.Year,
		arg.Runtime,
		pq.Array(arg.Genres),
	)
	var i CreateMovieRow
	err := row.Scan(&i.ID, &i.CreatedAt, &i.Version)
	return i, err
}

const deleteMovie = `-- name: DeleteMovie :exec
DELETE from movies where id = $1
`

func (q *Queries) DeleteMovie(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteMovie, id)
	return err
}

const getAllMovies = `-- name: GetAllMovies :many
select
  count(*) over(),
  id,
  created_at,
  title,
  year,
  runtime,
  genres,
  version
from
  movies
`

type GetAllMoviesRow struct {
	Count     int64
	ID        int64
	CreatedAt time.Time
	Title     string
	Year      int32
	Runtime   int32
	Genres    []string
	Version   int32
}

// TODO add pagination and filtering back
func (q *Queries) GetAllMovies(ctx context.Context) ([]GetAllMoviesRow, error) {
	rows, err := q.db.QueryContext(ctx, getAllMovies)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetAllMoviesRow
	for rows.Next() {
		var i GetAllMoviesRow
		if err := rows.Scan(
			&i.Count,
			&i.ID,
			&i.CreatedAt,
			&i.Title,
			&i.Year,
			&i.Runtime,
			pq.Array(&i.Genres),
			&i.Version,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getMovieByID = `-- name: GetMovieByID :one
select id,created_at,title,year,runtime,genres,version from movies where id = $1
`

type GetMovieByIDRow struct {
	ID        int64
	CreatedAt time.Time
	Title     string
	Year      int32
	Runtime   int32
	Genres    []string
	Version   int32
}

func (q *Queries) GetMovieByID(ctx context.Context, id int64) (GetMovieByIDRow, error) {
	row := q.db.QueryRowContext(ctx, getMovieByID, id)
	var i GetMovieByIDRow
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.Title,
		&i.Year,
		&i.Runtime,
		pq.Array(&i.Genres),
		&i.Version,
	)
	return i, err
}

const updateMovie = `-- name: UpdateMovie :exec
update movies
set title = $1, year = $2, runtime = $3, genres = $4, version = version+1, updated_at = now()
where id = $5 and version = $6
returning version
`

type UpdateMovieParams struct {
	Title   string
	Year    int32
	Runtime int32
	Genres  []string
	ID      int64
	Version int32
}

func (q *Queries) UpdateMovie(ctx context.Context, arg UpdateMovieParams) error {
	_, err := q.db.ExecContext(ctx, updateMovie,
		arg.Title,
		arg.Year,
		arg.Runtime,
		pq.Array(arg.Genres),
		arg.ID,
		arg.Version,
	)
	return err
}