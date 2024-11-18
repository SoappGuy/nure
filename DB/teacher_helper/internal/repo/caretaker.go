package repo

import (
	"github.com/jmoiron/sqlx"
	"teacher_helper/internal/model"
)

type CaretakerRepo struct {
	db *sqlx.DB
}

func NewCaretakerRepo(db *sqlx.DB) *CaretakerRepo {
	return &CaretakerRepo{db: db}
}

func (r *CaretakerRepo) GetAll() ([]model.Caretaker, error) {
	var caretakers []model.Caretaker
	err := r.db.Select(&caretakers, "SELECT * FROM Caretaker")
	return caretakers, err
}
