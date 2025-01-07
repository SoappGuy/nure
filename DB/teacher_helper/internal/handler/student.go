package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"teacher_helper/internal/model"
	"teacher_helper/internal/repo"
	"time"

	"github.com/go-pdf/fpdf"
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
	e.GET("/students/privileges/:id/info", h.GetPrivilege)
	e.GET("/students/privileges/:id/edit", h.EditPrivilege)
	e.PUT("/students/privileges/:privilege_ID", h.UpdatePrivilege)
	e.DELETE("/students/privileges/:id", h.DeletePrivilege)
	e.POST("/students/privileges", h.CreatePrivilege)

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

	e.GET("/students/:id", h.GetStudent)
	e.GET("/students/:id/info", h.StudentInfo)
	e.GET("/students/:id/edit", h.StudentInfoEdit)

	e.GET("/students", h.GetStudents)
	e.GET("/students/search", h.SearchStudents)
	e.POST("/students", h.CreateStudent)
	e.DELETE("/students/:id", h.DeleteStudent)
	e.PUT("/students/:id", h.UpdateStudent)

	e.GET("/students/print/general/:id", h.PrintStudentInfo)
	e.GET("/students/print/medcard/:id", h.PrintMedicalCard)
	e.GET("/students/print/grades/:id", h.PrintGrades)
}

type StudentsPage struct {
	Title    string
	Links    []Link
	Students []model.Student
}

type StudentPage struct {
	Links       []Link
	Student     model.Student
	MedicalCard *model.MedicalCard
	Connections []model.FamilyConnection
	Privileges  []model.Privilege
	Stats       []model.StudentStats
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

	studentsPage := StudentsPage{
		Title:    "Учні",
		Links:    links,
		Students: students,
	}

	return c.Render(http.StatusOK, "students.html/base", studentsPage)
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
	}

	connections, err := h.studentRepo.GetConnections(id)
	if err != nil {
		log.Error(err)
	}

	privileges, err := h.studentRepo.GetPrivileges(id)
	if err != nil {
		log.Error(err)
	}

	stats, err := h.studentRepo.GetStatsForStudent(id)
	if err != nil {
		log.Error(err)
	}

	links := NewLinks(PageTypeStudents)

	student_page := StudentPage{
		Links:       links,
		Student:     student,
		MedicalCard: medicalCard,
		Connections: connections,
		Privileges:  privileges,
		Stats:       stats,
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
		log.Error(err)
		return err
	}

	if err := query.Validate(); err != nil {
		log.Error(err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	students, err := h.studentRepo.GetWithParams(*query)
	if err != nil {
		log.Error(err)
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

// Connection

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

// MedicalCard

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

	medicalCard.WithDefauls()
	err := h.studentRepo.CreateMedicalCard(medicalCard)
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Can't add MedicalCard"})
	}

	return c.Render(http.StatusCreated, "student.html/medical-card", medicalCard)
}

// Privilege

func (h *StudentHandler) GetPrivilege(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid student ID"})
	}

	privilege, err := h.studentRepo.GetPrivilegeByID(id)
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Can't get Privilege"})
	}

	return c.Render(http.StatusOK, "student.html/privilege", privilege)
}

func (h *StudentHandler) EditPrivilege(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid student ID"})
	}

	privilege, err := h.studentRepo.GetPrivilegeByID(id)
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Can't get privilege"})
	}

	return c.Render(http.StatusOK, "student.html/privilege-edit", privilege)
}

func (h *StudentHandler) UpdatePrivilege(c echo.Context) error {
	privilege := new(model.Privilege)
	if err := c.Bind(privilege); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid privilege data"})
	}

	if err := h.studentRepo.UpdatePrivilege(privilege); err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Can't update privilege"})
	}

	privilege, err := h.studentRepo.GetPrivilegeByID(int(privilege.PrivilegeID))
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Can't get privilege"})
	}

	return c.Render(http.StatusOK, "student.html/privilege", privilege)
}

func (h *StudentHandler) DeletePrivilege(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid privilege ID"})
	}

	if err := h.studentRepo.DeletePrivilege(id); err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Can't delete privilege"})
	}

	return c.NoContent(http.StatusAccepted)
}

func (h *StudentHandler) CreatePrivilege(c echo.Context) error {
	privilege := new(model.Privilege)
	if err := c.Bind(privilege); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid parent data"})
	}

	id, err := h.studentRepo.CreatePrivilege(privilege)
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Can't create privilege"})
	}

	privilege, err = h.studentRepo.GetPrivilegeByID(int(id))
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Can't get parent"})
	}

	return c.Render(http.StatusCreated, "student.html/privilege", privilege)
}

