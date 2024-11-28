-- name: GetAccountByID :one
SELECT * from accounts where id = @id;

-- name: GetAccountByIDLocked :one
SELECT * from accounts where id = @id for update;

-- upsert
-- name: UpsertAccount :exec
INSERT into accounts (id, balance, available, pending, version, updated_at, created_at)
VALUES (@id, @balance, @available, @pending, @version, @updated_at, @created_at)
ON CONFLICT (id)
DO UPDATE SET
  balance    = EXCLUDED.balance,
  available  = EXCLUDED.available,
  pending    = EXCLUDED.pending,
  version    = EXCLUDED.version,
  updated_at = EXCLUDED.updated_at;
