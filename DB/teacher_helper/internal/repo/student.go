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
