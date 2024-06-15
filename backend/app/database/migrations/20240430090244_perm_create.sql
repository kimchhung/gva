-- Modify "permissions" table
ALTER TABLE `permissions` ADD COLUMN `created_at` timestamp NOT NULL, ADD COLUMN `updated_at` timestamp NOT NULL;
