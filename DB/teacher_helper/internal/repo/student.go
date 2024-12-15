package repo

import (
	"fmt"
	"teacher_helper/internal/model"

	"github.com/jmoiron/sqlx"
)

type StudentRepo struct {
	db *sqlx.DB
}

func NewStudentRepo(db *sqlx.DB) *StudentRepo {
	return &StudentRepo{db: db}
}

func (r *StudentRepo) GetWithParams(params QueryParams) ([]model.Student, error) {
	var order string
	if params.IsDescending {
		order = "DESC"
	} else {
		order = "ASC"
	}

	query := fmt.Sprintf(`
	SELECT 
		* 
	FROM 
		Student 
	WHERE
		firstname LIKE :query OR
		middlename LIKE :query OR
		lastname LIKE :query OR
		gender LIKE :query OR
		birthday LIKE :query OR
		form_of_education LIKE :query OR
		personal_file_number LIKE :query
	ORDER BY %s %s`, params.OrderBy, order)

	rows, err := r.db.NamedQuery(query, params)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var students []model.Student
	for rows.Next() {
		var student model.Student
		err := rows.StructScan(&student)
		if err != nil {
			return nil, err
		}
		students = append(students, student)
	}

	if err := rows.Close(); err != nil {
		return nil, err
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return students, nil
}

func (r *StudentRepo) GetAll() ([]model.Student, error) {
	var students []model.Student
	err := r.db.Select(&students, "SELECT * FROM Student")
	return students, err
}

func (r *StudentRepo) GetByID(id int) (model.Student, error) {
	var student model.Student
	err := r.db.Get(&student, "SELECT * FROM Student WHERE student_ID = ?", id)
	return student, err
}

func (r *StudentRepo) Create(student *model.Student) error {
	result, err := r.db.NamedExec(`
		INSERT INTO Student 
			(firstname, middlename, lastname, gender, birthday, form_of_education, personal_file_number, note)
		VALUES
			(:firstname, :middlename, :lastname, :gender, :birthday, :form_of_education, :personal_file_number, :note)`,
		student,
	)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return fmt.Errorf("Student not created")
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	student.StudentID = id

	return nil
}

func (r *StudentRepo) Update(student *model.Student) error {
	result, err := r.db.NamedExec(`
		UPDATE Student
		SET
			firstname = :firstname,
			middlename = :middlename,
			lastname = :lastname,
			gender = :gender,
			birthday = :birthday,
			form_of_education = :form_of_education,
			personal_file_number = :personal_file_number,
			note = :note
		WHERE
			student_ID = :student_ID
		`,
		student,
	)

	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return fmt.Errorf("Student not updated")
	}

	return nil
}

func (r *StudentRepo) Delete(id int) error {
	result, err := r.db.Exec("DELETE FROM Student WHERE student_ID = ?", id)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return fmt.Errorf("Student with ID %d not found", id)
	}

	return nil
}

