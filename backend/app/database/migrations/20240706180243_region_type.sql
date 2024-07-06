-- Modify "regions" table
ALTER TABLE `regions` ADD COLUMN `type` enum('continent','country','city','street','any') NOT NULL;
