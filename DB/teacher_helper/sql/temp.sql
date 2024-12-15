SELECT
    subj.subject_ID,
    subj.title,
    COUNT(DISTINCT l.lesson_ID) AS total_lessons,
    COUNT(
        DISTINCT CASE
            WHEN a.attendance = 'Відсутній' THEN l.lesson_ID
        END
    ) AS Visits,
    IFNULL(AVG(m.mark), 0) AS grade
FROM
    Student s
    CROSS JOIN Lesson l
    LEFT JOIN Subject subj ON l.subject_ID = subj.subject_ID
    LEFT JOIN Attendance a ON s.student_ID = a.student_ID
    AND l.lesson_ID = a.lesson_ID
    LEFT JOIN Mark m ON s.student_ID = m.student_ID
    AND l.lesson_ID = m.lesson_ID
WHERE
    s.student_ID = 1
GROUP BY
    s.student_ID,
    subj.subject_ID
ORDER BY
    s.lastname,
    s.firstname,
    subj.title;
