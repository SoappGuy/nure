package handler

import (
	"net/http"

	"teacher_helper/internal/model"
	"teacher_helper/internal/repo"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type StudentHandler struct {
	studentRepo *repo.StudentRepo
}

func NewStudentHandler(studentRepo *repo.StudentRepo) *StudentHandler {
	return &StudentHandler{studentRepo: studentRepo}
}

func (h *StudentHandler) RegisterRoutes(e *echo.Echo) {
	e.GET("/students", h.GetStudents)
}

type StudentsPage struct {
	Title    string
	Links    []Link
	Students []model.Student
}

func (h *StudentHandler) GetStudents(c echo.Context) error {
	query_params := repo.StudentParams{
		Query:        "%",
		OrderBy:      "lastname",
		IsDescending: false,
	}

	students, err := h.studentRepo.GetWithParams(query_params)
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Can't get students"})
	}

	links := NewLinks(PageTypeStudents)

	students_page := StudentsPage{
		Title:    "Учні",
		Links:    links,
		Students: students,
	}

	return c.Render(http.StatusOK, "students.html", students_page)
}
