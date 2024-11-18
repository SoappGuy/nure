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

func (r *CaretakerRepo) GetByID(id int) (model.Caretaker, error) {
	var caretaker model.Caretaker
	err := r.db.Get(&caretaker, "SELECT * FROM Caretaker WHERE caretaker_ID = $1", id)
	return caretaker, err
}

func (r *CaretakerRepo) Create(caretaker *model.Caretaker) (id int64, err error) {
	result, err := r.db.NamedExec(
		`INSERT INTO Caretaker 
			(firstname, middlename, lastname, phone, email) 
		VALUES 
			(:firstname, :middlename, :lastname, :phone, :email)", caretaker)
		`,
		caretaker,
	)

	if err == nil {
		return -1, err
	}

	id, err = result.LastInsertId()
	return id, err
}

func (r *CaretakerRepo) Update(caretaker *model.Caretaker) error {
	_, err := r.db.NamedExec(
		`UPDATE Caretaker 
		SET 
			firstname = :firstname, 
			middlename = :middlename, 
			lastname = :lastname, 
			phone = :phone, 
			email = :email 
		WHERE 
			caretaker_ID = :caretaker_ID
		`,
		caretaker,
	)
	return err
}

func (r *CaretakerRepo) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM Caretaker WHERE caretaker_ID = $1", id)
	return err
}
