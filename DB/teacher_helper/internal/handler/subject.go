package handler

import (
	"net/http"
	"strconv"
	"teacher_helper/internal/model"
	"teacher_helper/internal/repo"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type SubjectHandler struct {
	subjectRepo *repo.SubjectRepo
}

func NewSubjectHandler(subjectRepo *repo.SubjectRepo) *SubjectHandler {
	return &SubjectHandler{subjectRepo: subjectRepo}
}

func (h *SubjectHandler) RegisterRoutes(e *echo.Echo) {
	e.GET("/subjects", h.GetSubjects)
	e.GET("/subjects/search", h.SearchSubjects)
	e.POST("/subjects", h.CreateSubject)
	e.DELETE("/subjects/:id", h.DeleteSubject)
	e.PUT("/subjects/:id", h.UpdateSubject)

	e.GET("/subjects/select/:id", h.SubjectsSelect)

	e.GET("/subjects/search", h.SearchSubjects)
	e.POST("/subjects", h.CreateSubject)
	e.DELETE("/subjects/:id", h.DeleteSubject)
	e.PUT("/subjects/:id", h.UpdateSubject)

	e.GET("/subjects/:id", h.GetSubject)
	e.GET("/subjects/:id/info", h.SubjectInfo)
	e.GET("/subjects/:id/edit", h.SubjectInfoEdit)
}

type SubjectPage struct {
	Links   []Link
	Subject model.Subject
}

type SubjectsPage struct {
	Title    string
	Links    []Link
	Subjects []model.Subject
}

func (h *SubjectHandler) GetSubjects(c echo.Context) error {
	query_params := repo.QueryParams{
		Query:        "%",
		OrderBy:      "title",
		IsDescending: false,
	}

	subjects, err := h.subjectRepo.GetWithParams(query_params)
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Can't get subjects"})
	}

	links := NewLinks(PageTypeSubjects)

	subjectsPage := SubjectsPage{
		Title:    "Subjects",
		Links:    links,
		Subjects: subjects,
	}

	return c.Render(http.StatusOK, "subjects.html/base", subjectsPage)
}

func (h *SubjectHandler) GetSubject(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid subject ID"})
	}

	subject, err := h.subjectRepo.GetByID(id)
	if err != nil {
		log.Error(err)
	}

	links := NewLinks(PageTypeSubjects)

	subjectPage := SubjectPage{
		Links:   links,
		Subject: *subject,
	}

	return c.Render(http.StatusOK, "subject.html/base", subjectPage)

}

func (h *SubjectHandler) SubjectInfoEdit(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid subject ID"})
	}

	subject, err := h.subjectRepo.GetByID(id)
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Can't get subject"})
	}

	return c.Render(http.StatusOK, "subject.html/subject-info-edit", subject)
}

func (h *SubjectHandler) SubjectInfo(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid subject ID"})
	}

	subject, err := h.subjectRepo.GetByID(id)
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Can't get subject"})
	}

	return c.Render(http.StatusOK, "subject.html/subject-info", subject)
}

func (h *SubjectHandler) SearchSubjects(c echo.Context) error {
	query := new(repo.QueryParams)
	if err := c.Bind(query); err != nil {
		return err
	}

	if err := query.Validate(repo.QueryParams{OrderBy: "title"}); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	subjects, err := h.subjectRepo.GetWithParams(*query)
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Can't get subjects"})
	}

	return c.Render(http.StatusOK, "subjects.html/subjects", subjects)
}

func (h *SubjectHandler) CreateSubject(c echo.Context) error {
	subject := new(model.Subject)
	if err := c.Bind(subject); err != nil {
		return err
	}

	err := h.subjectRepo.Create(subject)
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Can't create subject"})
	}

	return c.Render(http.StatusOK, "subjects.html/subject", subject)
}

func (h *SubjectHandler) DeleteSubject(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid subject ID"})
	}

	if err := h.subjectRepo.Delete(id); err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Can't delete subject"})
	}

	return c.NoContent(http.StatusAccepted)
}

func (h *SubjectHandler) UpdateSubject(c echo.Context) error {
	subject := new(model.Subject)
	if err := c.Bind(subject); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid subject data"})
	}

	if err := h.subjectRepo.Update(subject); err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Can't update subject"})
	}

	return c.Render(http.StatusOK, "subject.html/subject-info", subject)
}

func (h *SubjectHandler) SubjectsSelect(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid subject ID"})
	}

	subject := new(model.Subject)
	if id != -1 {
		subject, err = h.subjectRepo.GetByID(id)
		if err != nil {
			log.Error(err)
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Can't get subject"})
		}
	}

	subjects, err := h.subjectRepo.GetWithParams(repo.QueryParams{
		Query:        "%",
		OrderBy:      "title",
		IsDescending: false,
	})

	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Can't get subjects"})
	}

	subjectsSelect := map[any]any{
		"Selected": subject,
		"Subjects": subjects,
	}

	return c.Render(http.StatusOK, "subjects.html/subjects-select", subjectsSelect)
}
