package repo

import (
	"fmt"
	"teacher_helper/internal/model"

	"github.com/jmoiron/sqlx"
)

type SubjectRepo struct {
	db *sqlx.DB
}

func NewSubjectRepo(db *sqlx.DB) *SubjectRepo {
	return &SubjectRepo{db: db}
}

func (r *SubjectRepo) GetWithParams(params QueryParams) ([]model.Subject, error) {
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
		Subject 
	WHERE
		title LIKE :query OR
		description LIKE :query
	ORDER BY %s %s`, params.OrderBy, order)

	rows, err := r.db.NamedQuery(query, params)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var subjects []model.Subject
	for rows.Next() {
		var subject model.Subject
		err := rows.StructScan(&subject)
		if err != nil {
			return nil, err
		}
		subjects = append(subjects, subject)
	}

	return subjects, nil
}

func (r *SubjectRepo) GetAll() ([]model.Subject, error) {
	var subjects []model.Subject
	err := r.db.Select(&subjects, "SELECT * FROM Subject")
	if err != nil {
		return nil, err
	}

	return subjects, nil
}

func (r *SubjectRepo) GetByID(id int) (*model.Subject, error) {
	var subject model.Subject
	err := r.db.Get(&subject, "SELECT * FROM Subject WHERE subject_ID = ?", id)
	if err != nil {
		return nil, err
	}

	return &subject, nil
}

func (r *SubjectRepo) Create(subject *model.Subject) error {
	result, err := r.db.NamedExec(
		`INSERT INTO Subject
			(title, description, number_of_hours)
		VALUES
			(:title, :description, :number_of_hours)`,
		subject,
	)
	if err != nil {
		return err
	}

	row, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if row == 0 {
		return fmt.Errorf("Subject not created")
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	subject.SubjectID = id

	return nil
}

func (r *SubjectRepo) Update(subject *model.Subject) error {
	result, err := r.db.NamedExec(
		`UPDATE Subject
		SET
			title = :title,
			description = :description,
			number_of_hours = :number_of_hours
		WHERE
			subject_ID = :subject_ID`,
		subject,
	)
	if err != nil {
		return err
	}

	row, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if row == 0 {
		return fmt.Errorf("Subject not updated")
	}

	return nil
}

func (r *SubjectRepo) Delete(id int) error {
	result, err := r.db.Exec("DELETE FROM Subject WHERE subject_ID = ?", id)
	if err != nil {
		return err
	}

	row, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if row == 0 {
		return fmt.Errorf("Subject with id %d not deleted", id)
	}

	return nil
}
