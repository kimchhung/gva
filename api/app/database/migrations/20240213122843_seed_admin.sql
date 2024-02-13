-- Add the initial admin to the database.
INSERT INTO `admins` (`created_at`, `updated_at`, `name`, `display_name`) VALUES ({{ TIME_VALUE }}, {{ TIME_VALUE }}, 'admin', 'super admin');
