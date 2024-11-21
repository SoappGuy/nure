package handler

import (
	"net/http"
	"strconv"

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
	e.GET("/caretakers/search", h.SearchCaretakers)
	e.POST("/caretakers", h.CreateCaretaker)
	e.DELETE("/caretakers/:id", h.DeleteCaretaker)
	e.PUT("/caretakers/:id", h.UpdateCaretaker)

	e.GET("/caretakers/:id", h.GetCaretaker)
	e.GET("/caretakers/:id/info", h.CaretakerInfo)
	e.GET("/caretakers/:id/edit", h.CaretakerInfoEdit)

	e.GET("/caretakers/familyconnections/:id/edit", h.EditConnection)
	e.PUT("/caretakers/familyconnections/:id", h.UpdateConnection)
	e.DELETE("/caretakers/familyconnections/:id", h.DeleteConnection)
	e.POST("/caretakers/familyconnections", h.CreateConnection)
}

type CaretakersPage struct {
	Title      string
	Links      []Link
	Caretakers []model.Caretaker
}

type CaretakerPage struct {
	Links       []Link
	Caretaker   model.Caretaker
	Connections []model.FamilyConnection
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

func (h *CaretakerHandler) GetCaretaker(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid caretaker ID"})
	}

	caretaker, err := h.caretakerRepo.GetByID(id)
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Can't get caretaker"})
	}

	links := NewLinks(PageTypeCaretakers)

	connections, err := h.caretakerRepo.GetConnections(id)

	caretaker_page := CaretakerPage{
		Links:       links,
		Caretaker:   caretaker,
		Connections: connections,
	}

	return c.Render(http.StatusOK, "caretaker.html/base", caretaker_page)
}

func (h *CaretakerHandler) CaretakerInfoEdit(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid caretaker ID"})
	}

	caretaker, err := h.caretakerRepo.GetByID(id)
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Can't get caretaker"})
	}

	return c.Render(http.StatusOK, "caretaker.html/caretaker-info-edit", caretaker)
}

func (h *CaretakerHandler) CaretakerInfo(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid caretaker ID"})
	}

	caretaker, err := h.caretakerRepo.GetByID(id)
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Can't get caretaker"})
	}

	return c.Render(http.StatusOK, "caretaker.html/caretaker-info", caretaker)
}

func (h *CaretakerHandler) SearchCaretakers(c echo.Context) error {
	query := new(repo.QueryParams)
	if err := c.Bind(query); err != nil {
		return err
	}

	if err := query.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	caretakers, err := h.caretakerRepo.GetWithParams(*query)
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Can't get caretakers"})
	}

	return c.Render(http.StatusOK, "caretakers.html/caretakers", caretakers)
}

func (h *CaretakerHandler) CreateCaretaker(c echo.Context) error {
	caretaker := new(model.Caretaker)
	if err := c.Bind(caretaker); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid caretaker data"})
	}

	if err := h.caretakerRepo.Create(caretaker); err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Can't create caretaker"})
	}

	return c.Render(http.StatusCreated, "caretakers.html/caretaker", caretaker)
}

func (h *CaretakerHandler) DeleteCaretaker(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid caretaker ID"})
	}

	if err := h.caretakerRepo.Delete(id); err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Can't delete caretaker"})
	}

	return c.NoContent(http.StatusAccepted)
}

func (h *CaretakerHandler) UpdateCaretaker(c echo.Context) error {
	caretaker := new(model.Caretaker)
	if err := c.Bind(caretaker); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid caretaker data"})
	}

	if err := h.caretakerRepo.Update(caretaker); err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Can't update caretaker"})
	}

	return c.Render(http.StatusOK, "caretaker.html/caretaker-info", caretaker)
}

func (h *CaretakerHandler) EditConnection(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid caretaker ID"})
	}

	kid, err := h.caretakerRepo.GetConnectionByID(id)
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Can't get kid"})
	}

	return c.Render(http.StatusOK, "caretaker.html/connection-edit", kid)
}

func (h *CaretakerHandler) UpdateConnection(c echo.Context) error {
	kid := new(model.FamilyConnectionWithIDs)
	if err := c.Bind(kid); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid kid data"})
	}

	if err := h.caretakerRepo.UpdateConnection(kid); err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Can't update kid"})
	}

	return c.Render(http.StatusOK, "caretaker.html/connection", kid)
}

func (h *CaretakerHandler) DeleteConnection(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid kid ID"})
	}

	if err := h.caretakerRepo.DeleteConnection(id); err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Can't delete kid"})
	}

	return c.NoContent(http.StatusAccepted)
}

func (h *CaretakerHandler) CreateConnection(c echo.Context) error {
	kid := new(model.FamilyConnectionWithIDs)
	if err := c.Bind(kid); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid kid data"})
	}

	id, err := h.caretakerRepo.CreateConnection(kid)
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Can't create kid"})
	}

	connection, err := h.caretakerRepo.GetConnectionByID(int(id))
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Can't get kid"})
	}

	return c.Render(http.StatusCreated, "caretaker.html/connection", connection)
}
