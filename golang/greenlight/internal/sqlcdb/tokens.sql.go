// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: tokens.sql

package sqlcdb

import (
	"context"
	"time"
)

const CreateScopedUserToken = `-- name: CreateScopedUserToken :exec
insert into tokens(hash, user_id, expire_at, scope)
values($1,$2,$3,$4)
`

type CreateScopedUserTokenParams struct {
	Hash     []byte
	UserID   int64
	ExpireAt time.Time
	Scope    string
}

func (q *Queries) CreateScopedUserToken(ctx context.Context, arg *CreateScopedUserTokenParams) error {
	_, err := q.db.ExecContext(ctx, CreateScopedUserToken,
		arg.Hash,
		arg.UserID,
		arg.ExpireAt,
		arg.Scope,
	)
	return err
}

const DeleteScopedUserTokens = `-- name: DeleteScopedUserTokens :exec
delete from tokens where scope = $1 and user_id = $2
`

type DeleteScopedUserTokensParams struct {
	Scope  string
	UserID int64
}

func (q *Queries) DeleteScopedUserTokens(ctx context.Context, arg *DeleteScopedUserTokensParams) error {
	_, err := q.db.ExecContext(ctx, DeleteScopedUserTokens, arg.Scope, arg.UserID)
	return err
}
