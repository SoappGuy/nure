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


DELETE FROM Student;


INSERT INTO
    Student (
        firstname,
        middlename,
        lastname,
        gender,
        birthday,
        form_of_education,
        personal_file_number,
        note
    )
VALUES
    /* sql-formatter-disable */
    ('Олександр', 'Петрович', 'Коваленко', 'M', '2009-05-15', 'Денна', 'AA-01', NULL),
    ('Марія', 'Іванівна', 'Шевченко', 'F', '2011-03-22', 'Денна', 'AA-02', 'Учасниця олімпіад'),
    ('Дмитро', 'Олегович', 'Мельник', 'M', '2011-11-30', 'Домашня', 'AБ-01', 'Спортсмен'),
    ('Софія', 'Андріївна', 'Бондаренко', 'F', '2010-07-19', 'Денна', 'ББ-01', 'Відмінниця'),
    ('Максим', 'Васильович', 'Ткаченко', 'M', '2011-09-05', 'Денна', 'ББ-02', 'Учасник наукового гуртка');
    /* sql-formatter-enable */
DELETE FROM Caretaker;


INSERT INTO
    Caretaker (firstname, middlename, lastname, phone, email)
VALUES
    /* sql-formatter-disable */
    ('Петро', 'Васильович', 'Коваленко', '+380671234567', 'pvk@example.com'),
    ('Ірина', 'Миколаївна', 'Шевченко', '+380671234568', 'ims@example.com'),
    ('Олег', 'Вікторович', 'Мельник', '+380671234569', 'ovm@example.com'),
    ('Андрій', 'Васильович', 'Бондаренко', '+380671234570', 'avb@example.com'),
    ('Василь', 'Миколайович', 'Ткаченко', '+380671234571', 'vmt@example.com');
    /* sql-formatter-enable */
    /* sql-formatter-disable */
    SET @caretaker_1 = ( SELECT caretaker_ID FROM Caretaker WHERE lastname = 'Коваленко');
    SET @caretaker_2 = ( SELECT caretaker_ID FROM Caretaker WHERE lastname = 'Шевченко');
    SET @caretaker_3 = ( SELECT caretaker_ID FROM Caretaker WHERE lastname = 'Бондаренко');
    SET @caretaker_4 = ( SELECT caretaker_ID FROM Caretaker WHERE lastname = 'Ткаченко');
    SET @caretaker_5 = ( SELECT caretaker_ID FROM Caretaker WHERE lastname = 'Мельник');
    
    SET @student_1 = ( SELECT student_ID FROM Student WHERE lastname = 'Коваленко');
    SET @student_2 = ( SELECT student_ID FROM Student WHERE lastname = 'Шевченко');
    SET @student_3 = ( SELECT student_ID FROM Student WHERE lastname = 'Бондаренко');
    SET @student_4 = ( SELECT student_ID FROM Student WHERE lastname = 'Ткаченко');
    SET @student_5 = ( SELECT student_ID FROM Student WHERE lastname = 'Мельник');
    /* sql-formatter-enable */
DELETE FROM FamilyConnection;


INSERT INTO
    FamilyConnection (kinship, caretaker_ID, student_ID)
VALUES
    ('Батько', @caretaker_1, @student_1),
    ('Мати', @caretaker_2, @student_2),
    ('Вітчим', @caretaker_3, @student_3),
    ('Брат', @caretaker_4, @student_4),
    ('Дідусь', @caretaker_5, @student_5);


DELETE FROM Subject;


INSERT INTO
    Subject (title, description, number_of_hours)
VALUES
    (
        'Математика',
        'Наука, яка первісно виникла як один з напрямків пошуку істини у сфері просторових відношень і обчислень, для практичних потреб людини рахувати, обчислювати, вимірювати, досліджувати форми та рух фізичних тіл.',
        68
    ),
    (
        'Українська Мова',
        'Національна мова українців.',
        85
    ),
    (
        'Інформатика',
        'Наука про інформацію, методи та засоби її опрацювання, у тому числі за допомогою обчислювальних систем.',
        34
    ),
    (
        'Фізика',
        'Природнича наука, яка досліджує загальні властивості матерії та явищ у ній, а також виявляє загальні закони, які керують цими явищами.',
        51
    ),
    (
        'Образотворче Мистецтво',
        'Мистецтво відображення сущого у вигляді різних образів, зокрема таких як художні образи на площині та в просторі.',
        17
    );


