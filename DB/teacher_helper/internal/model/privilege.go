package model

type Privilege struct {
	PrivilegeID    int64  `db:"privilege_ID" param:"privilege_ID" query:"privilege_ID" form:"privilege_ID"`
	StudentID      int64  `db:"student_ID" param:"student_ID" query:"student_ID" form:"student_ID"`
	Type           string `db:"type" param:"type" query:"type" form:"type"`
	Description    string `db:"description" param:"description" query:"description" form:"description"`
	ExpirationDate Date   `db:"expiration_date" param:"expiration_date" query:"expiration_date" form:"expiration_date"`
}
