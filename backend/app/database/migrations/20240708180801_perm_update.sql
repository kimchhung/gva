-- Modify "permissions" table
ALTER TABLE `permissions` MODIFY COLUMN `order` bigint NULL DEFAULT 0, ADD COLUMN `type` enum('dynamic','static') NULL DEFAULT "dynamic";
