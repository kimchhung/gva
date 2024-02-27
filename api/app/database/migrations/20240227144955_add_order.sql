-- Modify "permissions" table
ALTER TABLE `permissions` ADD COLUMN `order` bigint NOT NULL, ADD COLUMN `admin_permissions` bigint NULL, ADD INDEX `permissions_admins_permissions` (`admin_permissions`), ADD CONSTRAINT `permissions_admins_permissions` FOREIGN KEY (`admin_permissions`) REFERENCES `admins` (`id`) ON UPDATE NO ACTION ON DELETE SET NULL;
