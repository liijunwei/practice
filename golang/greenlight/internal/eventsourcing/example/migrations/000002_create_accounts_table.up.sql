CREATE TABLE IF NOT EXISTS accounts (
  id         uuid PRIMARY KEY,
  balance    numeric(50,32)  not null,
  available  numeric(50,32)  not null,
  pending    numeric(50,32)  not null,
  created_at timestamp not null,
  updated_at timestamp not null,
  version    int not null
);
