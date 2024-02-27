INSERT INTO
  `roles` (
    id,
    name,
    is_active,
    is_changeable,
    created_at,
    updated_at
  )
VALUES
  (
    1,
    'SUPER_ADMIN',
    1,
    0,
    NOW(),
    NOW()
  );

INSERT INTO
  `admins` (
    username,
    password,
    whitelist_ips,
    is_active,
    created_at,
    updated_at
  )
VALUES
  (
    'admin',
    '$2a$10$HS2I1JNF10kHQw/jzhpcduDrwuVQEVnCPAx8AaO/UkDy7I9W9.6D.',
    '["0.0.0.0"]',
    1,
    NOW(),
    NOW()
  );

INSERT INTO
  `admin_roles` (admin_id, role_id)
VALUES
  (1, 1);