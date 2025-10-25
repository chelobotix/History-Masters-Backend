package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type EventHandler interface {
	GetAll(c echo.Context) error
}

type eventHandler struct{}

func NewEventHandler() EventHandler {
	return &eventHandler{}
}

func (eh *eventHandler) GetAll(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]any{
		"status": "ok",
	})
}
