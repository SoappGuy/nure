package model

type AttendanceStatus string

const (
	Present AttendanceStatus = "Присутній"
	Absent  AttendanceStatus = "Відсутній"
)

type Attendance struct {
	AttendanceID int64            `db:"attendance_ID" param:"attendance_ID" form:"attendance_ID" query:"attendance_ID"`
	Attendance   AttendanceStatus `db:"attendance" param:"attendance" form:"attendance" query:"attendance"`
	Reason       *string          `db:"reason" param:"reason" form:"reason" query:"reason"`
	StudentID    int64            `db:"student_ID" param:"student_ID" form:"student_ID" query:"student_ID"`
	LessonID     int64            `db:"lesson_ID" param:"lesson_ID" form:"lesson_ID" query:"lesson_ID"`
}
