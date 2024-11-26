CREATE TABLE account (
  id         uuid PRIMARY KEY,
  balance    numeric not null,
  available  numeric not null,
  pending    numeric not null,
  created_at timestamp not null,
  updated_at timestamp not null,
  version    int not null
);
