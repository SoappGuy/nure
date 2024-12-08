package model

type Gender string

const (
	Male   Gender = "M"
	Female Gender = "F"
)

type FormOfEducation string

const (
	DayForm     FormOfEducation = "Денна"
	EveningForm FormOfEducation = "Вечірня"
	HomeForm    FormOfEducation = "Домашня"
)

type Student struct {
	StudentID          int64           `db:"student_ID" param:"student_ID" query:"student_ID" form:"student_ID"`
	Firstname          string          `db:"firstname" param:"firstname" query:"firstname" form:"firstname"`
	Middlename         string          `db:"middlename" param:"middlename" query:"middlename" form:"middlename"`
	Lastname           string          `db:"lastname" param:"lastname" query:"lastname" form:"lastname"`
	Gender             Gender          `db:"gender" param:"gender" query:"gender" form:"gender"`
	Birthday           Date            `db:"birthday" param:"birthday" query:"birthday" form:"birthday"`
	FormOfEducation    FormOfEducation `db:"form_of_education" param:"form_of_education" query:"form_of_education" form:"form_of_education"`
	PersonalFileNumber string          `db:"personal_file_number" param:"personal_file_number" query:"personal_file_number" form:"personal_file_number"`
	Note               *string         `db:"note" param:"note" query:"note" form:"note"`
}
