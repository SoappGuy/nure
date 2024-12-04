DROP TABLE IF EXISTS Mark;


DROP TABLE IF EXISTS Lesson;


DROP TABLE IF EXISTS Subject;


DROP TABLE IF EXISTS FamilyConnection;


DROP TABLE IF EXISTS Caretaker;


DROP TABLE IF EXISTS Student;


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
