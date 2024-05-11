-- +goose Up
CREATE TABLE IF NOT EXISTS volunteers(
    id INT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    firstName varchar(35) DEFAULT NULL,
    lastName varchar(35) DEFAULT NULL,
    password varchar(255) DEFAULT NULL,
    role varchar(35) DEFAULT NULL,
    email  VARCHAR(256)  DEFAULT '',
    phone varchar(45) DEFAULT NULL,
    deactevated bool DEFAULT FALSE,
    createdAt int(11) DEFAULT NULL
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS animal_shelters(
    id INT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    shelter_name varchar(90) DEFAULT NULL,
    description varchar(340) DEFAULT NULL,
    logo varchar(240) DEFAULT NULL,
    phone varchar(45) DEFAULT NULL,
    adress varchar(240) DEFAULT NULL
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS volunteers_on_shelters (
    id INT(11) UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    volunteerID INT(11) UNSIGNED,
    shelterID INT(11) UNSIGNED,
    CONSTRAINT mi_volunteer_assigned FOREIGN KEY(volunteerID) REFERENCES volunteers(id) ON DELETE SET NULL,
    CONSTRAINT mi_shelter_assigned FOREIGN KEY(shelterID) REFERENCES animal_shelters(id) ON DELETE SET NULL
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS animals(
    id INT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    age TINYINT unsigned DEFAULT NULL,
    name varchar(35) DEFAULT NULL,
    sex TINYINT DEFAULT NULL,
    description varchar(255) DEFAULT NULL,
    sterilized bool DEFAULT NULL,
    vaccinated bool DEFAULT NULL,
    onrainbow bool DEFAULT NULL,
    onhappines bool DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS animals_on_shelters (
    id INT(11) UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    animalID INT(11) UNSIGNED,
    shelterID INT(11) UNSIGNED,
    CONSTRAINT mi_animal_assigned FOREIGN KEY(animalID) REFERENCES animals(id) ON DELETE SET NULL,
    CONSTRAINT mi_animal_shelter_assigned FOREIGN KEY(shelterID) REFERENCES animal_shelters(id) ON DELETE SET NULL
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS animal_types(
    id INT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    name varchar(90) DEFAULT NULL
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS animals_on_types (
    id INT(11) UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    animal_typeID INT(11) UNSIGNED,
    animalID INT(11) UNSIGNED,
    CONSTRAINT mi_animal_types_assigned FOREIGN KEY(animal_typeID) REFERENCES animal_types(id) ON DELETE SET NULL,
    CONSTRAINT mi_types_assigned FOREIGN KEY(animalID) REFERENCES animals(id) ON DELETE SET NULL
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS photos (
    id INT(11) unsigned PRIMARY KEY AUTO_INCREMENT,
    filename varchar(256) DEFAULT NULL,
    origin varchar(256) DEFAULT NULL,
    file_type varchar(10) DEFAULT NULL
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS animals_photos (
    id INT(11) UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    animalID INT(11) UNSIGNED,
    photoID INT(11) UNSIGNED,
    CONSTRAINT mi_animals_photos_animalID FOREIGN KEY(animalID) REFERENCES animals(id) ON DELETE SET NULL,
    CONSTRAINT mi_animals_photos_photoID FOREIGN KEY(photoID) REFERENCES photos(id) ON DELETE SET NULL
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS news (
    id INT(11) unsigned PRIMARY KEY AUTO_INCREMENT,
    label varchar(256) DEFAULT NULL,
    description TEXT DEFAULT NULL,
    created_at int(11) DEFAULT NULL
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS news_photos (
    id INT(11) UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    newsID INT(11) UNSIGNED,
    photoID INT(11) UNSIGNED,
    CONSTRAINT mi_news_photos_newsID FOREIGN KEY(newsID) REFERENCES news(id) ON DELETE SET NULL,
    CONSTRAINT mi_news_photos_photoID FOREIGN KEY(photoID) REFERENCES photos(id) ON DELETE SET NULL
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS services (
    id INT(11) unsigned PRIMARY KEY AUTO_INCREMENT,
    volunteer_id INT(11) unsigned,
    name varchar(256) DEFAULT NULL,
    description TEXT DEFAULT NULL,
    created_at int(11) DEFAULT NULL,
    deleted_at int(11) DEFAULT NULL,
    updated_at int(11) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS services_photos (
    id INT(11) UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    serviceID INT(11) UNSIGNED,
    photoID INT(11) UNSIGNED,
    CONSTRAINT mi_services_photos_serviceID FOREIGN KEY(serviceID) REFERENCES services(id) ON DELETE SET NULL,
    CONSTRAINT mi_services_photos_photoID FOREIGN KEY(photoID) REFERENCES photos(id) ON DELETE SET NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS products (
    id INT(11) unsigned PRIMARY KEY AUTO_INCREMENT,
    name varchar(256),
    description TEXT,
    link varchar(256),
    created_at int(11) DEFAULT NULL,
    deleted_at int(11) DEFAULT NULL,
    updated_at int(11) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS products_photos (
    id INT(11) UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    productID INT(11) UNSIGNED,
    photoLink varchar(256),
    CONSTRAINT mi_products_photos_productID FOREIGN KEY(productID) REFERENCES products(id) ON DELETE SET NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS `passwordRecovery` (
    id int(11) unsigned NOT NULL AUTO_INCREMENT,
    `token` varchar(40) DEFAULT NULL,
    `volunteerID` int(11) unsigned DEFAULT NULL,
    PRIMARY KEY (`id`),
    KEY `volunteerID` (`volunteerID`),
    CONSTRAINT `passwordrecovery_ibfk_1` FOREIGN KEY (`volunteerID`) REFERENCES `volunteers` (`id`) ON DELETE CASCADE

) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS `usersEmail` (
    `email` varchar(255) NOT NULL,
    `volunteerID` int(11) unsigned DEFAULT NULL,
    `isConfirmed` tinyint(1) DEFAULT NULL,
    `createdAt` int(11) DEFAULT NULL,
    `verificationToken` varchar(600) DEFAULT NULL,
    PRIMARY KEY (`email`),
    KEY `volunteerID` (`volunteerID`),
    CONSTRAINT `usersEmail_ibfk_1` FOREIGN KEY (`volunteerID`) REFERENCES `volunteers` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


-- +goose Down
DROP TABLE IF EXISTS passwordRecovery;
DROP TABLE IF EXISTS usersEmail;
DROP TABLE  IF EXISTS products;
DROP TABLE IF EXISTS products_photos;
DROP TABLE  IF EXISTS services;
DROP TABLE IF EXISTS services_photos;
DROP TABLE IF EXISTS news_photos;
DROP TABLE IF EXISTS news;
DROP TABLE IF EXISTS animals_photos;
DROP TABLE IF EXISTS photos;
DROP TABLE IF EXISTS animals_on_types;
DROP TABLE IF EXISTS animal_types;
DROP TABLE IF EXISTS animals_on_shelters;
DROP TABLE IF EXISTS animals;
DROP TABLE IF EXISTS volunteers_on_shelters;
DROP TABLE IF EXISTS animal_shelters;
DROP TABLE IF EXISTS volunteers;

