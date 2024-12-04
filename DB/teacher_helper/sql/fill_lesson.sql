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
