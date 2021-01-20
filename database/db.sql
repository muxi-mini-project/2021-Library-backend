create database if not exists library;

use library;
    
create table users (
	user_id BIGINT primary key not null AUTO_INCREMENT,
	user_name varchar(50) not null,
	motto TEXT not null,
	password varchar(16) not null,
	user_picture varchar(100) not null
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

create table books (
	books_id BIGINT primary key not null AUTO_INCREMENT,
	books_name varchar(50) not null,
	books_picture varchar(50) not null,
	author varchar(20) not null,
	books_content text not null,
	click  int not null,
	kinds_id BIGINT not null
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

create table digest (
	Digest_id BIGINT primary key not null AUTO_INCREMENT,
	Digest_name varchar(50) not null,
	Digest_content text not null,
	Digest_idea text not null,
	books_id BIGINT not null,
        user_id BIGINT not null,
        DigestKinds_id BIGINT not null,
	locken BOOLEAN not null,
	time timestamp not null default current_timestamp on update current_timestamp
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

create table digestKinds (
	digestkinds_id BIGINT primary key not null AUTO_INCREMENT,
	digestkinds_name char(10) not null
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

create table BooksKinds (
        BooksKinds_id BIGINT PRIMARY KEY  not null AUTO_INCREMENT,
        BooksKinds_name varchar(100) not null,
	BooksKinds_picture char(50) not null,
	BooksKinds_number BIGINT not null
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

create table UtoB (
	id1 BIGINT PRIMARY KEY not null AUTO_INCREMENT,
	user_id BIGINT  not null,
	Books_id BIGINT not null
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


	
