INSERT INTO
  `permissions` (`name`, `key`, `group`, `order`)
VALUES
  ('Super Admin', 'ADMIN.SUPER', 'ADMIN', 1),
  ('View Admin', 'ADMIN.VIEW', 'ADMIN', 2),
  ('Modify Admin', 'ADMIN.MODIFY', 'ADMIN', 3),
  ('Delete Admin', 'ADMIN.DELETE', 'ADMIN', 4),
  (
    'Super Admin Role',
    'ADMIN_ROLE.SUPER',
    'ADMIN_ROLE',
    1
  ),
  (
    'View Admin Role',
    'ADMIN_ROLE.VIEW',
    'ADMIN_ROLE',
    2
  ),
  (
    'Modify Admin Role',
    'ADMIN_ROLE.MODIFY',
    'ADMIN_ROLE',
    3
  ),
  (
    'Delete Admin Role',
    'ADMIN_ROLE.DELETE',
    'ADMIN_ROLE',
    4
  )