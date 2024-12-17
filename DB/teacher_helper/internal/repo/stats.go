package repo

import (
	"teacher_helper/internal/model"

	"github.com/jmoiron/sqlx"
)

type StatsRepo struct {
	db *sqlx.DB
}

func NewStatsRepo(db *sqlx.DB) *StatsRepo {
	return &StatsRepo{db: db}
}

func (r *StatsRepo) GetNextMedCheck() ([]model.NextMedCheck, error) {
	var nextMedCheck []model.NextMedCheck
	err := r.db.Select(&nextMedCheck, `
SELECT
    Student.student_ID,
    CONCAT_WS(
        ' ',
        Student.lastname,
        Student.firstname,
        Student.middlename
    ) AS fullname,
    MedicalCard.next_inspection_date
FROM
    MedicalCard,
    Student
WHERE
    MedicalCard.student_ID = Student.student_ID
ORDER BY
    MedicalCard.next_inspection_date;
`)
	if err != nil {
		return nil, err
	}
	return nextMedCheck, nil
}

func (r *StatsRepo) GetStudentsRating() ([]model.StudentsRating, error) {
	var studentsRating []model.StudentsRating
	err := r.db.Select(&studentsRating, `
		SELECT
			Student.student_ID,
			CONCAT_WS(
				' ',
				Student.lastname,
				Student.firstname,
				Student.middlename
			) AS fullname,
			AVG(Mark.mark) AS avg_mark
		FROM
			Mark
			JOIN Student ON Mark.student_ID = Student.student_ID
		GROUP BY
			Student.student_ID
		ORDER BY
			avg_mark DESC;
	`)
	if err != nil {
		return nil, err
	}
	return studentsRating, nil
}

func (r *StatsRepo) GetPrivilegeSchedule() ([]model.PrivilegeSchedule, error) {
	var privilegeSchedule []model.PrivilegeSchedule
	err := r.db.Select(&privilegeSchedule, `
	SELECT
		Student.student_ID,
		CONCAT_WS(
			' ',
			Student.lastname,
			Student.firstname,
			Student.middlename
		) AS fullname,
		Privilege.type,
		Privilege.expiration_date
	FROM
		Privilege
		JOIN Student ON Privilege.student_ID = Student.student_ID
	ORDER BY
			CASE 
			WHEN Privilege.expiration_date IS NULL THEN 1
			ELSE 0
		END,
		Privilege.expiration_date;
	`)
	if err != nil {
		return nil, err
	}
	return privilegeSchedule, nil
}

func (r *StatsRepo) GetFormOfEducationCount() ([]model.FormOfEducationCount, error) {
	var formOfEducationCount []model.FormOfEducationCount
	err := r.db.Select(&formOfEducationCount, `
	SELECT
		Student.form_of_education,
		COUNT(Student.student_ID) AS count
	FROM
		Student
	GROUP BY
		Student.form_of_education;
	`)
	if err != nil {
		return nil, err
	}
	return formOfEducationCount, nil
}

func (r *StatsRepo) GetGenderCount() ([]model.GenderCount, error) {
	var genderCount []model.GenderCount
	err := r.db.Select(&genderCount, `
		SELECT
			gender,
			COUNT(*) AS count
		FROM
			Student
		GROUP BY
			gender;
	`)
	if err != nil {
		return nil, err
	}
	return genderCount, nil
}

func (r *StatsRepo) GetHealthGroupCount() ([]model.HealthGroupCount, error) {
	var healthGroupCount []model.HealthGroupCount
	err := r.db.Select(&healthGroupCount, `
SELECT
    MedicalCard.health_group,
    COUNT(*) AS count
FROM
    MedicalCard
GROUP BY
    MedicalCard.health_group;
`)
	if err != nil {
		return nil, err
	}
	return healthGroupCount, nil
}

func (r *StatsRepo) GetBloodGroupCount() ([]model.BloodGroupCount, error) {
	var bloodGroupCount []model.BloodGroupCount
	err := r.db.Select(&bloodGroupCount, `
		SELECT
			MedicalCard.blood_group,
			MedicalCard.rh_factor,
			COUNT(*) AS count
		FROM
			MedicalCard
		GROUP BY
			MedicalCard.blood_group,
			MedicalCard.rh_factor
		ORDER BY
			MedicalCard.blood_group,
			MedicalCard.rh_factor;
	`)
	if err != nil {
		return nil, err
	}
	return bloodGroupCount, nil
}
