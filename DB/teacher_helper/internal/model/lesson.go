package model

type Lesson struct {
	LessonID     int64    `db:"lesson_ID" param:"lesson_ID" query:"lesson_ID" form:"lesson_ID"`
	Topic        string   `db:"topic" param:"topic" query:"topic" form:"topic"`
	StartDate    Date     `db:"start_date" param:"start_date" query:"start_date" form:"start_date"`
	StartTime    NullTime `db:"start_time" param:"start_time" query:"start_time" form:"start_time"`
	LessonNumber int64    `db:"lesson_number" param:"lesson_number" query:"lesson_number" form:"lesson_number"`
	Subject      Subject  `db:""`
}

type LessonWithIDs struct {
	LessonID     int64    `db:"lesson_ID" param:"lesson_ID" query:"lesson_ID" form:"lesson_ID"`
	Topic        string   `db:"topic" param:"topic" query:"topic" form:"topic"`
	StartDate    Date     `db:"start_date" param:"start_date" query:"start_date" form:"start_date"`
	StartTime    NullTime `db:"start_time" param:"start_time" query:"start_time" form:"start_time"`
	LessonNumber int64    `db:"lesson_number" param:"lesson_number" query:"lesson_number" form:"lesson_number"`
	SubjectID    int64    `db:"subject_ID" param:"subject_ID" query:"subject_ID" form:"subject_ID"`
}
