package handler

import (
	"net/http"

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

func (h *CaretakerHandler) GetAllCaretakers(c echo.Context) error {
	caretakers, err := h.caretakerRepo.GetAll()
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Can't get caretakers"})
	}
	return c.Render(http.StatusOK, "caretakers.html", caretakers)
}
