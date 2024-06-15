INSERT INTO
  roles (
    id,
    name,
    is_changeable,
    is_enable,
    description,
    `order`,
    created_at,
    updated_at
  )
VALUES
  (
    1,
    'SUPER_ADMIN',
    0,
    1,
    "Super Admin can control everything",
    0,
    NOW(),
    NOW()
  );

INSERT INTO
  admins (
    id,
    username,
    password,
    whitelist_ips,
    is_enable,
    created_at,
    updated_at
  )
VALUES
  (
    1,
    'admin',
    '$2a$10$HS2I1JNF10kHQw/jzhpcduDrwuVQEVnCPAx8AaO/UkDy7I9W9.6D.',
    '["0.0.0.0"]',
    1,
    NOW(),
    NOW()
  );

INSERT INTO
  admin_roles (admin_id, role_id) VALUE(1, 1)