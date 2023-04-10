CREATE DATABASE IF NOT EXISTS dental_clinic_db;

USE dental_clinic_db;

CREATE TABLE IF NOT EXISTS dentist (
  id INT(11) NOT NULL AUTO_INCREMENT,
  name VARCHAR(50) NOT NULL,
  last_name VARCHAR(50) NOT NULL,
  license VARCHAR(50) NOT NULL,
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


INSERT INTO dentist (name, last_name, license) VALUES
  ("Juan", "Pérez", "12345"),
  ("María", "Gómez", "67890"),
  ("Carlos", "Hernández", "13579"),
  ("Laura", "García", "24680"),
  ("Pedro", "Rodríguez", "97531"),
  ("Ana", "Martínez", "86420"),
  ("Jorge", "González", "75319"),
  ("Marcela", "López", "24681"),
  ("Luis", "Díaz", "86421"),
  ("Marta", "Sánchez", "75310");

CREATE TABLE IF NOT EXISTS patient (
  id INT(11) NOT NULL AUTO_INCREMENT,
  name VARCHAR(50) NOT NULL,
  last_name VARCHAR(50) NOT NULL,
  domicilio VARCHAR(50) NOT NULL,
  dni INT(11) NOT NULL,
  email VARCHAR(50) NOT NULL,
  admission_date DATE NOT NULL,
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

INSERT INTO patient (name, last_name, domicilio, dni, email, admission_date) VALUES
  ("Ana", "García", "Calle Falsa 123", 12345678, "ana.garcia@gmail.com", "2022-01-15"),
  ("Pedro", "Martínez", "Avenida Siempreviva 456", 23456789, "pedro.martinez@yahoo.com", "2022-02-01"),
  ("Sofía", "López", "Calle Falsa 456", 34567890, "sofia.lopez@hotmail.com", "2022-03-10"),
  ("Carlos", "González", "Calle Real 789", 45678901, "carlos.gonzalez@gmail.com", "2022-03-15"),
  ("Laura", "Fernández", "Calle Mayor 1011", 56789012, "laura.fernandez@yahoo.com", "2022-04-05"),
  ("Pablo", "Sánchez", "Avenida del Sol 1213", 67890123, "pablo.sanchez@hotmail.com", "2022-04-10"),
  ("Lucía", "Romero", "Calle del Prado 1415", 78901234, "lucia.romero@gmail.com", "2022-05-01"),
  ("Miguel", "Gómez", "Calle Nueva 1617", 89012345, "miguel.gomez@yahoo.com", "2022-05-15"),
  ("Elena", "Hernández", "Avenida de la Libertad 1819", 90123456, "elena.hernandez@hotmail.com", "2022-06-01"),
  ("María", "Jiménez", "Calle Mayor 2021", 12345679, "maria.jimenez@gmail.com", "2022-06-15");

CREATE TABLE IF NOT EXISTS appointment (
  id INT(11) NOT NULL AUTO_INCREMENT,
  date DATE NOT NULL,
  hour TIME NOT NULL,
  patient_id INT(11) NOT NULL,
  dentist_id INT(11) NOT NULL,
  PRIMARY KEY (id),
  FOREIGN KEY (patient_id) REFERENCES patient(id),
  FOREIGN KEY (dentist_id) REFERENCES dentist(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

INSERT INTO appointment (date, hour, patient_id, dentist_id) VALUES
  ("2022-03-15", "10:00:00", 1, 1),
  ("2022-03-16", "09:00:00", 2, 3),
  ("2022-03-17", "11:00:00", 4, 5),
  ("2022-03-18", "15:00:00", 6, 7),
  ("2022-03-19", "14:00:00", 8, 9);