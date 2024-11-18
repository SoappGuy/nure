package repo

import (
	"fmt"
	"teacher_helper/internal/model"

	"github.com/jmoiron/sqlx"
)

type CaretakerRepo struct {
	db *sqlx.DB
}

func NewCaretakerRepo(db *sqlx.DB) *CaretakerRepo {
	return &CaretakerRepo{db: db}
}

type CaretakerParams struct {
	Query        string `db:"query"`
	OrderBy      string `db:"order_by"`
	IsDescending bool   `db:"is_descending"`
}

func (r *CaretakerRepo) GetWithParams(params CaretakerParams) ([]model.Caretaker, error) {
	query := `
	SELECT 
		* 
	FROM 
		Caretaker 
	WHERE
		firstname LIKE :query OR
		middlename LIKE :query OR
		lastname LIKE :query OR
		phone LIKE :query OR
		email LIKE :query
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

	var caretakers []model.Caretaker
	for rows.Next() {
		var caretaker model.Caretaker
		err := rows.StructScan(&caretaker)
		if err != nil {
			return nil, err
		}
		caretakers = append(caretakers, caretaker)
	}

	return caretakers, nil
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

	row, err := result.RowsAffected()
	if err != nil {
		return -1, err
	}

	if row == 0 {
		return -1, fmt.Errorf("CareTaker not created")
	}

	id, err = result.LastInsertId()
	if err != nil {
		return -1, err
	}

	return id, nil
}

func (r *CaretakerRepo) Update(caretaker *model.Caretaker) error {
	result, err := r.db.NamedExec(
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
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return fmt.Errorf("CareTaker not updated")
	}

	return nil
}

func (r *CaretakerRepo) Delete(id int) error {
	result, err := r.db.Exec("DELETE FROM Caretaker WHERE caretaker_ID = $1", id)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return fmt.Errorf("CareTaker with id %d not found", id)
	}

	return nil
}
