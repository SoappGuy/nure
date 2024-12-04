package model

type Subject struct {
	SubjectID     int64  `db:"subject_ID" param:"subject_ID" query:"subject_ID" form:"subject_ID"`
	Title         string `db:"title" param:"title" query:"title" form:"title"`
	Description   string `db:"description" param:"description" query:"description" form:"description"`
	NumberOfHours int64  `db:"number_of_hours" param:"number_of_hours" query:"number_of_hours" form:"number_of_hours"`
}
