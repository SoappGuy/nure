package model

type NextMedCheck struct {
	StudentID          int    `db:"student_ID" json:"student_ID"`
	Fullname           string `db:"fullname" json:"fullname"`
	NextInspectionDate Date   `db:"next_inspection_date" json:"next_inspection_date"`
}

type StudentsRating struct {
	StudentID int     `db:"student_ID" json:"student_ID"`
	Fullname  string  `db:"fullname" json:"fullname"`
	AvgMark   float32 `db:"avg_mark" json:"avg_mark"`
}

type PrivilegeSchedule struct {
	Fullname                string `db:"fullname" json:"fullname"`
	StudentID               int    `db:"student_ID" json:"student_ID"`
	PrivilegeType           string `db:"type" json:"type"`
	PrivilegeExpirationDate Date   `db:"expiration_date" json:"expiration_date"`
}

type FormOfEducationCount struct {
	FormOfEducation string `db:"form_of_education" json:"form_of_education"`
	Count           int    `db:"count" json:"count"`
}

type GenderCount struct {
	Gender Gender `db:"gender" json:"gender"`
	Count  int    `db:"count" json:"count"`
}

type HealthGroupCount struct {
	HealthGroup HealthGroup `db:"health_group" json:"health_group"`
	Count       int         `db:"count" json:"count"`
}

type BloodGroupCount struct {
	BloodGroup BloodGroup `db:"blood_group" json:"blood_group"`
	RhFactor   RhFactor   `db:"rh_factor" json:"rh_factor"`
	Count      int        `db:"count" json:"count"`
}