/* sql-formatter-disable */
SET @subject_1 = ( SELECT subject_ID FROM Subject WHERE title = 'Математика');
SET @subject_2 = ( SELECT subject_ID FROM Subject WHERE title = 'Українська Мова');
SET @subject_3 = ( SELECT subject_ID FROM Subject WHERE title = 'Інформатика');
SET @subject_4 = ( SELECT subject_ID FROM Subject WHERE title = 'Фізика');
SET @subject_5 = ( SELECT subject_ID FROM Subject WHERE title = 'Образотворче Мистецтво');
/* sql-formatter-enable */
DELETE FROM Lesson;


INSERT INTO
    Lesson (
        topic,
        start_date,
        start_time,
        lesson_number,
        subject_id
    )
VALUES
    /* sql-formatter-disable */
    ('Звичайні дроби та дії з ними', '2024-10-24', '09:00:00', 1, @subject_1),
    ('Десяткові дроби та дії з ними', '2024-10-25', '10:00:00', 2, @subject_1),
    ('Мішані дроби та дії з ними', '2024-10-26', '11:00:00', 3, @subject_1),
    ('Рівняння', '2024-10-27', '09:30:00',2 , @subject_1),
    ("Від'ємні числа", '2024-10-28', '08:30:00',1 , @subject_1),
    ('Фонетика', '2024-10-24', '12:00:00',4 , @subject_2),
    ('Синтаксис', '2024-10-25', '13:00:00',5 ,  @subject_2),
    ('Морфологія', '2024-10-26', '14:00:00',6 , @subject_2),
    ('Орфографія', '2024-10-27', '15:00:00',7 , @subject_2),
    ('Стилістика', '2024-10-28', '16:00:00',8 , @subject_2),
    ('Основи роботи з Microsoft Word', '2024-10-24', '09:00:00',1 , @subject_3),
    ('Основи роботи з Microsoft Excel', '2024-10-25', '10:00:00',2 , @subject_3),
    ('Основи роботи з Microsoft PowerPoint', '2024-10-26', '11:00:00',3 , @subject_3),
    ('Основи роботи з Microsoft Access', '2024-10-27', '09:30:00',2 , @subject_3),
    ('Основи програмування у Scratch', '2024-10-28', '08:30:00',1 , @subject_3),
    ('Механіка', '2024-10-24', '12:00:00',4 , @subject_4),
    ('Кінематика', '2024-10-25', '13:00:00',5 , @subject_4),
    ('Оптика', '2024-10-26', '14:00:00',6 , @subject_4),
    ('Електрика', '2024-10-27', '15:00:00',7 , @subject_4),
    ('Магнетизм', '2024-10-28', '16:00:00',8 , @subject_4),
    ('Малювання', '2024-10-24', '09:00:00',1 , @subject_5),
    ('Живопис', '2024-10-25', '10:00:00',2 , @subject_5),
    ('Скульптура', '2024-10-26', '11:00:00',3 , @subject_5),
    ('Графіка', '2024-10-27', '09:30:00',2 , @subject_5),
    ('Композиція', '2024-10-28', '08:30:00',1 , @subject_5);
    /* sql-formatter-enable */
    /* sql-formatter-disable */
    SET @student_1 = ( SELECT student_ID FROM Student WHERE firstname = 'Олександр' AND lastname = 'Коваленко');
    SET @student_2 = ( SELECT student_ID FROM Student WHERE firstname = 'Марія' AND lastname = 'Шевченко');
    SET @student_3 = ( SELECT student_ID FROM Student WHERE firstname = 'Максим' AND lastname = 'Ткаченко');
    
    SET @lesson_1 = ( SELECT lesson_ID FROM Lesson WHERE topic = 'Звичайні дроби та дії з ними');
    SET @lesson_2 = ( SELECT lesson_ID FROM Lesson WHERE topic = 'Основи роботи з Microsoft Word');
    SET @lesson_3 = ( SELECT lesson_ID FROM Lesson WHERE topic = 'Скульптура');
    /* sql-formatter-enable */
DELETE FROM Mark;


INSERT INTO
    Mark (mark, student_ID, lesson_ID)
VALUES
    (10, @student_1, @lesson_1),
    (12, @student_1, @lesson_2),
    (8, @student_1, @lesson_3),
    (11, @student_2, @lesson_1),
    (9, @student_2, @lesson_2),
    (9, @student_2, @lesson_3),
    (11, @student_3, @lesson_1),
    (11, @student_3, @lesson_2),
    (8, @student_3, @lesson_3);
