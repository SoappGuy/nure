package main

import (
	"class_manager/cmd/sql_queries"
	"context"
	"database/sql"
	"html/template"
	"io"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Templates struct {
	templates *template.Template
}

func (t *Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func newTemplate() *Templates {
	return &Templates{templates: template.Must(template.ParseGlob("views/*.html"))}
}

func main() {

	ctx := context.Background()

	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/class_manager?parseTime=true")
	if err != nil {
		panic(err)
	}

	queries := sql_queries.New(db)

	e := echo.New()
	e.Use(middleware.Logger())

	e.Renderer = newTemplate()

	state := NewApp()

	e.GET("/", func(c echo.Context) error {
		return c.Render(200, "index", state)
	})

	e.GET("/students", func(c echo.Context) error {
		students, err := queries.ListStudents(ctx)
		if err != nil {
			return err
		}

		if len(students) == 0 {
			return c.Render(404, "students-empty", students)
		}
		return c.Render(200, "students-table", students)
	})

	e.DELETE("/students", func(c echo.Context) error {
		students_selected := c.QueryParams()["student-selected"]
		return c.Render(200, "delete-students", students_selected)
	})

	e.POST("/students", func(c echo.Context) error {
		student := sql_queries.Student{
			StudentID:          0,
			Firstname:          "",
			Middlename:         "",
			Lastname:           "",
			Gender:             "",
			Birthday:           time.Time{},
			FormOfEducation:    "",
			PersonalFileNumber: "",
			Note:               sql.NullString{},
		}
		return c.Render(200, "oob-student", student)
	})

	e.Logger.Fatal(e.Start(":6969"))
}

func NewApp() AppState {
	return AppState{}
}

type AppState struct {
}
