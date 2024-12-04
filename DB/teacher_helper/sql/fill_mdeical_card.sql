/* sql-formatter-disable */
SET @student_1 = ( SELECT student_ID FROM Student WHERE lastname = 'Коваленко');
SET @student_2 = ( SELECT student_ID FROM Student WHERE lastname = 'Шевченко');
SET @student_3 = ( SELECT student_ID FROM Student WHERE lastname = 'Бондаренко');
SET @student_4 = ( SELECT student_ID FROM Student WHERE lastname = 'Ткаченко');
SET @student_5 = ( SELECT student_ID FROM Student WHERE lastname = 'Мельник');
/* sql-formatter-enable */
INSERT INTO
    MedicalCard (
        weight,
        height,
        health_group,
        blood_group,
        rh_factor,
        last_inspection_date,
        next_inspection_date,
        note,
        student_ID
    )
VALUES
    /* sql-formatter-disable */
    ( 50.0, 170.0, '2', 'I', '+', '2020-01-01', '2021-01-01', 'Здоровий', @student_1 ),
    ( 60.0, 180.0, '3', 'II', '-', '2020-01-01', '2021-01-01', 'Здоровий', @student_2 ),
    ( 70.0, 190.0, '4', 'III', '+', '2020-01-01', '2021-01-01', 'Здоровий', @student_3 ),
    ( 80.0, 200.0, '5', 'IV', '-', '2020-01-01', '2021-01-01', 'Здоровий', @student_4 ),
    ( 90.0, 210.0, '1', 'I', '+', '2020-01-01', '2021-01-01', 'Здоровий', @student_5 );
    /* sql-formatter-enable */
