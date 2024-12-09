package handler

import (
	"net/http"
	"strconv"
	"teacher_helper/internal/model"
	"teacher_helper/internal/repo"
	"teacher_helper/internal/service"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type LessonHandler struct {
	lessonRepo *repo.LessonRepo
}

func NewLessonHandler(lessonRepo *repo.LessonRepo) *LessonHandler {
	return &LessonHandler{lessonRepo: lessonRepo}
}

func (h *LessonHandler) RegisterRoutes(e *echo.Echo) {
	e.GET("/lessons", h.GetLessons)
	e.GET("/lessons/:year/:month/:day/:add", h.GetCallendar)
	e.POST("/lessons/:start_date", h.CreateLesson)

	e.GET("/lesson/:id", h.GetLesson)
	e.DELETE("/lesson/:id", h.DeleteLesson)
	e.GET("/lesson/:id/info", h.LessonInfo)
	e.GET("/lesson/:id/edit", h.LessonInfoEdit)
	e.PUT("/lesson/:id", h.UpdateLesson)
}

type LessonsPage struct {
	Title    string
	Links    []Link
	Calendar service.Calendar
	Lessons  []model.Lesson
}

type LessonPage struct {
	Title  string
	Links  []Link
	Lesson model.Lesson
}

func (h *LessonHandler) GetLessons(c echo.Context) error {
	today := time.Now()
	calendar, err := service.NewCalendar(today, *h.lessonRepo)
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, "Internal server error")
	}

	lessons, err := h.lessonRepo.GetLessonsAtDate(today.Year(), today.Month(), today.Day())
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, "Internal server error")
	}

	links := NewLinks(PageTypeLessons)

	lessonsPage := LessonsPage{
		Title:    "Уроки",
		Links:    links,
		Calendar: *calendar,
		Lessons:  lessons,
	}

	return c.Render(http.StatusOK, "lessons.html/base", lessonsPage)
}

func (h *LessonHandler) GetCallendar(c echo.Context) error {
	year, err := strconv.Atoi(c.Param("year"))
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, "Internal server error")
	}

	currMonth, err := strconv.Atoi(c.Param("month"))
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, "Internal server error")
	}

	add, err := strconv.Atoi(c.Param("add"))
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, "Internal server error")
	}

	var month time.Month
	if currMonth+add < 1 {
		month = time.Month(12)
		year--
	} else if currMonth+add > 12 {
		month = time.Month(1)
		year++
	} else {
		month = time.Month(currMonth + add)
	}

	currDay, err := strconv.Atoi(c.Param("day"))
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, "Internal server error")
	}

	today := time.Date(year, month, currDay, 0, 0, 0, 0, time.UTC)
	calendar, err := service.NewCalendar(today, *h.lessonRepo)
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, "Internal server error")
	}

	lessons, err := h.lessonRepo.GetLessonsAtDate(today.Year(), today.Month(), today.Day())
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, "Internal server error")
	}

	links := NewLinks(PageTypeLessons)

	lessonsPage := LessonsPage{
		Title:    "Уроки",
		Links:    links,
		Calendar: *calendar,
		Lessons:  lessons,
	}

	if c.Request().Header.Get("HX-Request") == "true" && !(c.Request().Header.Get("HX-History-Restore-Request") == "true") {
		return c.Render(http.StatusOK, "lessons.html/oob-content", lessonsPage)
	}

	return c.Render(http.StatusOK, "lessons.html/base", lessonsPage)
}

func (h *LessonHandler) CreateLesson(c echo.Context) error {
	lessonWithIDs := new(model.LessonWithIDs)
	if err := c.Bind(lessonWithIDs); err != nil {
		log.Error(err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid lesson data"})
	}

	id, err := h.lessonRepo.CreateLesson(lessonWithIDs)
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Can't create lesson"})
	}

	lesson, err := h.lessonRepo.GetLessonByID(id)
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Can't get lesson"})
	}

	return c.Render(http.StatusCreated, "lessons.html/lesson", lesson)
}

func (h *LessonHandler) GetLesson(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid lesson id"})
	}

	lesson, err := h.lessonRepo.GetLessonByID(int64(id))
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Can't get lesson"})
	}

	links := NewLinks(PageTypeLessons)

	lessonPage := LessonPage{
		Title:  "Уроки",
		Links:  links,
		Lesson: lesson,
	}
	return c.Render(http.StatusOK, "lesson.html/base", lessonPage)
}

func (h *LessonHandler) DeleteLesson(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid Lesson ID"})
	}

	if err := h.lessonRepo.DeleteLesson(id); err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Can't delete Lesson"})
	}

	return c.NoContent(http.StatusAccepted)
}

func (h *LessonHandler) LessonInfo(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid Lesson ID"})
	}

	lesson, err := h.lessonRepo.GetLessonByID(int64(id))
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Can't get Lesson"})
	}

	return c.Render(http.StatusOK, "lesson.html/lesson-info", lesson)
}

func (h *LessonHandler) LessonInfoEdit(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid Lesson ID"})
	}

	lesson, err := h.lessonRepo.GetLessonByID(int64(id))
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Can't get Lesson"})
	}

	return c.Render(http.StatusOK, "lesson.html/lesson-edit", lesson)
}

func (h *LessonHandler) UpdateLesson(c echo.Context) error {
	lessonWithIDs := new(model.LessonWithIDs)
	if err := c.Bind(lessonWithIDs); err != nil {
		log.Error(err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid lesson data"})
	}

	if err := h.lessonRepo.UpdateLesson(lessonWithIDs); err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Can't update lesson"})
	}

	lesson, err := h.lessonRepo.GetLessonByID(lessonWithIDs.LessonID)
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Can't get lesson"})
	}

	return c.Render(http.StatusOK, "lesson.html/lesson-info", lesson)
}
