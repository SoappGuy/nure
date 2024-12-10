package service

import (
	"fmt"
	"teacher_helper/internal/model"
	"teacher_helper/internal/repo"
)

type StudentAtLesson struct {
	model.Student
	LessonID   int64
	Marks      []model.Mark
	Attendance model.Attendance
}

func GetLessonDetails(lessonID int64, repo *repo.LessonRepo) ([]StudentAtLesson, error) {
	lessonDetails := []StudentAtLesson{}

	students, err := repo.GetAllStudents()
	if err != nil {
		return nil, err
	}

	for _, student := range students {
		fmt.Printf("Student: %v\n", student)
		marks, err := repo.GetMarksForStudentAtLesson(student.StudentID, lessonID)
		if err != nil {
			return nil, err
		}

		attendance, err := repo.GetAttendanceForStudentAtLesson(student.StudentID, lessonID)
		if err != nil {
			return nil, err
		}

		studentDetails := StudentAtLesson{
			Student:    student,
			Marks:      marks,
			Attendance: *attendance,
			LessonID:   lessonID,
		}

		lessonDetails = append(lessonDetails, studentDetails)
	}

	return lessonDetails, nil
}
