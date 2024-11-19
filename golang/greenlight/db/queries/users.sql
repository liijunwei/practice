-- name: CreateUser :one
insert into users(name, email, password_hash, status)
values(@name, @email, @password_hash, @status)
returning id, created_at, updated_at, version;

-- name: GetUserByEmail :one
select id, name, email, password_hash, status, version, created_at, updated_at from users where email = @email;

-- name: GetUserByToken :one
select u.id, u.created_at, u.name, u.email, u.password_hash::bytea password_hash, u.status, u.version
from users u
inner join tokens t on u.id = t.user_id
where
t.hash = @hash
and t.scope = @scope
and t.expire_at > now();

-- name: UpdateUser :one
update users
set name=$1, email=$2, password_hash=$3, status=$4, version=version+1, updated_at=now()
where id=$5 and version=$6
returning version;
