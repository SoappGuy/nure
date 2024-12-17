package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"teacher_helper/internal/model"
	"teacher_helper/internal/repo"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func dump(data interface{}) {
	b, _ := json.MarshalIndent(data, "", "  ")
	fmt.Print(string(b))
}

type StatsHandler struct {
	statsRepo *repo.StatsRepo
}

func NewStatsHandler(statsRepo *repo.StatsRepo) *StatsHandler {
	return &StatsHandler{statsRepo: statsRepo}
}

type StatsPage struct {
	Title                string
	Links                []Link
	NextMedCheck         []model.NextMedCheck
	StudentsRating       []model.StudentsRating
	PrivilegeSchedule    []model.PrivilegeSchedule
	FormOfEducationCount []model.FormOfEducationCount
	GenderCount          []model.GenderCount
	HealthGroupCount     []model.HealthGroupCount
	BloodGroupCount      []model.BloodGroupCount
}

func (h *StatsHandler) RegisterRoutes(e *echo.Echo) {
	e.GET("/stats", h.GetPage)
}

func (h *StatsHandler) GetPage(c echo.Context) error {
	nextMedCheck, err := h.statsRepo.GetNextMedCheck()
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	studentsRating, err := h.statsRepo.GetStudentsRating()
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	privilegeSchedule, err := h.statsRepo.GetPrivilegeSchedule()
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	formOfEducationCount, err := h.statsRepo.GetFormOfEducationCount()
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	genderCount, err := h.statsRepo.GetGenderCount()
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	healthGroupCount, err := h.statsRepo.GetHealthGroupCount()
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	bloodGroupCount, err := h.statsRepo.GetBloodGroupCount()
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, err)
	}

	links := NewLinks(PageTypeStats)

	stats_page := StatsPage{
		Title:                "Stats",
		Links:                links,
		NextMedCheck:         nextMedCheck,
		StudentsRating:       studentsRating,
		PrivilegeSchedule:    privilegeSchedule,
		FormOfEducationCount: formOfEducationCount,
		GenderCount:          genderCount,
		HealthGroupCount:     healthGroupCount,
		BloodGroupCount:      bloodGroupCount,
	}

	// log.Printf("%#v", stats_page)
	// dump(stats_page)

	return c.Render(http.StatusOK, "stats.html/base", stats_page)
}
