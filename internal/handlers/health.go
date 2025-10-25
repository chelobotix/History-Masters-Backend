package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type HealthHandler interface {
	ServerHealth(c echo.Context) error
}

type healthHandler struct{}

func NewHealthHandler() HealthHandler {
	return &healthHandler{}
}

func (h *healthHandler) ServerHealth(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]any{
		"server": "live",
	})
}
