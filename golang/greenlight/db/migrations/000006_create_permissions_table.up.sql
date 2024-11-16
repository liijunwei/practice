create table if not exists permissions(
  id bigserial primary key,
  code text not null,
  created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
  updated_at timestamp(0) with time zone NOT NULL DEFAULT NOW()
);

create table if not exists user_permissions(
  user_id bigint not null references users on delete cascade,
  permission_id bigint not null references permissions on delete cascade,
  created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
  updated_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
  primary key(user_id, permission_id)
);
