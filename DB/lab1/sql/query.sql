-- name: ListStudents :many
SELECT
    *
FROM
    Student
ORDER BY
    lastname;
