-- name: GetUserByID :one
select
  *
from
  users
where
  id = @id;

-- name: GetUserByEmail :one
select
  id,
  name,
  email,
  password_hash,
  status,
  version,
  created_at,
  updated_at
from
  users
where
  email = @email;

-- name: GetByToken :one
select
  u.id,
  u.created_at,
  u.name,
  u.email,
  u.password_hash,
  u.status,
  u.version
from
  users u
  inner join tokens t on u.id = t.user_id
where
  t.hash = @hash
  and t.scope = @scope
  and t.expire_at > @expire_at;
