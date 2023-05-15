CREATE TABLE IF NOT EXISTS animals(
    id int(10) unsigned NOT NULL PRIMARY KEY ,
    age smallint(5) unsigned DEFAULT NULL,
    name varchar(20) DEFAULT NULL,
    sex varchar(10) DEFAULT NULL,
    type varchar(40) DEFAULT NULL,
    description varchar(255) DEFAULT NULL,
    castrated bool DEFAULT NULL,
    sterilized bool DEFAULT NULL,
    vaccinated bool DEFAULT NULL,
    shelter varchar(30) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;