func (r *StudentRepo) GetConnections(id int) ([]model.FamilyConnection, error) {
	rows, err := r.db.Query(`
	SELECT
		family_connection_ID,
		kinship,
		Caretaker.caretaker_ID,
		Caretaker.firstname,
		Caretaker.middlename,
		Caretaker.lastname,
		Caretaker.phone,
		Caretaker.email,
		Student.student_ID,
		Student.firstname,
		Student.middlename,
		Student.lastname,
		Student.gender,
		Student.birthday,
		Student.form_of_education,
		Student.personal_file_number,
		Student.note
	FROM
		FamilyConnection
	JOIN
		Caretaker ON Caretaker.caretaker_ID = FamilyConnection.caretaker_ID
	JOIN
		Student ON Student.student_ID = FamilyConnection.student_ID
	WHERE
		FamilyConnection.student_ID = ?`,
		id,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var connections []model.FamilyConnection
	for rows.Next() {
		var i model.FamilyConnection
		if err := rows.Scan(
			&i.FamilyConnectionID,
			&i.Kinship,
			&i.Caretaker.CaretakerID,
			&i.Caretaker.Firstname,
			&i.Caretaker.Middlename,
			&i.Caretaker.Lastname,
			&i.Caretaker.Phone,
			&i.Caretaker.Email,
			&i.Student.StudentID,
			&i.Student.Firstname,
			&i.Student.Middlename,
			&i.Student.Lastname,
			&i.Student.Gender,
			&i.Student.Birthday,
			&i.Student.FormOfEducation,
			&i.Student.PersonalFileNumber,
			&i.Student.Note,
		); err != nil {
			return nil, err
		}
		connections = append(connections, i)
	}

	if err := rows.Close(); err != nil {
		return nil, err
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return connections, nil
}

func (r *StudentRepo) GetConnectionByID(id int) (*model.FamilyConnection, error) {
	qeury := `
	SELECT
		family_connection_id,
		kinship,
		Caretaker.caretaker_id,
		Caretaker.firstname,
		Caretaker.middlename,
		Caretaker.lastname,
		Caretaker.phone,
		Caretaker.email,
		Student.student_id,
		Student.firstname,
		Student.middlename,
		Student.lastname,
		Student.gender, 
		Student.birthday, 
		Student.form_of_education, 
		Student.personal_file_number, 
		Student.note
	FROM
		FamilyConnection
	JOIN
		Caretaker ON FamilyConnection.caretaker_id = Caretaker.caretaker_id
	JOIN
		Student ON FamilyConnection.student_id = Student.student_id
	WHERE
		family_connection_id = ?
	`

	rows, err := r.db.Query(qeury, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var (
		familyConnections []model.FamilyConnection
	)
	for rows.Next() {
		var i model.FamilyConnection
		if err := rows.Scan(
			&i.FamilyConnectionID,
			&i.Kinship,
			&i.Caretaker.CaretakerID,
			&i.Caretaker.Firstname,
			&i.Caretaker.Middlename,
			&i.Caretaker.Lastname,
			&i.Caretaker.Phone,
			&i.Caretaker.Email,
			&i.Student.StudentID,
			&i.Student.Firstname,
			&i.Student.Middlename,
			&i.Student.Lastname,
			&i.Student.Gender,
			&i.Student.Birthday,
			&i.Student.FormOfEducation,
			&i.Student.PersonalFileNumber,
			&i.Student.Note,
		); err != nil {
			return nil, err
		}
		familyConnections = append(familyConnections, i)
	}

	if err := rows.Close(); err != nil {
		return nil, err
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	if len(familyConnections) != 1 {
		return nil, fmt.Errorf("Expected 1 Connection, but got %d instead", len(familyConnections))
	}

	return &familyConnections[0], nil
}

func (r *StudentRepo) UpdateConnection(parent *model.FamilyConnectionWithIDs) error {
	result, err := r.db.NamedExec(
		`UPDATE FamilyConnection 
		SET 
			kinship = :kinship, 
			student_ID = :student_ID, 
			caretaker_ID = :caretaker_ID 
		WHERE 
			family_connection_ID = :family_connection_ID
		`,
		parent,
	)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return fmt.Errorf("FamilyConnection not updated")
	}

	return nil
}

func (r *StudentRepo) DeleteConnection(id int) error {
	result, err := r.db.Exec("DELETE FROM FamilyConnection WHERE family_connection_ID = ?", id)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return fmt.Errorf("FamilyConnection with id %d not found", id)
	}

	return nil
}

func (r *StudentRepo) CreateConnection(parent *model.FamilyConnectionWithIDs) (int64, error) {
	result, err := r.db.NamedExec(
		`INSERT INTO FamilyConnection 
			(student_ID, caretaker_ID, kinship) 
		VALUES 
			(:student_ID, :caretaker_ID, :kinship)`,
		parent,
	)
	if err != nil {
		return -1, err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return -1, err
	}

	if rows == 0 {
		return -1, fmt.Errorf("FamilyConnection not created")
	}

	id, err := result.LastInsertId()
	if err != nil {
		return -1, err
	}

	return id, nil
}

func (r *StudentRepo) GetStatsForStudent(id int) ([]model.StudentStats, error) {
	rows, err := r.db.NamedQuery(`
		SELECT
			s.student_ID,
			subj.subject_ID,
			subj.title as subject_title,
			COUNT(DISTINCT l.lesson_ID) AS total_lessons,
			COUNT(
				DISTINCT CASE
					WHEN a.attendance = 'Відсутній' THEN l.lesson_ID
				END
			) AS visits,
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
		s.student_ID = :student_ID
		GROUP BY
			s.student_ID,
			subj.subject_ID
		ORDER BY
			subj.title
		`,
		map[string]interface{}{"student_ID": id},
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var stats []model.StudentStats
	for rows.Next() {
		var i model.StudentStats
		if err := rows.StructScan(&i); err != nil {
			return nil, err
		}
		stats = append(stats, i)
	}

	if err := rows.Close(); err != nil {
		return nil, err
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return stats, nil
}
