-- Create "admin_roles" table
CREATE TABLE `admin_roles` (`admin_id` varchar(255) NOT NULL, `role_id` varchar(255) NOT NULL, PRIMARY KEY (`admin_id`, `role_id`), INDEX `admin_roles_role_id` (`role_id`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- Create "admins" table
CREATE TABLE `admins` (`id` varchar(255) NOT NULL, `created_at` timestamp NOT NULL, `updated_at` timestamp NOT NULL, `is_enable` bool NOT NULL DEFAULT 1, `deleted_at` bigint NOT NULL DEFAULT 0, `username` varchar(255) NOT NULL, `password` varchar(255) NOT NULL, `whitelist_ips` json NOT NULL, `display_name` varchar(255) NULL, PRIMARY KEY (`id`), INDEX `admin_deleted_at` (`deleted_at`), UNIQUE INDEX `admin_username_deleted_at` (`username`, `deleted_at`), UNIQUE INDEX `username` (`username`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- Create "comic_chapters" table
CREATE TABLE `comic_chapters` (`id` varchar(255) NOT NULL, `created_at` timestamp NOT NULL, `updated_at` timestamp NOT NULL, `chapter` bigint unsigned NOT NULL, `title` varchar(255) NULL, `volumn` varchar(255) NULL, `lang` varchar(255) NOT NULL, `up_count` bigint unsigned NOT NULL DEFAULT 0, `down_count` bigint unsigned NOT NULL DEFAULT 0, `is_last_chapter` bool NOT NULL DEFAULT 0, `comic_chapters` varchar(255) NULL, PRIMARY KEY (`id`), INDEX `comic_chapters_comics_chapters` (`comic_chapters`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- Create "comic_imgs" table
CREATE TABLE `comic_imgs` (`id` varchar(255) NOT NULL, `created_at` timestamp NOT NULL, `updated_at` timestamp NOT NULL, `b2key` varchar(255) NOT NULL, `height` bigint NOT NULL, `name` varchar(255) NOT NULL, `optimized_size` bigint NOT NULL, `size` bigint NOT NULL, `width` bigint NOT NULL, `comic_chapter_imgs` varchar(255) NULL, PRIMARY KEY (`id`), UNIQUE INDEX `b2key` (`b2key`), INDEX `comic_imgs_comic_chapters_imgs` (`comic_chapter_imgs`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- Create "comics" table
CREATE TABLE `comics` (`id` varchar(255) NOT NULL, `created_at` timestamp NOT NULL, `updated_at` timestamp NOT NULL, `chapter` bigint unsigned NOT NULL, `title` varchar(255) NOT NULL, `slug` varchar(255) NOT NULL, `covers` json NOT NULL, `status` varchar(255) NOT NULL, `is_translate_completed` bool NOT NULL DEFAULT 0, `up_count` bigint unsigned NOT NULL DEFAULT 0, `last_chapter_id` varchar(255) NULL, `final_chapter_id` varchar(255) NULL, PRIMARY KEY (`id`), INDEX `comics_comic_chapters_final_chapter` (`final_chapter_id`), INDEX `comics_comic_chapters_last_chapter` (`last_chapter_id`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- Create "genres" table
CREATE TABLE `genres` (`id` varchar(255) NOT NULL, `created_at` timestamp NOT NULL, `updated_at` timestamp NOT NULL, `name` varchar(255) NOT NULL, `type` enum('comic') NOT NULL, PRIMARY KEY (`id`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- Create "permissions" table
CREATE TABLE `permissions` (`id` varchar(255) NOT NULL, `created_at` timestamp NOT NULL, `updated_at` timestamp NOT NULL, `group` varchar(255) NOT NULL, `name` varchar(255) NOT NULL, `key` varchar(255) NOT NULL, `order` bigint NOT NULL, PRIMARY KEY (`id`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- Create "role_permissions" table
CREATE TABLE `role_permissions` (`role_id` varchar(255) NOT NULL, `permission_id` varchar(255) NOT NULL, PRIMARY KEY (`role_id`, `permission_id`), INDEX `role_permissions_permission_id` (`permission_id`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- Create "role_routes" table
CREATE TABLE `role_routes` (`role_id` varchar(255) NOT NULL, `route_id` varchar(255) NOT NULL, PRIMARY KEY (`role_id`, `route_id`), INDEX `role_routes_route_id` (`route_id`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- Create "roles" table
CREATE TABLE `roles` (`id` varchar(255) NOT NULL, `created_at` timestamp NOT NULL, `updated_at` timestamp NOT NULL, `is_enable` bool NOT NULL DEFAULT 1, `deleted_at` bigint NOT NULL DEFAULT 0, `name` varchar(255) NOT NULL, `description` varchar(255) NOT NULL, `order` bigint NOT NULL, `is_changeable` bool NOT NULL, PRIMARY KEY (`id`), INDEX `role_deleted_at` (`deleted_at`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- Create "routes" table
CREATE TABLE `routes` (`id` varchar(255) NOT NULL, `created_at` timestamp NOT NULL, `updated_at` timestamp NOT NULL, `is_enable` bool NOT NULL DEFAULT 1, `deleted_at` bigint NOT NULL DEFAULT 0, `path` varchar(255) NOT NULL, `component` varchar(255) NOT NULL, `redirect` varchar(255) NULL, `name` varchar(255) NOT NULL, `order` bigint NULL DEFAULT 0, `type` enum('cata_log','menu','button','external_link') NOT NULL DEFAULT "cata_log", `meta` json NOT NULL, `parent_id` varchar(255) NULL, PRIMARY KEY (`id`), INDEX `route_deleted_at` (`deleted_at`), UNIQUE INDEX `route_path_parent_id_type_deleted_at` (`path`, `parent_id`, `type`, `deleted_at`), INDEX `routes_routes_children` (`parent_id`), CONSTRAINT `routes_routes_children` FOREIGN KEY (`parent_id`) REFERENCES `routes` (`id`) ON UPDATE NO ACTION ON DELETE SET NULL) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- Modify "admin_roles" table
ALTER TABLE `admin_roles` ADD CONSTRAINT `admin_roles_admin_id` FOREIGN KEY (`admin_id`) REFERENCES `admins` (`id`) ON UPDATE NO ACTION ON DELETE CASCADE, ADD CONSTRAINT `admin_roles_role_id` FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`) ON UPDATE NO ACTION ON DELETE CASCADE;
-- Modify "comic_chapters" table
ALTER TABLE `comic_chapters` ADD CONSTRAINT `comic_chapters_comics_chapters` FOREIGN KEY (`comic_chapters`) REFERENCES `comics` (`id`) ON UPDATE NO ACTION ON DELETE SET NULL;
-- Modify "comic_imgs" table
ALTER TABLE `comic_imgs` ADD CONSTRAINT `comic_imgs_comic_chapters_imgs` FOREIGN KEY (`comic_chapter_imgs`) REFERENCES `comic_chapters` (`id`) ON UPDATE NO ACTION ON DELETE SET NULL;
-- Modify "comics" table
ALTER TABLE `comics` ADD CONSTRAINT `comics_comic_chapters_final_chapter` FOREIGN KEY (`final_chapter_id`) REFERENCES `comic_chapters` (`id`) ON UPDATE NO ACTION ON DELETE SET NULL, ADD CONSTRAINT `comics_comic_chapters_last_chapter` FOREIGN KEY (`last_chapter_id`) REFERENCES `comic_chapters` (`id`) ON UPDATE NO ACTION ON DELETE SET NULL;
-- Modify "role_permissions" table
ALTER TABLE `role_permissions` ADD CONSTRAINT `role_permissions_permission_id` FOREIGN KEY (`permission_id`) REFERENCES `permissions` (`id`) ON UPDATE NO ACTION ON DELETE CASCADE, ADD CONSTRAINT `role_permissions_role_id` FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`) ON UPDATE NO ACTION ON DELETE CASCADE;
-- Modify "role_routes" table
ALTER TABLE `role_routes` ADD CONSTRAINT `role_routes_role_id` FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`) ON UPDATE NO ACTION ON DELETE CASCADE, ADD CONSTRAINT `role_routes_route_id` FOREIGN KEY (`route_id`) REFERENCES `routes` (`id`) ON UPDATE NO ACTION ON DELETE CASCADE;
