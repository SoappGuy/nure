INSERT INTO
    MedicalCard (
        student_ID,
        weight,
        height,
        health_group,
        blood_group,
        rh_factor,
        last_inspection_date,
        next_inspection_date
    )
VALUES
    /* sql-formatter-disable */
    (1, 45.2, 156.5, '1', 'I', '+', '2024-09-10', '2025-03-10'), -- Олександр Коваленко
    (2, 47.0, 158.0, '1', 'II', '-', '2024-09-11', '2025-03-11'), -- Марія Петренко
    (3, 50.5, 162.3, '2', 'III', '+', '2024-09-12', '2025-03-12'), -- Дмитро Сидоренко
    (4, 48.7, 159.2, '1', 'II', '+', '2024-09-13', '2025-03-13'), -- Анна Бойко
    (5, 52.3, 163.8, '2', 'IV', '-', '2024-09-14', '2025-03-14'), -- Максим Шевченко
    (6, 46.8, 157.0, '1', 'I', '+', '2024-09-15', '2025-03-15'), -- Олена Захаренко
    (7, 49.4, 160.5, '1', 'III', '-', '2024-09-16', '2025-03-16'), -- Іван Кравчук
    (8, 47.2, 158.8, '1', 'II', '+', '2024-09-17', '2025-03-17'), -- Катерина Лисенко
    (9, 51.5, 162.0, '2', 'IV', '-', '2024-09-18', '2025-03-18'), -- Артем Гончаренко
    (10, 48.9, 159.5, '1', 'I', '+', '2024-09-19', '2025-03-19'); -- Вікторія Савченко
    /* sql-formatter-enable */
