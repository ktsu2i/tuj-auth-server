
-- +migrate Up
CREATE TABLE IF NOT EXISTS `users` (
	`tuid` INT UNSIGNED NOT NULL PRIMARY KEY,
	`email` VARCHAR(255) NOT NULL UNIQUE,
	`hashed_password` VARCHAR(255) NOT NULL,
	`first_name` VARCHAR(255) NOT NULL,
	`middle_name` VARCHAR(255),
	`last_name` VARCHAR(255) NOT NULL,
	`departments` JSON NOT NULL,
	`is_admin` BOOLEAN NOT NULL DEFAULT FALSE,
	`is_coordinator` BOOLEAN NOT NULL DEFAULT FALSE,
	`is_faculty` BOOLEAN NOT NULL DEFAULT FALSE,
	`is_staff` BOOLEAN NOT NULL DEFAULT FALSE,
	`created_at` DATETIME(6) DEFAULT NULL,
	`updated_at` DATETIME(6) DEFAULT NULL
);

-- +migrate Down
DROP TABLE IF EXISTS `users`;