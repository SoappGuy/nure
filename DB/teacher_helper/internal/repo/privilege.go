package repo

import (
	"fmt"
	"teacher_helper/internal/model"
)

func (r *StudentRepo) GetPrivileges(id int) ([]model.Privilege, error) {
	rows, err := r.db.Query(`
	SELECT
		privilege_ID,
		type,
		description,
		expiration_date,
		student_ID
	FROM
		Privilege
	WHERE
		student_ID = ?`,
		id,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var privileges []model.Privilege
	for rows.Next() {
		var i model.Privilege
		if err := rows.Scan(
			&i.PrivilegeID,
			&i.Type,
			&i.Description,
			&i.ExpirationDate,
			&i.StudentID,
		); err != nil {
			return nil, err
		}
		privileges = append(privileges, i)
	}

	if err := rows.Close(); err != nil {
		return nil, err
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return privileges, nil
}

func (r *StudentRepo) GetPrivilegeByID(id int) (*model.Privilege, error) {
	rows, err := r.db.Query(`
		SELECT
			privilege_ID,
			type,
			description,
			expiration_date,
			student_id
		FROM
			Privilege
		WHERE
			privilege_ID = ?
		`,
		id,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var (
		privileges []model.Privilege
	)
	for rows.Next() {
		var i model.Privilege
		if err := rows.Scan(
			&i.PrivilegeID,
			&i.Type,
			&i.Description,
			&i.ExpirationDate,
			&i.StudentID,
		); err != nil {
			return nil, err
		}
		privileges = append(privileges, i)
	}

	if err := rows.Close(); err != nil {
		return nil, err
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	if len(privileges) != 1 {
		return nil, fmt.Errorf("Expected 1 Privilege, but got %d instead", len(privileges))
	}

	return &privileges[0], nil
}

func (r *StudentRepo) UpdatePrivilege(privilege *model.Privilege) error {
	fmt.Printf("\n\n%#v\n\n", privilege)
	result, err := r.db.NamedExec(
		`UPDATE Privilege
		SET 
			type = :type,
			description = :description,
			expiration_date = :expiration_date,
			student_ID = :student_ID
		WHERE 
			privilege_ID = :privilege_ID
		`,
		privilege,
	)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return fmt.Errorf("Privilege not updated")
	}

	return nil
}

func (r *StudentRepo) DeletePrivilege(id int) error {
	result, err := r.db.Exec("DELETE FROM Privilege WHERE privilege_ID = ?", id)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return fmt.Errorf("Privilege with id %d not found", id)
	}

	return nil
}

func (r *StudentRepo) CreatePrivilege(privilege *model.Privilege) (int64, error) {
	result, err := r.db.NamedExec(
		`INSERT INTO Privilege
			(student_ID, type, description, expiration_date)
		VALUES 
			(:student_ID, :type, :description, :expiration_date)`,
		privilege,
	)
	if err != nil {
		return -1, err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return -1, err
	}

	if rows == 0 {
		return -1, fmt.Errorf("Privilege not created")
	}

	id, err := result.LastInsertId()
	if err != nil {
		return -1, err
	}

	return id, nil
}
