// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: query.sql

package sqlcdb

import (
	"context"
)

const createTODO = `-- name: CreateTODO :exec
INSERT INTO
  todos(title)
VALUES
  (?1)
`

func (q *Queries) CreateTODO(ctx context.Context, title string) error {
	_, err := q.db.ExecContext(ctx, createTODO, title)
	return err
}

const deleteTODO = `-- name: DeleteTODO :exec
DELETE FROM
  todos
WHERE
  id = ?1
`

func (q *Queries) DeleteTODO(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteTODO, id)
	return err
}

const listTODOs = `-- name: ListTODOs :many
SELECT
  id,
  title
FROM
  todos
`

func (q *Queries) ListTODOs(ctx context.Context) ([]Todo, error) {
	rows, err := q.db.QueryContext(ctx, listTODOs)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Todo
	for rows.Next() {
		var i Todo
		if err := rows.Scan(&i.ID, &i.Title); err != nil {
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
