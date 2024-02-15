-- Add the initial admin to the database.
INSERT INTO `roles` (`created_at`, `updated_at`, `name`) VALUES (NOW(), NOW(), 'Super Admin');
INSERT INTO `admins` (`created_at`, `updated_at`, `name`, `display_name`) VALUES (NOW(), NOW(), 'admin', 'ADMIN');
INSERT INTO `admin_roles` (`admin_id`, `role_id`) VALUES (1, 1) ON DUPLICATE KEY UPDATE `admin_id` = `admin_roles`.`admin_id`, `role_id` = `admin_roles`.`role_id`;
