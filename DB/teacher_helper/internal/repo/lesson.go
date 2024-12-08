package repo

import (
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
			AND start_date = ?`,
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
