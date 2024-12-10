package model

type Mark struct {
	MarkID    int64 `db:"mark_ID" param:"mark_ID" form:"mark_ID" query:"mark_ID"`
	Mark      int64 `db:"mark" param:"mark" form:"mark" query:"mark"`
	StudentID int64 `db:"student_ID" param:"student_ID" form:"student_ID" query:"student_ID"`
	LessonID  int64 `db:"lesson_ID" param:"lesson_ID" form:"lesson_ID" query:"lesson_ID"`
}
