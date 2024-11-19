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


CREATE TABLE Subject_ (
    subject_ID INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(32) NOT NULL,
    descript TEXT,
    number_of_hours int NOT NULL
);


CREATE TABLE Lesson (
    lesson_ID INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    topic TEXT NOT NULL,
    start_date DATE NOT NULL,
    start_time TIME NOT NULL,
    subject_ID INT NOT NULL,
    FOREIGN KEY (subject_ID) REFERENCES Subject_.subject_ID
);


CREATE TABLE Mark (
    mark_ID INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    mark INT NOT NULL,
    student_ID INT NOT NULL,
    lesson_ID INT NOT NULL,
    FOREIGN KEY (student_ID) REFERENCES Student.student_ID,
    FOREIGN KEY (lesson_ID) REFERENCES Lesson.lesson_ID
);
