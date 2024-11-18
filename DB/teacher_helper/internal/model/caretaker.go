package model

type Caretaker struct {
	CaretakerID int    `db:"caretaker_ID"`
	Firstname   string `db:"firstname"`
	Middlename  string `db:"middlename"`
	Lastname    string `db:"lastname"`
	Phone       string `db:"phone"`
	Email       string `db:"email"`
}
