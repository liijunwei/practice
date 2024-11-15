create table if not exists tokens(
  hash bytea primary key,
  user_id bigint not null references users on delete cascade,
  scope text not null,
  expire_at timestamp(0) with time zone not null,
  created_at timestamp(0) with time zone not null default now(),
  updated_at timestamp(0) with time zone not null default now()
);
