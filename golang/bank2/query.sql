-- name: CreateUser :one
insert into
  users(
    id,
    username,
    password,
    email
  )
values
  (
    @id,
    @username,
    @password,
    @email
  ) returning *;

-- name: CreateAccount :one
insert into
  accounts(
    id,
    user_id,
    currency,
    available
  )
values
  (
    @id,
    @user_id,
    @currency,
    @available
  ) returning *;

-- name: GetAllAccounts :many
select
  a.id account_id,
  u.username,
  u.email,
  a.currency,
  a.available,
  a.lock_version,
  a.created_at,
  a.updated_at
from
  users u
  join accounts a on u.id = a.user_id
order by
  u.username asc,
  a.currency asc,
  a.available desc;

-- name: GetAccount :one
select
  *
from
  accounts
where
  id = @id;

-- name: CreditAccount :one
update
  accounts
set
  available = available + @amount,
  lock_version = lock_version + 1
where
  id = @id returning *;

-- name: DebitAccount :one
update
  accounts
set
  available = available - @amount,
  lock_version = lock_version + 1
where
  id = @id returning *;

-- name: CreateAccountEvent :exec
insert into
  account_events(id, account_id, amount)
values
  (@id, @account_id, @amount);

-- name: CreateTransaction :one
insert into
  transactions(
    id,
    from_account_id,
    to_account_id,
    amount,
    description,
    kind
  )
values
  (
    @id,
    @from_account_id,
    @to_account_id,
    @amount,
    @description,
    @kind
  ) returning *;
