-- Create "admins" table
CREATE TABLE `admins` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NOT NULL,
  `updated_at` timestamp NOT NULL,
  `is_enable` bool NOT NULL DEFAULT 1,
  `deleted_at` bigint NOT NULL DEFAULT 0,
  `username` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `whitelist_ips` json NOT NULL,
  `display_name` varchar(255) NULL,
  PRIMARY KEY (`id`),
  INDEX `admin_deleted_at` (`deleted_at`),
  UNIQUE INDEX `admin_username_deleted_at` (`username`, `deleted_at`),
  UNIQUE INDEX `username` (`username`)
) CHARSET utf8mb4 COLLATE utf8mb4_bin;

-- Create "roles" table
CREATE TABLE `roles` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NOT NULL,
  `updated_at` timestamp NOT NULL,
  `is_enable` bool NOT NULL DEFAULT 1,
  `deleted_at` bigint NOT NULL DEFAULT 0,
  `name` varchar(255) NOT NULL,
  `description` varchar(255) NOT NULL,
  `order` bigint NOT NULL,
  `is_changeable` bool NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `role_deleted_at` (`deleted_at`)
) CHARSET utf8mb4 COLLATE utf8mb4_bin;

-- Create "admin_roles" table
CREATE TABLE `admin_roles` (
  `admin_id` bigint NOT NULL,
  `role_id` bigint NOT NULL,
  PRIMARY KEY (`admin_id`, `role_id`),
  INDEX `admin_roles_role_id` (`role_id`),
  CONSTRAINT `admin_roles_admin_id` FOREIGN KEY (`admin_id`) REFERENCES `admins` (`id`) ON UPDATE NO ACTION ON DELETE CASCADE,
  CONSTRAINT `admin_roles_role_id` FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`) ON UPDATE NO ACTION ON DELETE CASCADE
) CHARSET utf8mb4 COLLATE utf8mb4_bin;

-- Create "permissions" table
CREATE TABLE `permissions` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `group` varchar(255) NOT NULL,
  `name` varchar(255) NOT NULL,
  `key` varchar(255) NOT NULL,
  `order` bigint NOT NULL,
  PRIMARY KEY (`id`)
) CHARSET utf8mb4 COLLATE utf8mb4_bin;

-- Create "role_permissions" table
CREATE TABLE `role_permissions` (
  `role_id` bigint NOT NULL,
  `permission_id` bigint NOT NULL,
  PRIMARY KEY (`role_id`, `permission_id`),
  INDEX `role_permissions_permission_id` (`permission_id`),
  CONSTRAINT `role_permissions_permission_id` FOREIGN KEY (`permission_id`) REFERENCES `permissions` (`id`) ON UPDATE NO ACTION ON DELETE CASCADE,
  CONSTRAINT `role_permissions_role_id` FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`) ON UPDATE NO ACTION ON DELETE CASCADE
) CHARSET utf8mb4 COLLATE utf8mb4_bin;

-- Create "routes" table
CREATE TABLE `routes` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NOT NULL,
  `updated_at` timestamp NOT NULL,
  `is_enable` bool NOT NULL DEFAULT 1,
  `deleted_at` bigint NOT NULL DEFAULT 0,
  `path` varchar(255) NOT NULL,
  `component` varchar(255) NOT NULL,
  `redirect` varchar(255) NOT NULL,
  `name` varchar(255) NOT NULL,
  `type` bigint NOT NULL,
  `title` varchar(255) NOT NULL,
  `meta` json NOT NULL,
  `parent_id` bigint NULL,
  PRIMARY KEY (`id`),
  INDEX `route_deleted_at` (`deleted_at`),
  INDEX `routes_routes_children` (`parent_id`),
  CONSTRAINT `routes_routes_children` FOREIGN KEY (`parent_id`) REFERENCES `routes` (`id`) ON UPDATE NO ACTION ON DELETE
  SET
    NULL
) CHARSET utf8mb4 COLLATE utf8mb4_bin;

-- Create "role_routes" table
CREATE TABLE `role_routes` (
  `role_id` bigint NOT NULL,
  `route_id` bigint NOT NULL,
  PRIMARY KEY (`role_id`, `route_id`),
  INDEX `role_routes_route_id` (`route_id`),
  CONSTRAINT `role_routes_role_id` FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`) ON UPDATE NO ACTION ON DELETE CASCADE,
  CONSTRAINT `role_routes_route_id` FOREIGN KEY (`route_id`) REFERENCES `routes` (`id`) ON UPDATE NO ACTION ON DELETE CASCADE
) CHARSET utf8mb4 COLLATE utf8mb4_bin;