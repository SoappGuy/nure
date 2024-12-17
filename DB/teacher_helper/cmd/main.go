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
			"\n------------------------------------------------------\n",
		CustomTimeFormat: "15:04:05",
	}))

	renderer := template.NewTemplates()
	e.Renderer = renderer

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
	e.Static("/dist", "static/dist")

	e.GET("/", func(c echo.Context) error {
		return c.Render(200, "main.html/base", nil)
	})
	app.StudentHandler.RegisterRoutes(e)
	app.CaretakerHandler.RegisterRoutes(e)
	app.QueryHandler.RegisterRoutes(e)
	app.SubjectHandler.RegisterRoutes(e)
	app.LessonHandler.RegisterRoutes(e)
	app.StatsHandler.RegisterRoutes(e)

	log.Println("Started on :6969")
	e.Start(":6969")
}

type App struct {
	DB *sqlx.DB

	StudentHandler   *handler.StudentHandler
	CaretakerHandler *handler.CaretakerHandler
	QueryHandler     *handler.QueryHandler
	SubjectHandler   *handler.SubjectHandler
	LessonHandler    *handler.LessonHandler
	StatsHandler     *handler.StatsHandler
}

func NewApp(db *sqlx.DB) *App {
	studentRepo := repo.NewStudentRepo(db)
	caretakerRepo := repo.NewCaretakerRepo(db)

	studentHandler := handler.NewStudentHandler(studentRepo)
	caretakerHandler := handler.NewCaretakerHandler(caretakerRepo)

	queryRepo := repo.NewQueryRepo(db)
	queryHandler := handler.NewQueryHandler(queryRepo)

	subjectRepo := repo.NewSubjectRepo(db)
	subjectHandler := handler.NewSubjectHandler(subjectRepo)

	lessonRepo := repo.NewLessonRepo(db)
	lessonHandler := handler.NewLessonHandler(lessonRepo)

	statsRepo := repo.NewStatsRepo(db)
	statsHandler := handler.NewStatsHandler(statsRepo)

	return &App{
		DB:               db,
		StudentHandler:   studentHandler,
		CaretakerHandler: caretakerHandler,
		QueryHandler:     queryHandler,
		SubjectHandler:   subjectHandler,
		LessonHandler:    lessonHandler,
		StatsHandler:     statsHandler,
	}
}
