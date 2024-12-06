/* sql-formatter-disable */
DROP TABLE IF EXISTS StudentAdress;
DROP TABLE IF EXISTS Adress;
DROP TABLE IF EXISTS Privilege;
DROP TABLE IF EXISTS MedicalCard;
DROP TABLE IF EXISTS Attendance;
DROP TABLE IF EXISTS Mark;
DROP TABLE IF EXISTS Lesson;
DROP TABLE IF EXISTS Subject;
DROP TABLE IF EXISTS FamilyConnection;
DROP TABLE IF EXISTS Caretaker;
DROP TABLE IF EXISTS Student;
/* sql-formatter-enable */
CREATE TABLE Student (
    student_ID INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    firstname VARCHAR(16) NOT NULL,
    middlename VARCHAR(16) NOT NULL,
    lastname VARCHAR(16) NOT NULL,
    gender ENUM("M", "F") NOT NULL,
    birthday DATE NOT NULL,
    form_of_education ENUM("Денна", "Вечірня", "Домашня") NOT NULL,
    personal_file_number VARCHAR(16) NOT NULL,
    note TEXT
);


CREATE TABLE Caretaker (
    caretaker_ID INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    firstname VARCHAR(16) NOT NULL,
    middlename VARCHAR(16) NOT NULL,
    lastname VARCHAR(16) NOT NULL,
    phone VARCHAR(16) NOT NULL,
    email VARCHAR(32) NOT NULL
);


CREATE TABLE FamilyConnection (
    family_connection_ID INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    kinship ENUM(
        'Батько',
        'Мати',
        'Вітчим',
        'Мачеха',
        'Дідусь',
        'Бабуся',
        'Тітка',
        'Дядько',
        'Брат',
        'Сестра',
        'Інше'
    ) NOT NULL,
    caretaker_ID INT NOT NULL,
    student_ID INT NOT NULL,
    FOREIGN KEY (caretaker_ID) REFERENCES Caretaker (caretaker_ID) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (student_ID) REFERENCES Student (student_ID) ON DELETE CASCADE ON UPDATE CASCADE
);


CREATE TABLE Subject (
    subject_ID INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(32) NOT NULL,
    description TEXT,
    number_of_hours int NOT NULL
);


CREATE TABLE Lesson (
    lesson_ID INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    topic TEXT NOT NULL,
    start_date DATE NOT NULL,
    start_time TIME NOT NULL,
    lesson_number TINYINT NOT NULL,
    subject_ID INT NOT NULL,
    FOREIGN KEY (subject_ID) REFERENCES Subject (subject_ID) ON DELETE CASCADE ON UPDATE CASCADE
);


CREATE TABLE Mark (
    mark_ID INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    mark INT NOT NULL,
    student_ID INT NOT NULL,
    lesson_ID INT NOT NULL,
    FOREIGN KEY (student_ID) REFERENCES Student (student_ID) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (lesson_ID) REFERENCES Lesson (lesson_ID) ON DELETE CASCADE ON UPDATE CASCADE
);


CREATE TABLE Attendance (
    attendance_ID INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    attendance ENUM('Відсутній', 'Присутній') NOT NULL,
    reason TEXT,
    student_ID INT NOT NULL,
    lesson_ID INT NOT NULL,
    FOREIGN KEY (student_ID) REFERENCES Student (student_ID) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (lesson_ID) REFERENCES Lesson (lesson_ID) ON DELETE CASCADE ON UPDATE CASCADE
);


CREATE TABLE MedicalCard (
    medical_card_ID INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    student_ID INT NOT NULL,
    weight FLOAT NOT NULL,
    height FLOAT NOT NULL,
    health_group ENUM('1', '2', '3', '4', '5') NOT NULL,
    blood_group ENUM('I', 'II', 'III', 'IV') NOT NULL,
    rh_factor ENUM('+', '-') NOT NULL,
    last_inspection_date DATE NOT NULL,
    next_inspection_date DATE NOT NULL,
    note TEXT,
    FOREIGN KEY (student_ID) REFERENCES Student (student_ID) ON DELETE CASCADE ON UPDATE CASCADE
);


CREATE TABLE Privilege (
    privilege_ID INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    student_ID INT NOT NULL,
    type VARCHAR(32) NOT NULL,
    description TEXT,
    expiration_date DATE,
    FOREIGN KEY (student_ID) REFERENCES Student (student_ID) ON DELETE CASCADE ON UPDATE CASCADE
);


CREATE TABLE Adress (
    adress_ID INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    country VARCHAR(32) NOT NULL,
    region VARCHAR(32) NOT NULL,
    city VARCHAR(32) NOT NULL,
    street VARCHAR(32) NOT NULL,
    biulding VARCHAR(32) NOT NULL,
    apartment VARCHAR(32) NOT NULL,
    postal_code VARCHAR(32) NOT NULL
);


CREATE TABLE StudentAdress (
    student_adress_ID INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    student_ID INT NOT NULL,
    adress_ID INT NOT NULL,
    type ENUM('Постійна', 'Тимчасова', 'Реєстрації') NOT NULL,
    FOREIGN KEY (student_ID) REFERENCES Student (student_ID) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (adress_ID) REFERENCES Adress (adress_ID) ON DELETE CASCADE ON UPDATE CASCADE
);
