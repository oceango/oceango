create table users (
    id int(11) UNSIGNED AUTO_INCREMENT,
    name varchar(25),
    telephone varchar(11),
    password varchar(255),
    PRIMARY KEY (`id` )
)ENGINE=InnoDB DEFAULT CHARSET=utf8;