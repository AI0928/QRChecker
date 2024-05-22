-- DROP DATABASE IF EXISTS sample_db;
-- CREATE DATABASE sample_db;
SET CHARACTER SET utf8;
DROP TABLE IF EXISTS users;

CREATE TABLE users (
  `id` INT NOT NULL AUTO_INCREMENT,
  `user_id` INT,
  `学部` VARCHAR(255),
  `学科` VARCHAR(255),
  `専攻` VARCHAR(255),
  `課程` VARCHAR(255),
  `学年` VARCHAR(255),
  `name` VARCHAR(255),
  `email` VARCHAR(255),
  `password` VARCHAR(255),
  PRIMARY KEY (`id`)
);



INSERT INTO users VALUES (1, 1000, "経営情報", "経営情報", "経営情報システム", "学士", 3, "山田太郎", "yamada@example.com", "yamada");

