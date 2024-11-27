// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: user.sql

package sqlcquery

import (
	"context"
	"time"

	"github.com/ericlagergren/decimal"
	"github.com/gofrs/uuid"
)

const GetAccountByID = `-- name: GetAccountByID :one
SELECT id, balance, available, pending, created_at, updated_at, version from account where id = $1
`

func (q *Queries) GetAccountByID(ctx context.Context, id uuid.UUID) (*Account, error) {
	row := q.db.QueryRow(ctx, GetAccountByID, id)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.Balance,
		&i.Available,
		&i.Pending,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Version,
	)
	return &i, err
}

const GetAccountByIDLocked = `-- name: GetAccountByIDLocked :one
SELECT id, balance, available, pending, created_at, updated_at, version from account where id = $1
`

func (q *Queries) GetAccountByIDLocked(ctx context.Context, id uuid.UUID) (*Account, error) {
	row := q.db.QueryRow(ctx, GetAccountByIDLocked, id)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.Balance,
		&i.Available,
		&i.Pending,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Version,
	)
	return &i, err
}

const Save = `-- name: Save :exec
INSERT into account (id, balance, available, pending, version, updated_at, created_at)
VALUES ($1, $2, $3, $4, $5, $6, $7)
`

type SaveParams struct {
	ID        uuid.UUID
	Balance   decimal.Big
	Available decimal.Big
	Pending   decimal.Big
	Version   int32
	UpdatedAt time.Time
	CreatedAt time.Time
}

func (q *Queries) Save(ctx context.Context, arg *SaveParams) error {
	_, err := q.db.Exec(ctx, Save,
		arg.ID,
		arg.Balance,
		arg.Available,
		arg.Pending,
		arg.Version,
		arg.UpdatedAt,
		arg.CreatedAt,
	)
	return err
}
