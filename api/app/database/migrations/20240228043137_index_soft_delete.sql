-- Modify "admins" table
ALTER TABLE
  `admins`
ADD
  INDEX `admin_deleted_at` (`deleted_at`),
ADD
  UNIQUE INDEX `admin_username_deleted_at` (`username`, `deleted_at`);

-- Modify "roles" table
ALTER TABLE
  `roles`
ADD
  INDEX `role_deleted_at` (`deleted_at`);