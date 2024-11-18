package repo

import (
	"fmt"
	"teacher_helper/internal/model"

	"github.com/jmoiron/sqlx"
)

type FamilyConnectionRepo struct {
	db *sqlx.DB
}

func NewFamilyConnectionRepo(db *sqlx.DB) *FamilyConnectionRepo {
	return &FamilyConnectionRepo{db: db}
}

type FamilyConnectionParams struct {
	Query        string `db:"query"`
	OrderBy      string `db:"order_by"`
	IsDescending bool   `db:"is_descending"`
}

func (r *FamilyConnectionRepo) GetWithParams(params FamilyConnectionParams) ([]model.FamilyConnection, error) {
	query := `
	SELECT 
		* 
	FROM 
		FamilyConnection 
	WHERE
		kinship LIKE :query OR
		student_id LIKE :query OR
		caretaker_id LIKE :query
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

	var familyConnections []model.FamilyConnection
	for rows.Next() {
		var familyConnection model.FamilyConnection
		err := rows.StructScan(&familyConnection)
		if err != nil {
			return nil, err
		}
		familyConnections = append(familyConnections, familyConnection)
	}
	return familyConnections, nil
}

func (r *FamilyConnectionRepo) GetAll() ([]model.FamilyConnection, error) {
	var familyConnections []model.FamilyConnection
	err := r.db.Select(&familyConnections, "SELECT * FROM FamilyConnection")
	return familyConnections, err
}

func (r *FamilyConnectionRepo) GetByID(id int) (model.FamilyConnection, error) {
	var familyConnection model.FamilyConnection
	err := r.db.Get(&familyConnection, "SELECT * FROM FamilyConnection WHERE family_connection_id = $1", id)
	return familyConnection, err
}

func (r *FamilyConnectionRepo) Create(familyConnection *model.FamilyConnection) (int64, error) {
	result, err := r.db.NamedExec(
		`INSERT INTO FamilyConnection 
			(kinship, student_id, caretaker_id) 
		VALUES 
			(:kinship, :student_id, :caretaker_id)
		`,
		familyConnection,
	)
	if err == nil {
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

func (r *FamilyConnectionRepo) Update(familyConnection *model.FamilyConnection) error {
	result, err := r.db.NamedExec(
		`UPDATE FamilyConnection 
		SET 
			kinship = :kinship, 
			student_id = :student_id, 
			caretaker_id = :caretaker_id 
		WHERE 
			family_connection_id = :family_connection_id
		`,
		familyConnection,
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

func (r *FamilyConnectionRepo) Delete(id int) error {
	result, err := r.db.Exec("DELETE FROM FamilyConnection WHERE family_connection_id = $1", id)

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return fmt.Errorf("FamilyConnection with ID %d not found", id)
	}

	return nil
}
