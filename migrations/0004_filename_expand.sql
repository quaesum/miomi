-- +goose Up
ALTER TABLE photos MODIFY COLUMN `filename` varchar(512);