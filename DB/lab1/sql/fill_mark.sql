/* sql-formatter-disable */
SET @student_1 = ( SELECT student_ID FROM Student WHERE firstname = "Олександр" AND lastname = "Коваленко");
SET @student_2 = ( SELECT student_ID FROM Student WHERE firstname = "Марія" AND lastname = "Шевченко");
SET @student_3 = ( SELECT student_ID FROM Student WHERE firstname = "Максим" AND lastname = "Ткаченко");

SET @lesson_1 = ( SELECT lesson_ID FROM Lesson WHERE topic = "Звичайні дроби та дії з ними");
SET @lesson_2 = ( SELECT lesson_ID FROM Lesson WHERE topic = "Основи роботи з Microsoft Word");
SET @lesson_3 = ( SELECT lesson_ID FROM Lesson WHERE topic = "Скульптура");
/* sql-formatter-enable */
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
