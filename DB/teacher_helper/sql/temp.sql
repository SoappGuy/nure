SELECT
    firstname,
    lastname,
    COUNT(mark)
FROM
    Student,
    Mark
WHERE
    Student.student_ID = Mark.student_ID
GROUP BY
    firstname,
    lastname
HAVING
    COUNT(mark) >= ALL (
        SELECT
            COUNT(mark)
        FROM
            Student,
            Mark
        WHERE
            Student.student_ID = Mark.student_ID
        GROUP BY
            firstname,
            lastname
    );