func (h *StudentHandler) PrintStudentInfo(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid student ID"})
	}

	student, err := h.studentRepo.GetByID(id)
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Can't get student"})
	}

	caretakers, err := h.studentRepo.GetCaretakers(id)
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Can't get CareTakers"})
	}

	pdf := fpdf.New("P", "mm", "A6", "")

	pdf.AddUTF8Font("Roboto", "", "roboto.ttf")
	pdf.SetFont("Roboto", "", 14)

	pdf.AddPage()
	pdf.Cell(0, 10, fmt.Sprintf("%s %s %s", student.Lastname, student.Firstname, student.Middlename))
	pdf.Ln(8)

	pdf.SetFont("Roboto", "", 12)
	var gender string
	if student.Gender == model.Male {
		gender = "Хлопець"
	} else {
		gender = "Дівчина"
	}

	pdf.Cell(0, 10, gender)
	pdf.Ln(8)
	pdf.Cell(0, 10, "Дата народження: "+student.Birthday.Format("2006-01-02"))
	pdf.Ln(8)
	pdf.Cell(0, 10, "Форма навчання: "+string(student.FormOfEducation))
	pdf.Ln(8)
	pdf.Cell(0, 10, "№ особової справи: "+student.PersonalFileNumber)
	pdf.Ln(8)
	if student.Note != nil {
		pdf.Cell(0, 10, *student.Note)
		pdf.Ln(8)
	}

	pdf.Ln(16)
	pdf.Cell(0, 10, "Опікуни:")
	pdf.SetFont("Roboto", "", 8)

	for _, caretaker := range caretakers {
		pdf.Ln(6)
		pdf.Cell(0, 10, fmt.Sprintf("%s %s %s", caretaker.Lastname, caretaker.Firstname, caretaker.Middlename))
		pdf.Ln(4)
		pdf.Cell(0, 10, caretaker.Phone)
		pdf.Ln(4)
		pdf.Cell(0, 10, caretaker.Email)
	}

	err = pdf.OutputFileAndClose("student_info.pdf")
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Can't create PDF"})
	}

	return c.File("student_info.pdf")
}

func (h *StudentHandler) PrintMedicalCard(c echo.Context) error {
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
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Can't get MedicalCard"})
	}

	pdf := fpdf.New("P", "mm", "A6", "")

	pdf.AddUTF8Font("Roboto", "", "roboto.ttf")
	pdf.SetFont("Roboto", "", 16)

	pdf.AddPage()
	pdf.Cell(0, 10, fmt.Sprintf("%s %s %s", student.Lastname, student.Firstname, student.Middlename))
	pdf.Ln(8)

	pdf.SetFont("Roboto", "", 12)
	pdf.Cell(0, 10, fmt.Sprintf("Група крові: %s", string(medicalCard.BloodGroup)))
	pdf.Ln(8)
	pdf.Cell(0, 10, fmt.Sprintf("Резус-фактор: %s", string(medicalCard.RhFactor)))
	pdf.Ln(8)
	pdf.Cell(0, 10, fmt.Sprintf("Вага: %.1f кг", medicalCard.Weight))
	pdf.Ln(8)
	pdf.Cell(0, 10, fmt.Sprintf("Зріст: %.1f см", medicalCard.Height))
	pdf.Ln(8)
	pdf.Cell(0, 10, fmt.Sprintf("Група здоров'я: %s", medicalCard.HealthGroup))
	pdf.Ln(16)
	pdf.Cell(0, 10, "Медичні огляди:")
	pdf.Ln(8)
	pdf.Cell(0, 10, fmt.Sprintf("Попередній - %s", (*medicalCard).LastInspectionDate.Format("2006-01-02")))
	pdf.Ln(8)
	pdf.Cell(0, 10, fmt.Sprintf("Наступний - %s", (*medicalCard).NextInspectionDate.Format("2006-01-02")))

	err = pdf.OutputFileAndClose("medical_card.pdf")
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Can't create PDF"})
	}

	return c.File("medical_card.pdf")
}

func (h *StudentHandler) PrintGrades(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid student ID"})
	}

	student, err := h.studentRepo.GetByID(id)
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Can't get student"})
	}

	stats, err := h.studentRepo.GetStatsForStudent(id)
	if err != nil {
		log.Error(err)
	}

	pdf := fpdf.New("L", "mm", "A5", "")

	pdf.AddUTF8Font("Roboto", "", "roboto.ttf")
	pdf.SetFont("Roboto", "", 16)

	pdf.AddPage()

	pdf.Cell(0, 10, fmt.Sprintf("%s %s %s", student.Lastname, student.Firstname, student.Middlename))
	pdf.Ln(8)

	pdf.Cell(0, 10, "Оцінки:")
	pdf.Ln(8)

	pdf.SetFont("Roboto", "", 12)
	// Розміри колонок
	columnWidth := 90.0 // Ширина кожної колонки (A4 ширина ≈ 210 мм, дві колонки по 90 мм + відступи)
	margin := 10.0      // Відступи між колонками

	// Координати початку
	leftColumnX := margin
	rightColumnX := margin + columnWidth + 10 // Відступ між колонками

	// Висота рядка
	lineHeight := 8.0

	// Малювання двох колонок
	for i, stat := range stats {
		// Визначаємо колонку: ліва або права
		if i%2 == 0 {
			// Ліва колонка
			pdf.SetXY(leftColumnX, margin+float64(i/2)*lineHeight+20)
		} else {
			// Права колонка
			pdf.SetXY(rightColumnX, margin+float64(i/2)*lineHeight+20)
		}

		pdf.CellFormat(80, 10, stat.SubjectTitle, "0", 0, "L", false, 0, "")               // Перший стовпець (80 мм)
		pdf.CellFormat(30, 10, fmt.Sprintf("%.2f", stat.Grade), "0", 0, "L", false, 0, "") // Другий стовпець (30 мм)
	}

	pdf.Ln(20)
	pdf.Cell(0, 10, fmt.Sprintf("Станом на: %s", time.Now().Format("2006-01-02")))

	err = pdf.OutputFileAndClose("grades.pdf")
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Can't create PDF"})
	}

	return c.File("grades.pdf")
}
