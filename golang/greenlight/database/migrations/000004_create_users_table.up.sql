create table if not exists users(
  id bigserial primary key,
  name text not null,
  email citext unique not null, -- case insensitive text
  password_hash bytea not null,
  status varchar(50) not null,
  version integer not null default 1,
  created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
  updated_at timestamp(0) with time zone NOT NULL DEFAULT NOW()
);
