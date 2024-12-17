SELECT
    Student.student_ID,
    CONCAT_WS(
        ' ',
        Student.lastname,
        Student.firstname,
        Student.middlename
    ) AS fullname,
    MedicalCard.next_inspection_date
FROM
    MedicalCard,
    Student
WHERE
    MedicalCard.student_ID = Student.student_ID
ORDER BY
    MedicalCard.next_inspection_date;


SELECT
    Student.student_ID,
    CONCAT_WS(
        ' ',
        Student.lastname,
        Student.firstname,
        Student.middlename
    ) AS fullname,
    AVG(Mark.mark) AS avg_mark
FROM
    Mark
    JOIN Student ON Mark.student_ID = Student.student_ID
GROUP BY
    Student.student_ID
ORDER BY
    AVG(Mark.mark) DESC;


SELECT
    Student.student_ID,
    CONCAT_WS(
        ' ',
        Student.lastname,
        Student.firstname,
        Student.middlename
    ) AS fullname,
    Privilege.type,
    Privilege.expiration_date
FROM
    Privilege
    JOIN Student ON Privilege.student_ID = Student.student_ID;


SELECT
    form_of_education,
    COUNT(*) AS COUNT
FROM
    Student
GROUP BY
    form_of_education;


SELECT
    gender,
    COUNT(*) AS COUNT
FROM
    Student
GROUP BY
    gender;


SELECT
    MedicalCard.health_group,
    COUNT(*) AS COUNT
FROM
    MedicalCard
GROUP BY
    MedicalCard.health_group;


SELECT
    MedicalCard.blood_group,
    MedicalCard.rh_factor,
    COUNT(*) AS COUNT
FROM
    MedicalCard
GROUP BY
    MedicalCard.blood_group,
    MedicalCard.rh_factor
ORDER BY
    MedicalCard.blood_group,
    MedicalCard.rh_factor;
