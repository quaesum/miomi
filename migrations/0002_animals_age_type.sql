-- +goose Up
ALTER TABLE animals ADD COLUMN `ageType` varchar(10);
ALTER TABLE animals MODIFY COLUMN `description` varchar(512);
