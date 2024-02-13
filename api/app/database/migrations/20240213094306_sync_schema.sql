-- Modify "admins" table
ALTER TABLE `admins` DROP COLUMN `deleted_at`, ADD COLUMN `display_name` varchar(255) NOT NULL;
-- Modify "articles" table
ALTER TABLE `articles` ADD COLUMN `created_at` timestamp NOT NULL, ADD COLUMN `updated_at` timestamp NOT NULL;
