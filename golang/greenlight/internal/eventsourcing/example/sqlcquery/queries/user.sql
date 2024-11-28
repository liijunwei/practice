-- name: GetAccountByID :one
SELECT * from account where id = @id;

-- name: GetAccountByIDLocked :one
SELECT * from account where id = @id for update;

-- upsert
-- name: CreateAccount :exec
INSERT into account (id, balance, available, pending, version, updated_at, created_at)
VALUES (@id, @balance, @available, @pending, @version, @updated_at, @created_at)
ON CONFLICT (id)
DO UPDATE SET balance = @balance, available = @available, pending = @pending, version = @version, updated_at = @updated_at;
