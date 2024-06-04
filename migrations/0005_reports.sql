-- +goose Up
    CREATE TABLE IF NOT EXISTS `reports` (
        `id` int(11) unsigned NOT NULL AUTO_INCREMENT PRIMARY KEY,
        `senderID` int(11) unsigned,
        `label` varchar(256) DEFAULT NULL,
        `description` text DEFAULT NULL,
        CONSTRAINT `reports_r_1` FOREIGN KEY (`senderID`) REFERENCES `volunteers` (`id`) ON DELETE CASCADE
    )ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- +goose Down
DROP TABLE reports