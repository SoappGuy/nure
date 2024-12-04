/* CREATE TABLE Privilege ( */
/*     privilege_ID INT NOT NULL AUTO_INCREMENT PRIMARY KEY, */
/*     student_ID INT NOT NULL, */
/*     type VARCHAR(32) NOT NULL, */
/*     description TEXT, */
/*     expiration_date DATE, */
/*     FOREIGN KEY (student_ID) REFERENCES Student (student_ID) ON DELETE CASCADE ON UPDATE CASCADE */
/* ); */

INSERT INTO Privilege (
    student_ID,
    type,
    description,
    expiration_date
) VALUES
    ( 1, 'Соціальна', 'Знижка на проїзд', '2021-01-01'),
    ( 2, 'Медична', 'Звільнення від фізкультури', '2021-01-01'),
    ( 3, 'Соціальна', 'Знижка на проїзд', '2021-01-01'),
    ( 4, 'Медична', 'Звільнення від фізкультури', '2021-01-01'),
    ( 5, 'Соціальна', 'Знижка на проїзд', '2021-01-01');
