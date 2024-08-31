-- Modify "role_routes" table
ALTER TABLE `role_routes` DROP FOREIGN KEY `role_routes_menu_id`, DROP FOREIGN KEY `role_routes_role_id`;
-- Create "genre_mangas" table
CREATE TABLE `genre_mangas` (`genre_id` varchar(255) NOT NULL, `manga_id` varchar(255) NOT NULL, PRIMARY KEY (`genre_id`, `manga_id`), INDEX `genre_mangas_manga_id` (`manga_id`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- Create "genres" table
CREATE TABLE `genres` (`id` varchar(255) NOT NULL, `created_at` timestamp NOT NULL, `updated_at` timestamp NOT NULL, `is_enable` bool NOT NULL DEFAULT 1, `deleted_at` bigint NOT NULL DEFAULT 0, `name` varchar(255) NOT NULL, `name_id` varchar(255) NOT NULL, PRIMARY KEY (`id`), INDEX `genre_deleted_at` (`deleted_at`), UNIQUE INDEX `genre_name_id_deleted_at` (`name_id`, `deleted_at`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- Create "manga_chapters" table
CREATE TABLE `manga_chapters` (`id` varchar(255) NOT NULL, `created_at` timestamp NOT NULL, `updated_at` timestamp NOT NULL, `title` varchar(255) NOT NULL, `img_url` varchar(255) NOT NULL, `number` bigint unsigned NOT NULL, `provider_name` varchar(255) NOT NULL, `chapter_updated_at` timestamp NOT NULL, `manga_id` varchar(255) NOT NULL, PRIMARY KEY (`id`), INDEX `manga_chapters_mangas_chapters` (`manga_id`), UNIQUE INDEX `mangachapter_provider_name_manga_id_number` (`provider_name`, `manga_id`, `number`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- Create "mangas" table
CREATE TABLE `mangas` (`id` varchar(255) NOT NULL, `created_at` timestamp NOT NULL, `updated_at` timestamp NOT NULL, `is_enable` bool NOT NULL DEFAULT 1, `deleted_at` bigint NOT NULL DEFAULT 0, `name_id` varchar(255) NOT NULL, `name` varchar(255) NOT NULL, `desc` varchar(255) NOT NULL, `prodiver` varchar(255) NOT NULL, `thumbnail_url` varchar(255) NOT NULL, `authors` json NOT NULL, PRIMARY KEY (`id`), INDEX `manga_deleted_at` (`deleted_at`), UNIQUE INDEX `manga_name_name_id_deleted_at` (`name`, `name_id`, `deleted_at`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- Modify "genre_mangas" table
ALTER TABLE `genre_mangas` ADD CONSTRAINT `genre_mangas_genre_id` FOREIGN KEY (`genre_id`) REFERENCES `genres` (`id`) ON UPDATE NO ACTION ON DELETE CASCADE, ADD CONSTRAINT `genre_mangas_manga_id` FOREIGN KEY (`manga_id`) REFERENCES `mangas` (`id`) ON UPDATE NO ACTION ON DELETE CASCADE;
-- Modify "manga_chapters" table
ALTER TABLE `manga_chapters` ADD CONSTRAINT `manga_chapters_mangas_chapters` FOREIGN KEY (`manga_id`) REFERENCES `mangas` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION;
-- Drop "menus" table
DROP TABLE `menus`;
-- Drop "regions" table
DROP TABLE `regions`;
-- Drop "role_routes" table
DROP TABLE `role_routes`;
-- Drop "todos" table
DROP TABLE `todos`;
