CREATE TABLE IF NOT EXISTS volunteers(
    id INT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    firstName varchar(35) DEFAULT NULL,
    lastName varchar(35) DEFAULT NULL,
    password varchar(255) DEFAULT NULL,
    email  VARCHAR(256)  DEFAULT '',
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
    description TEXT DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS news_photos (
    id INT(11) UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    newsID INT(11) UNSIGNED,
    photoID INT(11) UNSIGNED,
    CONSTRAINT mi_news_photos_newsID FOREIGN KEY(newsID) REFERENCES news(id) ON DELETE SET NULL,
    CONSTRAINT mi_news_photos_photoID FOREIGN KEY(photoID) REFERENCES photos(id) ON DELETE SET NULL
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8;

