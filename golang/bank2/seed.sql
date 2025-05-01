INSERT INTO
  "users" (
    "id",
    "username",
    "password",
    "email",
    "created_at",
    "updated_at"
  )
VALUES
  (
    '6171b3ed-36b7-47ef-b070-ac064ae78907',
    'Darrell.Marquardt',
    '8e4d9nGCg31Ni7C',
    'Jaiden_Heathcote222@example.net',
    '2025-04-28 16:33:36',
    '2025-04-28 16:33:36'
  ),
  (
    'fed31898-c5c0-48dc-937d-b981b876f1f8',
    'Darrell.Marquardt',
    '8e4d9nGCg31Ni7C',
    'Jaiden_Heathcote111@example.net',
    '2025-04-28 16:33:36',
    '2025-04-28 16:33:36'
  );

INSERT INTO
  "accounts" (
    "id",
    "user_id",
    "currency",
    "available",
    "created_at",
    "updated_at"
  )
VALUES
  (
    '79726a51-56bb-4bca-b1f0-2cbded8cac8a',
    'fed31898-c5c0-48dc-937d-b981b876f1f8',
    'CNY',
    '100',
    '2025-04-28 16:33:36',
    '2025-04-28 16:33:36'
  ),
  (
    '83b93359-f715-4fb2-9c1e-f3e5645793e5',
    '6171b3ed-36b7-47ef-b070-ac064ae78907',
    'CNY',
    '100',
    '2025-04-28 16:33:36',
    '2025-04-28 16:33:36'
  );
