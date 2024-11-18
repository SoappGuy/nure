package repo

import (
	"github.com/jmoiron/sqlx"
	"teacher_helper/internal/model"
)

type FamilyConnectionRepo struct {
	db *sqlx.DB
}

func NewFamilyConnectionRepo(db *sqlx.DB) *FamilyConnectionRepo {
	return &FamilyConnectionRepo{db: db}
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

func (r *FamilyConnectionRepo) Create(familyConnection *model.FamilyConnection) (id int64, err error) {
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

	id, err = result.LastInsertId()
	return id, err
}

func (r *FamilyConnectionRepo) Update(familyConnection *model.FamilyConnection) error {
	_, err := r.db.NamedExec(
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
	return err
}

func (r *FamilyConnectionRepo) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM FamilyConnection WHERE family_connection_id = $1", id)
	return err
}
