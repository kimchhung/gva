-- Modify "admins" table
ALTER TABLE `admins` ADD UNIQUE INDEX `username` (`username`);
-- Create "todos" table
CREATE TABLE `todos` (`id` bigint NOT NULL AUTO_INCREMENT, `created_at` timestamp NOT NULL, `updated_at` timestamp NOT NULL, `name` varchar(255) NOT NULL, PRIMARY KEY (`id`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
