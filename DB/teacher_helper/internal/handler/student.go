package handler

import (
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
	e.GET("/medicalCard/:id/info", h.GetMedicalCard)
	e.GET("/medicalCard/:id/edit", h.EditMedicalCard)
	e.PUT("/medicalCard/:id", h.UpdateMedicalCard)
	e.POST("/medicalCard", h.CreateMedicalCard)

	e.GET("/students/familyconnections/:id/info", h.GetConnection)
	e.GET("/students/familyconnections/:id/edit", h.EditConnection)
	e.PUT("/students/familyconnections/:family_connection_ID", h.UpdateConnection)
	e.DELETE("/students/familyconnections/:id", h.DeleteConnection)
	e.POST("/students/familyconnections", h.CreateConnection)

	e.GET("/students/select/:id", h.StudentsSelect)

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
	Links       []Link
	Student     model.Student
	MedicalCard model.MedicalCard
	Connections []model.FamilyConnection
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

	medicalCard, err := h.studentRepo.GetMedicalCard(id)
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Can't get medical card"})
	}

	links := NewLinks(PageTypeStudents)

	connections, err := h.studentRepo.GetConnections(id)
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Can't get connections"})
	}

	student_page := StudentPage{
		Links:       links,
		Student:     student,
		MedicalCard: *medicalCard,
		Connections: connections,
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

	return c.Render(http.StatusCreated, "students.html/student", student)
}

func (h *StudentHandler) DeleteStudent(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid student ID"})
	}

	if err = h.studentRepo.Delete(id); err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Can't delete student"})
	}

	return c.NoContent(http.StatusAccepted)
}

func (h *StudentHandler) UpdateStudent(c echo.Context) error {
	student := new(model.Student)
	if err := c.Bind(student); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid caretaker data"})
	}

	if err := h.studentRepo.Update(student); err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Can't update student"})
	}

	return c.Render(http.StatusOK, "student.html/student-info", student)
}

func (h *StudentHandler) StudentsSelect(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid student ID"})
	}

	student := model.Student{}
	if id != -1 {
		student, err = h.studentRepo.GetByID(id)
		if err != nil {
			log.Error(err)
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Can't get student"})
		}
	}

	students, err := h.studentRepo.GetWithParams(repo.QueryParams{
		Query:        "%",
		OrderBy:      "lastname",
		IsDescending: false,
	})
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Can't get students"})
	}

	studentsSelect := map[any]any{
		"Selected": student,
		"Students": students,
	}

	return c.Render(http.StatusOK, "students.html/students-select", studentsSelect)
}

func (h *StudentHandler) GetConnection(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid student ID"})
	}

	parent, err := h.studentRepo.GetConnectionByID(id)
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Can't get parent"})
	}

	return c.Render(http.StatusOK, "student.html/connection", parent)
}

func (h *StudentHandler) EditConnection(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid student ID"})
	}

	parent, err := h.studentRepo.GetConnectionByID(id)
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Can't get parent"})
	}

	return c.Render(http.StatusOK, "student.html/connection-edit", parent)
}

func (h *StudentHandler) UpdateConnection(c echo.Context) error {
	conenction := new(model.FamilyConnectionWithIDs)
	if err := c.Bind(conenction); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid parent data"})
	}

	if err := h.studentRepo.UpdateConnection(conenction); err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Can't update parent"})
	}

	parent, err := h.studentRepo.GetConnectionByID(int(conenction.FamilyConnectionID))
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Can't get parent"})
	}

	return c.Render(http.StatusOK, "student.html/connection", parent)
}

func (h *StudentHandler) DeleteConnection(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid parent ID"})
	}

	if err := h.studentRepo.DeleteConnection(id); err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Can't delete parent"})
	}

	return c.NoContent(http.StatusAccepted)
}

func (h *StudentHandler) CreateConnection(c echo.Context) error {
	parent := new(model.FamilyConnectionWithIDs)
	if err := c.Bind(parent); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid parent data"})
	}

	id, err := h.studentRepo.CreateConnection(parent)
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Can't create parent"})
	}

	connection, err := h.studentRepo.GetConnectionByID(int(id))
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Can't get parent"})
	}

	return c.Render(http.StatusCreated, "student.html/connection", connection)
}

func (h *StudentHandler) GetMedicalCard(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid student ID"})
	}

	medicalCard, err := h.studentRepo.GetMedicalCard(id)
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Can't get MedicalCard"})
	}

	return c.Render(http.StatusOK, "student.html/medical-card", medicalCard)
}

func (h *StudentHandler) EditMedicalCard(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid student ID"})
	}

	medicalCard, err := h.studentRepo.GetMedicalCard(id)
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Can't get MedicalCard"})
	}

	return c.Render(http.StatusOK, "student.html/medical-card-edit", medicalCard)
}

func (h *StudentHandler) UpdateMedicalCard(c echo.Context) error {
	medicalCard := new(model.MedicalCard)
	if err := c.Bind(medicalCard); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid MedicalCard data"})
	}

	if err := h.studentRepo.UpdateMedicalCard(medicalCard); err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Can't update MedicalCard"})
	}

	return c.Render(http.StatusOK, "student.html/medical-card", medicalCard)
}

func (h *StudentHandler) CreateMedicalCard(c echo.Context) error {
	medicalCard := new(model.MedicalCard)
	if err := c.Bind(medicalCard); err != nil {
		return err
	}

	err := h.studentRepo.CreateMedicalCard(medicalCard)
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Can't add MedicalCard"})
	}

	return c.Render(http.StatusCreated, "students.html/medical-card", medicalCard)
}
