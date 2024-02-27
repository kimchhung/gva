-- Create "admins" table
CREATE TABLE `admins` (`id` bigint NOT NULL AUTO_INCREMENT, `created_at` timestamp NOT NULL, `updated_at` timestamp NOT NULL, `username` varchar(255) NOT NULL, `password` varchar(255) NOT NULL, `whitelist_ips` json NOT NULL, `is_active` bool NOT NULL DEFAULT 1, `display_name` varchar(255) NOT NULL, PRIMARY KEY (`id`), UNIQUE INDEX `username` (`username`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- Create "roles" table
CREATE TABLE `roles` (`id` bigint NOT NULL AUTO_INCREMENT, `created_at` timestamp NOT NULL, `updated_at` timestamp NOT NULL, `name` varchar(255) NOT NULL, `is_active` bool NOT NULL, `is_changeable` bool NOT NULL, PRIMARY KEY (`id`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- Create "admin_roles" table
CREATE TABLE `admin_roles` (`admin_id` bigint NOT NULL, `role_id` bigint NOT NULL, PRIMARY KEY (`admin_id`, `role_id`), INDEX `admin_roles_role_id` (`role_id`), CONSTRAINT `admin_roles_admin_id` FOREIGN KEY (`admin_id`) REFERENCES `admins` (`id`) ON UPDATE NO ACTION ON DELETE CASCADE, CONSTRAINT `admin_roles_role_id` FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`) ON UPDATE NO ACTION ON DELETE CASCADE) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- Create "permissions" table
CREATE TABLE `permissions` (`id` bigint NOT NULL AUTO_INCREMENT, `name` varchar(255) NOT NULL, `key` varchar(255) NOT NULL, `group` varchar(255) NOT NULL, PRIMARY KEY (`id`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- Create "role_permissions" table
CREATE TABLE `role_permissions` (`role_id` bigint NOT NULL, `permission_id` bigint NOT NULL, PRIMARY KEY (`role_id`, `permission_id`), INDEX `role_permissions_permission_id` (`permission_id`), CONSTRAINT `role_permissions_permission_id` FOREIGN KEY (`permission_id`) REFERENCES `permissions` (`id`) ON UPDATE NO ACTION ON DELETE CASCADE, CONSTRAINT `role_permissions_role_id` FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`) ON UPDATE NO ACTION ON DELETE CASCADE) CHARSET utf8mb4 COLLATE utf8mb4_bin;
