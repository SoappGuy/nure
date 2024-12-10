package repo

import (
	"database/sql"
	"errors"
	"fmt"
	"teacher_helper/internal/model"
	"time"

	"github.com/jmoiron/sqlx"
)

type LessonRepo struct {
	db *sqlx.DB
}

func NewLessonRepo(db *sqlx.DB) *LessonRepo {
	return &LessonRepo{db: db}
}

func (r *LessonRepo) GetLessonsAtDate(year int, month time.Month, day int) ([]model.Lesson, error) {
	var lessons []model.Lesson
	err := r.db.Select(
		&lessons,
		`SELECT 
			* 
		FROM 
			Lesson, Subject 
		WHERE 
			Lesson.subject_ID = Subject.subject_ID 
			AND start_date = ?
		ORDER BY lesson_number ASC, start_time ASC`,
		fmt.Sprintf("%d-%d-%d", year, month, day),
	)
	return lessons, err
}

func (r *LessonRepo) GetLessonByID(lessonID int64) (model.Lesson, error) {
	var lesson model.Lesson
	err := r.db.Get(
		&lesson,
		`SELECT 
			* 
		FROM 
			Lesson, Subject 
		WHERE 
			Lesson.subject_ID = Subject.subject_ID 
			AND lesson_ID = ?`,
		lessonID,
	)
	return lesson, err
}

func (r *LessonRepo) CreateLesson(lesson *model.LessonWithIDs) (int64, error) {
	result, err := r.db.NamedExec(
		`INSERT INTO Lesson
			(topic, start_date, start_time, lesson_number, subject_ID)
		VALUES
			(:topic, :start_date, :start_time, :lesson_number, :subject_ID)`,
		lesson,
	)
	if err != nil {
		return -1, err
	}

	row, err := result.RowsAffected()
	if err != nil {
		return -1, err
	}

	if row == 0 {
		return -1, fmt.Errorf("Lesson not created")
	}

	id, err := result.LastInsertId()
	if err != nil {
		return -1, err
	}

	return id, nil
}

func (r *LessonRepo) DeleteLesson(id int) error {
	result, err := r.db.Exec("DELETE FROM Lesson WHERE lesson_ID = ?", id)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return fmt.Errorf("Lesson with id %d not found", id)
	}

	return nil
}

func (r *LessonRepo) UpdateLesson(LessonWithIDs *model.LessonWithIDs) error {
	result, err := r.db.NamedExec(
		`UPDATE Lesson
		SET 
			topic = :topic,
			start_date = :start_date,
			start_time = :start_time,
			lesson_number = :lesson_number,
			subject_ID = :subject_ID
		WHERE 
			lesson_ID = :lesson_ID
		`,
		LessonWithIDs,
	)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return fmt.Errorf("Lesson not updated")
	}

	return nil
}

func (r *LessonRepo) GetAllStudents() ([]model.Student, error) {
	var students []model.Student
	err := r.db.Select(&students, "SELECT * FROM Student ORDER BY lastname ASC, firstname ASC, middlename ASC")
	return students, err
}

func (r *LessonRepo) GetMarksForStudentAtLesson(studentID, lessonID int64) ([]model.Mark, error) {
	var marks []model.Mark
	err := r.db.Select(
		&marks,
		`SELECT
			*
		FROM
			Mark
		WHERE
			student_ID = ?
			AND lesson_ID = ?`,
		studentID,
		lessonID,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return []model.Mark{}, nil
		}
		return nil, err
	}
	return marks, err
}

func (r *LessonRepo) GetAttendanceForStudentAtLesson(studentID, lessonID int64) (*model.Attendance, error) {
	var attendance model.Attendance
	err := r.db.Get(
		&attendance,
		`SELECT
			*
		FROM
			Attendance
		WHERE
			student_ID = ?
			AND lesson_ID = ?`,
		studentID,
		lessonID,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			attendance, err := r.CreateAttendanceOnDemand(studentID, lessonID)
			if err != nil {
				return nil, err
			}
			return attendance, nil
		}
		return nil, err
	}
	return &attendance, nil
}

func (r *LessonRepo) CreateAttendanceOnDemand(studentID, lessonID int64) (*model.Attendance, error) {
	result, err := r.db.Exec(
		`INSERT INTO Attendance
			(attendance, reason, student_ID, lesson_ID)
		VALUES
			(?, 'без уточнень', ?, ?)`,
		model.Present,
		studentID,
		lessonID,
	)
	if err != nil {
		return nil, err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}

	if rows == 0 {
		return nil, fmt.Errorf("Attendance not created")
	}

	attendance := model.Attendance{
		Attendance: model.Present,
		StudentID:  studentID,
		LessonID:   lessonID,
	}

	return &attendance, nil
}

func (r *LessonRepo) CreateMark(mark *model.Mark) error {
	result, err := r.db.NamedExec(
		`INSERT INTO Mark
			(mark, student_ID, lesson_ID)
		VALUES
			(:mark, :student_ID, :lesson_ID)`,
		mark,
	)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return fmt.Errorf("Mark not created")
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	mark.MarkID = id
	return nil
}

func (r *LessonRepo) DeleteMark(id int) error {
	result, err := r.db.Exec("DELETE FROM Mark WHERE mark_ID = ?", id)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return fmt.Errorf("Mark with id %d not found", id)
	}

	return nil
}

func (r *LessonRepo) FlipAttendanceWithID(id int) (*model.Attendance, error) {
	_, err := r.db.Exec(
		`
		UPDATE Attendance
		SET
			attendance = CASE
				WHEN attendance = 'Відсутній' THEN 'Присутній'
				WHEN attendance = 'Присутній' THEN 'Відсутній'
			END
		WHERE
			attendance_ID = ?;
		`,
		id,
	)
	if err != nil {
		return nil, err
	}

	var attendance model.Attendance
	err = r.db.Get(
		&attendance,
		`
		SELECT
			*
		FROM
			Attendance
		WHERE
			attendance_ID = ?;
		`,
		id,
	)
	if err != nil {
		return nil, err
	}

	return &attendance, nil
}

func (r *LessonRepo) GetAttendanceByID(id int) (*model.Attendance, error) {
	attendance := new(model.Attendance)
	err := r.db.Get(attendance, `SELECT * FROM Attendance WHERE attendance_ID = ?`, id)
	return attendance, err
}

func (r *LessonRepo) UpdateAttendance(attendance *model.Attendance) error {
	result, err := r.db.NamedExec(
		`
		UPDATE Attendance
		SET
			reason = :reason
		WHERE
			attendance_ID = :attendance_ID
		`,
		attendance,
	)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return fmt.Errorf("Attendance not updated")
	}

	return nil
}
