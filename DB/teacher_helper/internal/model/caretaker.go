package model

type Caretaker struct {
	CaretakerID int64  `db:"caretaker_ID" param:"caretaker_ID" query:"caretaker_ID" form:"caretaker_ID"`
	Firstname   string `db:"firstname" param:"firstname" query:"firstname" form:"firstname"`
	Middlename  string `db:"middlename" param:"middlename" query:"middlename" form:"middlename"`
	Lastname    string `db:"lastname" param:"lastname" query:"lastname" form:"lastname"`
	Phone       string `db:"phone" param:"phone" query:"phone" form:"phone"`
	Email       string `db:"email" param:"email" query:"email" form:"email"`
}
