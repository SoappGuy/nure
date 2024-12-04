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
