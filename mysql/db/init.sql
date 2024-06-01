-- DROP DATABASE IF EXISTS sample_db;
-- CREATE DATABASE sample_db;
SET CHARACTER SET utf8mb4;
DROP TABLE IF EXISTS users;

CREATE TABLE users (
  `id` INT NOT NULL AUTO_INCREMENT,
  `user_id` INT,
  `faculty` VARCHAR(255),
  `department` VARCHAR(255),
  `major` VARCHAR(255),
  `course` VARCHAR(255),
  `grade` VARCHAR(255),
  `name` VARCHAR(255),
  `email` VARCHAR(255),
  `password` VARCHAR(255),
  PRIMARY KEY (`id`)
);



INSERT INTO users VALUES (1, 1000, "経営情報", "経営情報", "経営情報システム", "学士", 3, "山田太郎", "yamada@example.com", "yamada");

