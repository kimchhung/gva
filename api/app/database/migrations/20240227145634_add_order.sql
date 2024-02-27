-- Modify "permissions" table
ALTER TABLE `permissions` DROP COLUMN `admin_permissions`, DROP FOREIGN KEY `permissions_admins_permissions`;
