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
