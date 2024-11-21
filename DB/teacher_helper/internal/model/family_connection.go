package model

type Kinship string

const (
	Father      Kinship = "Батько"
	Mother      Kinship = "Мати"
	StepFather  Kinship = "Вітчим"
	StepMother  Kinship = "Мачуха"
	Uncle       Kinship = "Дядько"
	Aunt        Kinship = "Тітка"
	Brother     Kinship = "Брат"
	Sister      Kinship = "Сестра"
	Grandfather Kinship = "Дідусь"
	Grandmother Kinship = "Бабуся"
	Other       Kinship = "Інше"
)

type FamilyConnectionWithIDs struct {
	FamilyConnectionID int64   `db:"family_connection_ID" param:"family_connection_ID" query:"family_connection_ID" form:"family_connection_ID"`
	Student            int64   `db:"student_ID" param:"student_ID" query:"student_ID" form:"student_ID"`
	Caretaker          int64   `db:"caretaker_ID" param:"caretaker_ID" query:"caretaker_ID" form:"caretaker_ID"`
	Kinship            Kinship `db:"kinship" param:"kinship" query:"kinship" form:"kinship"`
}

type FamilyConnection struct {
	FamilyConnectionID int64   `db:"family_connection_ID" `
	Kinship            Kinship `db:"kinship"`

	// tags `db:""` are set for sqlx to mark this sctructs as embedded for parsing
	Student   Student   `db:""`
	Caretaker Caretaker `db:""`
}

// type FamilyConnectionFlat struct {
// 	FamilyConnectionID int64  `json:"family_connection_id"`
// 	Kinship            string `json:"kinship"`
// 	StudentID          int64  `json:"student_id"`
//
