package model

type Gender string

const (
	Male   Gender = "M"
	Female Gender = "F"
)

type FormOfEducation string

const (
	Day     FormOfEducation = "Денна"
	Evening FormOfEducation = "Вечірня"
	Home    FormOfEducation = "Домашня"
)

type Student struct {
	StudentID          int64           `db:"student_ID"`
	Firstname          string          `db:"firstname"`
	Middlename         string          `db:"middlename"`
	Lastname           string          `db:"lastname"`
	Gender             Gender          `db:"gender"`
	Birthday           string          `db:"birthday"`
	FormOfEducation    FormOfEducation `db:"form_of_education"`
	PersonalFileNumber string          `db:"personal_file_number"`
	Note               *string         `db:"note"`
}
