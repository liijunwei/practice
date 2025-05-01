INSERT INTO
  "users" (
    "id",
    "username",
    "password",
    "email"
  )
VALUES
  (
    '6171b3ed-36b7-47ef-b070-ac064ae78907',
    'Darrell.Marquardt',
    '8e4d9nGCg31Ni7C',
    'Jaiden_Heathcote222@example.net'
  ),
  (
    'fed31898-c5c0-48dc-937d-b981b876f1f8',
    'Darrell.Marquardt',
    '8e4d9nGCg31Ni7C',
    'Jaiden_Heathcote111@example.net'
  );

INSERT INTO
  "accounts" (
    "id",
    "user_id",
    "currency",
    "available"
  )
VALUES
  (
    '79726a51-56bb-4bca-b1f0-2cbded8cac8a',
    'fed31898-c5c0-48dc-937d-b981b876f1f8',
    'CNY',
    '100'
  ),
  (
    '83b93359-f715-4fb2-9c1e-f3e5645793e5',
    '6171b3ed-36b7-47ef-b070-ac064ae78907',
    'CNY',
    '100'
  );
