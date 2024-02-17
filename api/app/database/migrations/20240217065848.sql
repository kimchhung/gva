-- Modify "admins" table
ALTER TABLE `admins` DROP COLUMN `name`, ADD COLUMN `username` varchar(255) NOT NULL, ADD COLUMN `password` varchar(255) NOT NULL;
