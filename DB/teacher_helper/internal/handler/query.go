package handler

import (
	"fmt"
	"net/http"
	"reflect"

	"teacher_helper/internal/repo"

	"github.com/labstack/echo/v4"
)

type QueryHandler struct {
	queryRepo *repo.QueryRepo
}

func NewQueryHandler(queryRepo *repo.QueryRepo) *QueryHandler {
	return &QueryHandler{queryRepo: queryRepo}
}

type QueryPage struct {
	Title string
	Links []Link
}

func (h *QueryHandler) RegisterRoutes(e *echo.Echo) {
	e.GET("/data", h.Query)
	e.GET("/query", h.GetPage)
}

func (h *QueryHandler) GetPage(c echo.Context) error {
	links := NewLinks(PageTypeQuery)

	query_page := QueryPage{
		Title: "Query",
		Links: links,
	}

	return c.Render(http.StatusOK, "query.html/base", query_page)
}

func (h *QueryHandler) Query(c echo.Context) error {
	query := c.FormValue("query")

	result, err := h.queryRepo.Query(query)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	for _, v := range result {
		convertStrings(v)
	}

	return c.JSON(http.StatusOK, result)
}

func convertStrings(in map[string]interface{}) {
	for k, v := range in {
		t := reflect.TypeOf(v)
		if t != nil {
			switch t.Kind() {
			case reflect.Slice:
				in[k] = fmt.Sprintf("%s", v)
			default:
				// do nothing
			}
		}
	}
}
