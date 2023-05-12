CREATE TABLE IF NOT EXISTS animals(
    id int(10) unsigned NOT NULL PRIMARY KEY ,
    animal varchar(40) DEFAULT NULL,
    pet_name varchar(20) DEFAULT NULL,
    gender varchar(10) DEFAULT NULL,
    location varchar(30) DEFAULT NULL,
    age smallint(5) unsigned DEFAULT NULL,
    photo varchar(255) DEFAULT NULL,
    createdAt varchar(50) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;