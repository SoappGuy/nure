package handler

import (
	"fmt"
	"net/http"
	"strconv"

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
	e.GET("/students/search", h.SearchStudents)
	e.POST("/students", h.CreateStudent)
	e.DELETE("/students/:id", h.DeleteStudent)
	e.PUT("/students/:id", h.UpdateStudent)

	e.GET("/students/:id", h.GetStudent)
	e.GET("/students/:id/info", h.StudentInfo)
	e.GET("/students/:id/edit", h.StudentInfoEdit)
}

type StudentsPage struct {
	Title    string
	Links    []Link
	Students []model.Student
}

type StudentPage struct {
	Links   []Link
	Student model.Student
}

func (h *StudentHandler) GetStudents(c echo.Context) error {
	query_params := repo.QueryParams{
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

	return c.Render(http.StatusOK, "students.html/base", students_page)
}

func (h *StudentHandler) GetStudent(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid student ID"})
	}

	student, err := h.studentRepo.GetByID(id)
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Can't get student"})
	}

	links := NewLinks(PageTypeStudents)

	student_page := StudentPage{
		Links:   links,
		Student: student,
	}

	return c.Render(http.StatusOK, "student.html/base", student_page)
}

func (h *StudentHandler) StudentInfoEdit(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid student ID"})
	}

	student, err := h.studentRepo.GetByID(id)
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Can't get student"})
	}

	return c.Render(http.StatusOK, "student.html/student-info-edit", student)
}

func (h *StudentHandler) StudentInfo(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid student ID"})
	}

	student, err := h.studentRepo.GetByID(id)
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Can't get student"})
	}

	return c.Render(http.StatusOK, "student.html/student-info", student)
}

func (h *StudentHandler) SearchStudents(c echo.Context) error {
	query := new(repo.QueryParams)
	if err := c.Bind(query); err != nil {
		return err
	}

	if err := query.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	students, err := h.studentRepo.GetWithParams(*query)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Can't get students"})
	}

	return c.Render(http.StatusOK, "students.html/students", students)
}

func (h *StudentHandler) CreateStudent(c echo.Context) error {
	student := new(model.Student)
	if err := c.Bind(student); err != nil {
		return err
	}

	err := h.studentRepo.Create(student)
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Can't add student"})
	}

	return c.Render(http.StatusOK, "students.html/student", student)
}

func (h *StudentHandler) DeleteStudent(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid student ID"})
	}

	err = h.studentRepo.Delete(id)
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Can't delete student"})
	}

	return c.NoContent(http.StatusOK)
}

func (h *StudentHandler) UpdateStudent(c echo.Context) error {
	student := new(model.Student)
	if err := c.Bind(student); err != nil {
		return err
	}

	fmt.Printf("\n%#v\n", student)

	err := h.studentRepo.Update(student)
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Can't update student"})
	}

	return c.Render(http.StatusOK, "student.html/student-info", student)
}
