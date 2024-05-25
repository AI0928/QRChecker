use sample_db;

DROP TABLE IF EXISTS lectures;

CREATE TABLE lectures (
  `id` INT NOT NULL AUTO_INCREMENT,
  `number` INT,
  `name` VARCHAR(255),
  `teacher` VARCHAR(255),
  `start_time` TIME,
  `lecture_time` TIME,
  PRIMARY KEY (`id`)
);

INSERT INTO lectures VALUES (1, 1000, "授業１", "鈴木一郎" , '12:00:00', '1:30:00');
