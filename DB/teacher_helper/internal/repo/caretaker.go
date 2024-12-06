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

func (r *CaretakerRepo) GetWithParams(params QueryParams) ([]model.Caretaker, error) {
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
		Caretaker 
	WHERE
		firstname LIKE :query OR
		middlename LIKE :query OR
		lastname LIKE :query OR
		phone LIKE :query OR
		email LIKE :query
	ORDER BY %s %s`, params.OrderBy, order)

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
	err := r.db.Get(&caretaker, "SELECT * FROM Caretaker WHERE caretaker_ID = ?", id)
	return caretaker, err
}

func (r *CaretakerRepo) Create(caretaker *model.Caretaker) (err error) {
	result, err := r.db.NamedExec(
		`INSERT INTO Caretaker 
			(firstname, middlename, lastname, phone, email) 
		VALUES 
			(:firstname, :middlename, :lastname, :phone, :email)`,
		caretaker,
	)

	row, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if row == 0 {
		return fmt.Errorf("CareTaker not created")
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	caretaker.CaretakerID = id

	return nil
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
		return fmt.Errorf("Caretaker not updated")
	}

	return nil
}

func (r *CaretakerRepo) Delete(id int) error {
	result, err := r.db.Exec("DELETE FROM Caretaker WHERE caretaker_ID = ?", id)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return fmt.Errorf("CareTaker with id %d not deleted", id)
	}

	return nil
}

func (r *CaretakerRepo) GetConnections(id int) ([]model.FamilyConnection, error) {
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
		FamilyConnection.caretaker_ID = ?`,
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

func (r *CaretakerRepo) GetConnectionByID(id int) (*model.FamilyConnection, error) {
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

func (r *CaretakerRepo) UpdateConnection(kid *model.FamilyConnectionWithIDs) error {
	result, err := r.db.NamedExec(
		`UPDATE FamilyConnection 
		SET 
			kinship = :kinship, 
			student_ID = :student_ID, 
			caretaker_ID = :caretaker_ID 
		WHERE 
			family_connection_ID = :family_connection_ID
		`,
		kid,
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

func (r *CaretakerRepo) DeleteConnection(id int) error {
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

func (r *CaretakerRepo) CreateConnection(kid *model.FamilyConnectionWithIDs) (int64, error) {
	result, err := r.db.NamedExec(
		`INSERT INTO FamilyConnection 
			(student_ID, caretaker_ID, kinship) 
		VALUES 
			(:student_ID, :caretaker_ID, :kinship)`,
		kid,
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
