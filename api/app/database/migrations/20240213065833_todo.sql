-- Create "admins" table
CREATE TABLE `admins` (`id` bigint NOT NULL AUTO_INCREMENT, `created_at` timestamp NOT NULL, `updated_at` timestamp NOT NULL, `deleted_at` timestamp NOT NULL, `name` varchar(255) NOT NULL, PRIMARY KEY (`id`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- Create "articles" table
CREATE TABLE `articles` (`id` bigint NOT NULL AUTO_INCREMENT, `title` varchar(255) NOT NULL, `content` longtext NOT NULL, PRIMARY KEY (`id`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
