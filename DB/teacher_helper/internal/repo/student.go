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

type StudentParams struct {
	Query        string `db:"query"`
	OrderBy      string `db:"order_by"`
	IsDescending bool   `db:"is_descending"`
}

func (r *StudentRepo) GetWithParams(params StudentParams) ([]model.Student, error) {
	query := `
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
	ORDER BY :order_by`
	if params.IsDescending {
		query += " DESC"
	} else {
		query += " ASC"
	}

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
	err := r.db.Get(&student, "SELECT * FROM Student WHERE id = $1", id)
	return student, err
}

func (r *StudentRepo) Create(student model.Student) (int64, error) {
	result, err := r.db.NamedExec(`
		INSERT INTO Student 
			(firstname, middlename, lastnamem, gender, birthday, form_of_education, personal_file_number, note)
		VALUES
			(:firstname, :middlename, :lastname, :gender, :birthday, :form_of_education, :personal_file_number, :note)`,
		student,
	)
	if err != nil {
		return -1, err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return -1, err
	}

	if rows == 0 {
		return -1, fmt.Errorf("Student not created")
	}

	id, err := result.LastInsertId()
	if err != nil {
		return -1, err
	}

	return id, nil
}

func (r *StudentRepo) Update(student model.Student) error {
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
	result, err := r.db.Exec("DELETE FROM Student WHERE student_ID = $1", id)
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
