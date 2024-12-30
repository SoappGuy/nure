package repo

import (
	"fmt"
	"teacher_helper/internal/model"
)

func (r *StudentRepo) GetMedicalCard(id int) (*model.MedicalCard, error) {
	medicalCard := new(model.MedicalCard)
	err := r.db.Get(medicalCard, "SELECT * FROM MedicalCard WHERE student_id = ?", id)
	if err != nil {
		return nil, err
	}

	return medicalCard, nil
}

func (r *StudentRepo) CreateMedicalCard(medicalCard *model.MedicalCard) error {
	result, err := r.db.NamedExec(`
		INSERT INTO MedicalCard
			(student_ID, weight, height, health_group, blood_group, rh_factor, last_inspection_date, next_inspection_date)
		VALUES
			(:student_ID, :weight, :height, :health_group, :blood_group, :rh_factor, :last_inspection_date, :next_inspection_date)
		`,
		medicalCard,
	)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return fmt.Errorf("MedicalCard not created")
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	medicalCard.MedicalCardID = id

	return nil
}

func (r *StudentRepo) UpdateMedicalCard(medicalCard *model.MedicalCard) error {
	result, err := r.db.NamedExec(
		`UPDATE MedicalCard
		SET 
			weight = :weight,
			height = :height,
			health_group = :health_group,
			blood_group = :blood_group,
			rh_factor = :rh_factor,
			last_inspection_date = :last_inspection_date,
			next_inspection_date = :next_inspection_date
		WHERE 
			medical_card_ID = :medical_card_ID
		`,
		medicalCard,
	)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return fmt.Errorf("MedicalCard not updated")
	}

	return nil
}

func (r *StudentRepo) DeleteMedicalCard(id int) error {
	result, err := r.db.Exec("DELETE FROM MedicalCard WHERE medical_card_ID = ?", id)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return fmt.Errorf("MedicalCard with id %d not deleted", id)
	}

	return nil
}
