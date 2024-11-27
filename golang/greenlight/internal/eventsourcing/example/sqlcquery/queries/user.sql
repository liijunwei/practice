-- name: GetAccountByID :one
SELECT * from account where id = @id;

-- name: GetAccountByIDLocked :one
SELECT * from account where id = @id;

-- name: CreateAccount :exec
INSERT into account (id, balance, available, pending, version, updated_at, created_at)
VALUES (@id, @balance, @available, @pending, @version, @updated_at, @created_at);
