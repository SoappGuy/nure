package handler

import (
	"net/http"

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
	e.GET("/students", h.GetAllStudents)
}

func (h *StudentHandler) GetAllStudents(c echo.Context) error {
	students, err := h.studentRepo.GetAll()
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Can't get students"})
	}
	return c.Render(http.StatusOK, "students.html", students)
}
