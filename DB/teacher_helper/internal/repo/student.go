package repo

import (
	"github.com/jmoiron/sqlx"
	"teacher_helper/internal/model"
)

type StudentRepo struct {
	db *sqlx.DB
}

func NewStudentRepo(db *sqlx.DB) *StudentRepo {
	return &StudentRepo{db: db}
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

	id, err := result.LastInsertId()
	return id, err
}

func (r *StudentRepo) Update(student model.Student) error {
	_, err := r.db.NamedExec(`
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
	return err
}

func (r *StudentRepo) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM Student WHERE student_ID = $1", id)
	return err
}
