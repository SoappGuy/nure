package handler

import (
	"net/http"

	"teacher_helper/internal/model"
	"teacher_helper/internal/repo"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type CaretakerHandler struct {
	caretakerRepo *repo.CaretakerRepo
}

func NewCaretakerHandler(caretakerRepo *repo.CaretakerRepo) *CaretakerHandler {
	return &CaretakerHandler{caretakerRepo: caretakerRepo}
}

func (h *CaretakerHandler) RegisterRoutes(e *echo.Echo) {
	e.GET("/caretakers", h.GetAllCaretakers)
}

type CaretakersPage struct {
	Title      string
	Links      []Link
	Caretakers []model.Caretaker
}

func (h *CaretakerHandler) GetAllCaretakers(c echo.Context) error {
	query_params := repo.QueryParams{
		Query:        "%",
		OrderBy:      "lastname",
		IsDescending: false,
	}

	caretakers, err := h.caretakerRepo.GetWithParams(query_params)
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Can't get caretakers"})
	}

	links := NewLinks(PageTypeCaretakers)

	caretakers_page := CaretakersPage{
		Title:      "Опікуни",
		Links:      links,
		Caretakers: caretakers,
	}

	return c.Render(http.StatusOK, "caretakers.html/base", caretakers_page)
}
