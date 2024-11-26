-- name: GetAccountByID :one
SELECT * from account where id = $1;

-- name: GetAccountByIDLocked :one
SELECT * from account where id = $1 for update;

-- name: Save :exec
INSERT into account
  (id, balance, available, pending, version, updated_at, created_at)
  VALUES ($1, $2, $3, $4, $5, $6, $7)
  ON CONFLICT (id) DO UPDATE SET balance = $2, available = $3, pending = $4, version = $5, updated_at = $6;
