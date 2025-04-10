// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: query.sql

package sqlcdb

import (
	"context"
)

const checkOriginalExists = `-- name: CheckOriginalExists :one
SELECT
  id, original, shorturl, created_at, updated_at
FROM
  shorturls
WHERE
  original = ?1
`

// TODO try this first, then optimize with bloom filter to see it's performance gain
func (q *Queries) CheckOriginalExists(ctx context.Context, original string) (Shorturl, error) {
	row := q.db.QueryRowContext(ctx, checkOriginalExists, original)
	var i Shorturl
	err := row.Scan(
		&i.ID,
		&i.Original,
		&i.Shorturl,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const createShorturl = `-- name: CreateShorturl :one
INSERT INTO
  shorturls(original, shorturl)
VALUES
  (?1, ?2) RETURNING id, original, shorturl, created_at, updated_at
`

type CreateShorturlParams struct {
	Original string
	Shorturl string
}

func (q *Queries) CreateShorturl(ctx context.Context, arg CreateShorturlParams) (Shorturl, error) {
	row := q.db.QueryRowContext(ctx, createShorturl, arg.Original, arg.Shorturl)
	var i Shorturl
	err := row.Scan(
		&i.ID,
		&i.Original,
		&i.Shorturl,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getOriginalByShort = `-- name: GetOriginalByShort :one
SELECT
  id, original, shorturl, created_at, updated_at
FROM
  shorturls
WHERE
  shorturl = ?1
`

func (q *Queries) GetOriginalByShort(ctx context.Context, shorturl string) (Shorturl, error) {
	row := q.db.QueryRowContext(ctx, getOriginalByShort, shorturl)
	var i Shorturl
	err := row.Scan(
		&i.ID,
		&i.Original,
		&i.Shorturl,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listShorturls = `-- name: ListShorturls :many
SELECT
  id, original, shorturl, created_at, updated_at
FROM
  shorturls
`

func (q *Queries) ListShorturls(ctx context.Context) ([]Shorturl, error) {
	rows, err := q.db.QueryContext(ctx, listShorturls)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Shorturl
	for rows.Next() {
		var i Shorturl
		if err := rows.Scan(
			&i.ID,
			&i.Original,
			&i.Shorturl,
			&i.CreatedAt,
			&i.UpdatedAt,
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
