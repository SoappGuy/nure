package handler

import (
	"net/http"
	"slices"
	"strconv"
	"strings"
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

	e.DELETE("/mark/:id", h.DeleteMark)
	e.POST("/mark/:lesson_id/:student_id", h.CreateMark)

	e.PUT("/attendance/:id/flip", h.FlipAttendance)
	e.GET("/attendance/:id/edit", h.AttendanceInfoEdit)
	e.GET("/attendance/:id/info", h.AttendanceInfo)
	e.PUT("/attendance/:attendance_ID", h.UpdateAttendance)
}

func (h *LessonHandler) GetLessons(c echo.Context) error {
	filters := repo.Filters{}

	today := time.Now()
	calendar, err := service.NewCalendar(today, *h.lessonRepo, filters)
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, "Internal server error")
	}

	lessons, err := h.lessonRepo.GetLessonsAtDate(today.Year(), today.Month(), today.Day(), filters)
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

	filters := repo.Filters{}
	params := c.QueryParams()
	if len(params) > 0 {
		if subjectIDs, exist := params["subject_ID"]; exist {
			if slices.Contains(subjectIDs, "-1") {
				filters.SubjectIDs = ""
			} else {
				filters.SubjectIDs = strings.Join(subjectIDs, ",")
			}
		}
	}

	calendar, err := service.NewCalendar(today, *h.lessonRepo, filters)
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, "Internal server error")
	}

	lessons, err := h.lessonRepo.GetLessonsAtDate(today.Year(), today.Month(), today.Day(), filters)
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

	isHtmx := c.Request().Header.Get("HX-Request") == "true"
	isHistoryRestore := c.Request().Header.Get("HX-History-Restore-Request") == "true"

	if isHtmx && !isHistoryRestore {
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

	details, err := service.GetLessonDetails(int64(id), h.lessonRepo)
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Can't get lesson details"})
	}

	links := NewLinks(PageTypeLessons)

	lessonPage := LessonPage{
		Title:   "Уроки",
		Links:   links,
		Lesson:  lesson,
		Details: details,
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

func (h *LessonHandler) CreateMark(c echo.Context) error {
	mark := new(model.Mark)
	if err := c.Bind(mark); err != nil {
		log.Error(err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid mark data"})
	}

	err := h.lessonRepo.CreateMark(mark)
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Can't create mark"})
	}

	return c.Render(http.StatusCreated, "lesson.html/mark", mark)
}

func (h *LessonHandler) DeleteMark(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid Mark ID"})
	}

	if err := h.lessonRepo.DeleteMark(id); err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Can't delete Mark"})
	}

	return c.NoContent(http.StatusAccepted)
}

func (h *LessonHandler) FlipAttendance(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid Attendance id"})
	}

	attendance, err := h.lessonRepo.FlipAttendanceWithID(id)
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Can't flip Attendance with given id"})
	}

	if attendance.Attendance == model.Absent {
		return c.Render(http.StatusOK, "lesson.html/attendance-absent", attendance)
	} else {
		return c.Render(http.StatusOK, "lesson.html/attendance-present", attendance)
	}
}

func (h *LessonHandler) AttendanceInfo(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid Attendance id"})
	}

	attendance, err := h.lessonRepo.GetAttendanceByID(id)
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Can't get Attendance with given ID"})
	}

	if attendance.Attendance == model.Absent {
		return c.Render(http.StatusOK, "lesson.html/attendance-absent", attendance)
	} else {
		return c.Render(http.StatusOK, "lesson.html/attendance-present", attendance)
	}
}

func (h *LessonHandler) AttendanceInfoEdit(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid Attendance id"})
	}

	attendance, err := h.lessonRepo.GetAttendanceByID(id)
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Can't get Attendance with given ID"})
	}

	return c.Render(http.StatusOK, "lesson.html/attendance-edit", attendance)
}

func (h *LessonHandler) UpdateAttendance(c echo.Context) error {
	attendance := new(model.Attendance)
	if err := c.Bind(attendance); err != nil {
		log.Error(err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid Attendance data"})
	}

	if err := h.lessonRepo.UpdateAttendance(attendance); err != nil {
		log.Error(err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Can't update Attendance"})
	}

	if attendance.Attendance == model.Absent {
		return c.Render(http.StatusOK, "lesson.html/attendance-absent", attendance)
	} else {
		return c.Render(http.StatusOK, "lesson.html/attendance-present", attendance)
	}
}

type LessonsPage struct {
	Title    string
	Links    []Link
	Calendar service.Calendar
	Lessons  []model.Lesson
}

type LessonPage struct {
	Title   string
	Links   []Link
	Lesson  model.Lesson
	Details []service.StudentAtLesson
}
