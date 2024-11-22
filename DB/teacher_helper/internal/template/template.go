package template

import (
	"fmt"
	"html/template"
	"io"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
)

type Templates struct {
	templates map[string]*template.Template
}

func (t *Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	split := strings.Split(name, "/")

	category := split[0]
	name = split[1]

	tmpl, ok := t.templates[category]
	if !ok {
		err := fmt.Errorf("Template not found -> %s/%s", category, name)
		return err
	}

	return tmpl.ExecuteTemplate(w, name, data)

}

func NewTemplates() *Templates {
	funcs := template.FuncMap{
		"or":  func(a, b bool) bool { return a || b },
		"mod": func(a, b int) int { return a % b },
		"add": func(a, b int) int { return a + b },
		"seq": func(n int) []int {
			s := make([]int, n)
			for i := range s {
				s[i] = i + 1
			}
			return s
		},
		"initial": func(word string) string {
			if len(word) == 0 {
				return ""
			}
			return string([]rune(word)[0])
		},
		"version": func() string {
			return string(time.Now().Unix())
		},
	}

	templates := make(map[string]*template.Template)

	templates["students.html"] = template.Must(
		template.New("students.html").
			Funcs(funcs).
			ParseFiles("templates/base.html", "templates/pages/students.html"),
	)

	templates["student.html"] = template.Must(
		template.New("student.html").
			Funcs(funcs).
			ParseFiles("templates/base.html", "templates/pages/student.html"),
	)

	templates["caretakers.html"] = template.Must(
		template.New("caretakers.html").
			Funcs(funcs).
			ParseFiles("templates/base.html", "templates/pages/caretakers.html"),
	)

	templates["caretaker.html"] = template.Must(
		template.New("caretaker.html").
			Funcs(funcs).
			ParseFiles("templates/base.html", "templates/pages/caretaker.html"),
	)

	templates["query.html"] = template.Must(
		template.New("query.html").
			Funcs(funcs).
			ParseFiles("templates/base.html", "templates/pages/query.html"),
	)

	return &Templates{
		templates: templates,
	}
}
