-- Modify "admins" table
ALTER TABLE `admins` ADD COLUMN `deleted_at` int NOT NULL;
-- Modify "roles" table
ALTER TABLE `roles` ADD COLUMN `deleted_at` int NOT NULL;
