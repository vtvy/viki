create database viki;
use viki;

create table users (
	id 			int AUTO_INCREMENT,
    username 	varchar(50) NOT NULL unique,
    pass		varchar(100) NOT NULL,
	created_at timestamp DEFAULT CURRENT_TIMESTAMP,
	updated_at timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- INSERT
insert into users(username, pass) values("admin", "1234");