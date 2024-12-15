package service

import (
	"teacher_helper/internal/model"
	"teacher_helper/internal/repo"
	"time"
)

type Day struct {
	Year     int
	Month    int
	Day      int
	Lessons  []model.Lesson
	Selected bool
}

type Calendar struct {
	CurrentMonth     int
	CurrentMonthName string
	CurrentYear      int
	CurrentDay       int
	CurrentDate      model.Date
	Days             []Day
}

func NewCalendar(now time.Time, repo repo.LessonRepo, filters repo.Filters) (*Calendar, error) {
	year, _ := now.ISOWeek()
	today := now.Day()
	month := now.Month()

	currentMonthDays := daysInMonth(now.Month(), year)

	var previousMonthDays int

	if previousMonthNumber := now.Month() - 1; previousMonthNumber < 1 {
		previousMonthDays = daysInMonth(12, year-1)
	} else {
		previousMonthDays = daysInMonth(previousMonthNumber, year)
	}

	startDaysFrom := startDaysFrom(now)

	days := make([]Day, 42)

	for i := 0; i < 42; i++ {
		dayNumber := startDaysFrom + i
		dayMonth := month

		if dayNumber < 1 {
			dayNumber = previousMonthDays + dayNumber
			dayMonth = month - 1
		} else if dayNumber > currentMonthDays {
			dayNumber = dayNumber - currentMonthDays
			dayMonth = month + 1
		}

		lessons, err := repo.GetLessonsAtDate(year, dayMonth, dayNumber, filters)
		if err != nil {
			return nil, err
		}

		days[i] = Day{
			Year:     year,
			Month:    int(dayMonth),
			Day:      dayNumber,
			Selected: dayNumber == today && dayMonth == month,
			Lessons:  lessons,
		}
	}

	return &Calendar{
		CurrentMonth:     int(month),
		CurrentMonthName: monthName(month),
		CurrentYear:      year,
		CurrentDay:       today,
		CurrentDate:      model.Date(now),
		Days:             days,
	}, nil
}

func daysInMonth(m time.Month, year int) int {
	return time.Date(year, m+1, 0, 0, 0, 0, 0, time.UTC).Day()
}

func startDaysFrom(now time.Time) int {
	// year, week := now.ISOWeek()
	monthStart := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.UTC)
	weekday := monthStart.Weekday()

	if weekday == 0 {
		weekday = 7
	}

	return monthStart.Day() - int(weekday) + 1
}

func monthName(month time.Month) string {
	switch month {
	case 1:
		return "Січень"
	case 2:
		return "Лютий"
	case 3:
		return "Березень"
	case 4:
		return "Квітень"
	case 5:
		return "Травень"
	case 6:
		return "Червень"
	case 7:
		return "Липень"
	case 8:
		return "Серпень"
	case 9:
		return "Вересень"
	case 10:
		return "Жовтень"
	case 11:
		return "Листопад"
	case 12:
		return "Грудень"
	default:
		return ""
	}
}
