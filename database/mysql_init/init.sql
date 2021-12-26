CREATE DATABASE main CHARACTER SET utf8mb4;

DROP TABLE IF EXISTS posts;

CREATE TABLE posts (
  `id` INT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(255),
  `code` TEXT,
  `description` TEXT,
  `user_id`; INT,
  PRIMARY KEY (`id`)
);

INSERT INTO posts (id, name, code, description, user_id) VALUES (1, "Sample1", "{}", "description1", 1);
INSERT INTO posts (id, name, code, description, user_id) VALUES (2, "Sample2", "{}", "description2", 1);
