package handler

type Link struct {
	Id     string
	Title  string
	Icon   string
	Path   string
	Active bool
}

type PageType string

const (
	PageTypeMain PageType = "main"

	PageTypeStudents PageType = "students"
	PageTypeStudent  PageType = "student"

	PageTypeCaretakers PageType = "caretakers"
	PageTypeCaretaker  PageType = "caretaker"

	PageTypeSubjects PageType = "subjects"
	PageTypeSubject  PageType = "subject"

	PageTypeLessons PageType = "lessons"
	PageTypeLesson  PageType = "lesson"

	PageTypeMarks PageType = "marks"
	PageTypeMark  PageType = "mark"
)

func NewLink(id, title, icon string, path string, active bool) Link {
	return Link{
		Id:     id,
		Title:  title,
		Icon:   icon,
		Path:   path,
		Active: active,
	}
}

func NewLinks(active PageType) []Link {
	return []Link{
		NewLink("home_link", "Головна", "home", "/", active == PageTypeMain),
		NewLink("students_link", "Учні", "groups", "/students", active == PageTypeStudents),
		NewLink("caretakers_link", "Опікуни", "family_restroom", "/caretakers", active == PageTypeCaretakers),
		NewLink("subjects_link", "Предмети", "history_edu", "/subjects", active == PageTypeSubjects),
		NewLink("lessons_link", "Уроки", "book", "/lessons", active == PageTypeLessons),
		NewLink("marks_link", "Оцінки", "hotel_class", "/marks", active == PageTypeMarks),
	}
}
