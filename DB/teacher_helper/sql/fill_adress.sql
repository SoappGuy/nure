/* CREATE TABLE Adress ( */
/*     adress_ID INT NOT NULL AUTO_INCREMENT PRIMARY KEY, */
/*     country VARCHAR(32) NOT NULL, */
/*     region VARCHAR(32) NOT NULL, */
/*     city VARCHAR(32) NOT NULL, */
/*     street VARCHAR(32) NOT NULL, */
/*     biulding VARCHAR(32) NOT NULL, */
/*     apartment VARCHAR(32) NOT NULL, */
/*     postal_code VARCHAR(32) NOT NULL */
/* ); */
INSERT INTO
    Adress (
        country,
        region,
        city,
        street,
        biulding,
        apartment,
        postal_code
    )
VALUES
    /* sql-formatter-disable */
    ('Україна', 'Київська', 'Київ', 'Шевченка', '1', '1', '01001'),
    ('Україна', 'Київська', 'Київ', 'Шевченка', '1', '2', '01001'),
    ('Україна', 'Київська', 'Київ', 'Шевченка', '1', '3', '01001'),
    ('Україна', 'Київська', 'Київ', 'Шевченка', '1', '4', '01001'),
    ('Україна', 'Київська', 'Київ', 'Шевченка', '1', '5', '01001');
    /* sql-formatter-enable */
