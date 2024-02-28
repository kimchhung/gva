-- Modify "admins" table
ALTER TABLE `admins` MODIFY COLUMN `deleted_at` bigint NOT NULL DEFAULT 0;
-- Modify "roles" table
ALTER TABLE `roles` MODIFY COLUMN `deleted_at` bigint NOT NULL DEFAULT 0;
