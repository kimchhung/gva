-- Modify "routes" table
ALTER TABLE
  `routes`
MODIFY
  COLUMN `type` enum('cata_log', 'menu', 'button', 'external_link') NOT NULL DEFAULT "cata_log";