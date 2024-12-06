package model

//	type Caretaker struct {
//		CaretakerID int64  `db:"caretaker_ID" param:"caretaker_ID" query:"caretaker_ID" form:"caretaker_ID"`
//		Firstname   string `db:"firstname" param:"firstname" query:"firstname" form:"firstname"`
//		Middlename  string `db:"middlename" param:"middlename" query:"middlename" form:"middlename"`
//		Lastname    string `db:"lastname" param:"lastname" query:"lastname" form:"lastname"`
//		Phone       string `db:"phone" param:"phone" query:"phone" form:"phone"`
//		Email       string `db:"email" param:"email" query:"email" form:"email"`
//	}
//
// CREATE TABLE MedicalCard (
//
//	medical_card_ID INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
//	student_ID INT NOT NULL,
//	weight FLOAT NOT NULL,
//	height FLOAT NOT NULL,
//	health_group ENUM('1', '2', '3', '4', '5') NOT NULL,
//	blood_group ENUM('I', 'II', 'III', 'IV') NOT NULL,
//	rh_factor ENUM('+', '-') NOT NULL,
//	last_inspection_date DATE NOT NULL,
//	next_inspection_date DATE NOT NULL,
//	note TEXT,
//	FOREIGN KEY (student_ID) REFERENCES Student (student_ID) ON DELETE CASCADE ON UPDATE CASCADE
//
// );
type HealthGroup string

const (
	FirstHealthGroup  HealthGroup = "1"
	SecondHealthGroup HealthGroup = "2"
	ThirdHealthGroup  HealthGroup = "3"
	FourthHealthGroup HealthGroup = "4"
	FifthHealthGroup  HealthGroup = "5"
)

type BloodGroup string

const (
	FirstBloodGroup  BloodGroup = "I"
	SecondBloodGroup BloodGroup = "II"
	ThirdBloodGroup  BloodGroup = "III"
	FourthBloodGroup BloodGroup = "IV"
)

type RhFactor string

const (
	PositiveRhFactor RhFactor = "+"
	NegativeRhFactor RhFactor = "-"
)

type MedicalCard struct {
	MedicalCardID      int64       `db:"medical_card_ID" param:"medical_card_ID" query:"medical_card_ID" form:"medical_card_ID"`
	StudentID          int64       `db:"student_ID" param:"student_ID" query:"student_ID" form:"student_ID"`
	Weight             float64     `db:"weight" param:"weight" query:"weight" form:"weight"`
	Height             float64     `db:"height" param:"height" query:"height" form:"height"`
	HealthGroup        HealthGroup `db:"health_group" param:"health_group" query:"health_group" form:"health_group"`
	BloodGroup         BloodGroup  `db:"blood_group" param:"blood_group" query:"blood_group" form:"blood_group"`
	RhFactor           RhFactor    `db:"rh_factor" param:"rh_factor" query:"rh_factor" form:"rh_factor"`
	LastInspectionDate Date        `db:"last_inspection_date" param:"last_inspection_date" query:"last_inspection_date" form:"last_inspection_date"`
	NextInspectionDate Date        `db:"next_inspection_date" param:"next_inspection_date" query:"next_inspection_date" form:"next_inspection_date"`
	Note               *string     `db:"note" param:"note" query:"note" form:"note"`
}
