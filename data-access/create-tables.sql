DROP TABLE IF EXISTS student;
CREATE TABLE student (
  id         INT AUTO_INCREMENT NOT NULL,
  student_name     VARCHAR(128) NOT NULL,
  PRIMARY KEY (`id`)
);

INSERT INTO student
  (student_name)
VALUES
    ('Bartek'),
    ('Maciek'),
    ('Piotrek');