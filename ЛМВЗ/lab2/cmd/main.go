package main

import (
	// "time"
	// "context"
	// "database/sql"
	"html/template"
	"io"

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
	tmpl := template.New("").Funcs(template.FuncMap{})

	tmpl = template.Must(tmpl.ParseGlob("views/*.html"))

	return &Templates{templates: tmpl}
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Renderer = newTemplate()

	// db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/foodcourt?parseTime=true")
	// if err != nil {
	// 	panic(err)
	// }
	//
	// queries := sql_queries.New(db)

	e.Static("/css", "static/css")
	e.Static("/images", "static/images")
	e.Static("/scripts", "static/scripts")

	e.GET("/", func(c echo.Context) error {
		return c.Render(200, "main", nil)
	})

	e.Logger.Fatal(e.Start(":6969"))
}
