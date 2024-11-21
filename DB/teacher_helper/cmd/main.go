package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"log"
	"teacher_helper/internal/handler"
	"teacher_helper/internal/repo"
	"teacher_helper/internal/template"
)

func main() {
	e := echo.New()

	DefaultSkipper := middleware.DefaultSkipper
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Skipper: DefaultSkipper,
		Format: "------------------------------------------------------" +
			"\n" + `time: ${time_custom}` +
			"\n" + `method: ${method}` +
			"\n" + `uri: ${uri}` +
			"\n" + `status: ${status}` +
			"\n" + `error: ${error}` +
			// "\n" + `id: ${id}` +
			// "\n" + `remote_ip: ${remote_ip}` +
			// "\n" + `host: ${host}` +
			// "\n" + `user_agent: ${user_agent}` +
			// "\n" + `latency: ${latency}` +
			// "\n" + `latency_human: ${latency_human}` +
			// "\n" + `bytes_in/bytes_out: ${bytes_in}/${bytes_out}` +
			"\n------------------------------------------------------\n",
		CustomTimeFormat: "15:04:05",
	}))

	renderer := template.NewTemplates()
	e.Renderer = renderer

	// log.Printf("Loaded templates: %v", renderer.Templates.DefinedTemplates())

	dsn := "root:password@tcp(127.0.0.1:3306)/class_manager?parseTime=true"
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		log.Fatalf("Error encountered while connecting to DB: %v", err)
	}
	defer db.Close()

	app := NewApp(db)

	e.Static("/css", "static/css")
	e.Static("/images", "static/img")
	e.Static("/scripts", "static/js")

	app.StudentHandler.RegisterRoutes(e)
	app.CaretakerHandler.RegisterRoutes(e)

	log.Println("Started on :6969")
	e.Start(":6969")
}

type App struct {
	DB *sqlx.DB

	StudentHandler   *handler.StudentHandler
	CaretakerHandler *handler.CaretakerHandler
}

func NewApp(db *sqlx.DB) *App {
	studentRepo := repo.NewStudentRepo(db)
	caretakerRepo := repo.NewCaretakerRepo(db)

	studentHandler := handler.NewStudentHandler(studentRepo)
	caretakerHandler := handler.NewCaretakerHandler(caretakerRepo)

	return &App{
		DB:               db,
		StudentHandler:   studentHandler,
		CaretakerHandler: caretakerHandler,
	}
}
