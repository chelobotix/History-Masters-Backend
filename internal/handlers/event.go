package handlers

import (
	"net/http"

	"github.com/labstack/echo"
)

type EventHandler interface {
	GetEvents(c echo.Context) error
}

type eventHandler struct{}

func NewEventHandler() EventHandler {
	return &eventHandler{}
}

func (eh *eventHandler) GetEvents(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]any{
		"status": "ok",
	})
}
