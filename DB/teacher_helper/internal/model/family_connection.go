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

type FamilyConnection struct {
	FamilyConnectionID int64     `json:"family_connection_id"`
	Kinship            Kinship   `json:"kinship"`
	Student            Student   `json:"student_id"`
	Caretaker          Caretaker `json:"caretaker_id"`
}
