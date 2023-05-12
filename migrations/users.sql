CREATE TABLE IF NOT EXISTS users(
    id int(11) unsigned NOT NULL AUTO_INCREMENT,
    firstName varchar(35) DEFAULT NULL,
    lastName varchar(35) DEFAULT NULL,
    password varchar(255) DEFAULT NULL,
    companyID int(11) unsigned DEFAULT NULL,
    companyName varchar(50) DEFAULT NULL,
    createdAt int(11) DEFAULT NULL,
    PRIMARY KEY (id)
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8;
