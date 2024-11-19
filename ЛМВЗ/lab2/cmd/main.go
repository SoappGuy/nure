package main

import (
	"fmt"

	sql "github.com/jmoiron/sqlx"

	"html/template"
	"io"
	"math/rand"
	"strconv"
	"time"

	"foodcourt/cmd/sql_queries"

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
	tmpl := template.New("").Funcs(template.FuncMap{
		"version": func() string {
			return strconv.Itoa(int(time.Now().Unix()))
		},
		"seq": func(n int) []int {
			seq := make([]int, n)
			for i := range seq {
				seq[i] = i
			}
			return seq
		},
		"repeat": func(item interface{}, n int) []interface{} {
			seq := make([]interface{}, n)
			for i := range seq {
				seq[i] = item
			}
			return seq
		},
	})

	tmpl = template.Must(tmpl.ParseGlob("views/*.html"))

	return &Templates{templates: tmpl}
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
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
	e.Renderer = newTemplate()

	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/foodcourt?parseTime=true")
	if err != nil {
		panic(err)
	}

	queries := sql_queries.New(db)

	e.Static("/css", "static/css")
	e.Static("/images", "static/images")
	e.Static("/scripts", "static/scripts")

	e.GET("/", func(c echo.Context) error {
		return c.Render(200, "main", nil)
	})

	e.GET("/page/:page", func(c echo.Context) error {
		page := c.Param("page")
		return c.Render(200, page, nil)
	})

	e.GET("/n/food/:type", func(c echo.Context) error {
		foodType := c.Param("type")
		if foodType == "Всі" {
			foodType = "%"
		}

		id := c.FormValue("id")
		id += "%"

		foodstuff, err := queries.ListFoodStuff(c.Request().Context(), sql_queries.ListFoodStuffArgs{Type: foodType, ID: id})
		if err != nil {
			fmt.Printf("\n\n%v\n\n", err)
			return c.String(500, err.Error())
		}

		return c.Render(200, "foodstuff_n", foodstuff)
	})

	e.GET("/food/:type", func(c echo.Context) error {
		foodType := c.Param("type")
		if foodType == "Всі" {
			foodType = "%"
		}

		id := c.FormValue("id")
		id += "%"

		foodstuff, err := queries.ListFoodStuff(c.Request().Context(), sql_queries.ListFoodStuffArgs{Type: foodType, ID: id})
		if err != nil {
			fmt.Printf("\n\n%v\n\n", err)
			return c.String(500, err.Error())
		}

		return c.Render(200, "foodstuff", foodstuff)
	})

	e.GET("/food/print/:id", func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.String(400, "Invalid ID")
		}

		queries.IncrementFrequency(c.Request().Context(), int64(id))

		foodstuff, err := queries.GetFoodStuff(c.Request().Context(), int64(id))
		if err != nil {
			return c.String(500, err.Error())
		}

		weight := rand.Float64() * 2

		bill := Bill{
			Item:          *foodstuff,
			ShopAddress:   "вул. Шевченка, 1",
			DateOfPacking: time.Now().Format("2006-01-02 15:04:05"),
			DateToUse:     time.Now().AddDate(0, 0, 7).Format("2006-01-02 15:04:05"),
			Weight:        fmt.Sprintf("%.3f", weight),
			Price:         fmt.Sprintf("%.2f", weight*foodstuff.Price),
			Barcode:       "||| || ||||| ||| |",
		}

		return c.Render(200, "foodstuff_print", bill)
	})
	e.Logger.Fatal(e.Start(":6969"))
	fmt.Println("Server started at :6969")
}

type Bill struct {
	Item          sql_queries.Foodstuff
	ShopAddress   string
	DateOfPacking string
	DateToUse     string
	Weight        string
	Price         string
	Barcode       string
	Place         string
}
