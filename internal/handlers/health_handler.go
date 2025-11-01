package handlers

import (
	"myapp/config"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type HealthHandler interface {
	ServerHealth(c echo.Context) error
}

type healthHandler struct {
	Logger *zap.Logger
}

func NewHealthHandler(mainDependencies *config.MainDependencies) HealthHandler {
	return &healthHandler{
		Logger: mainDependencies.Logger,
	}
}

func (h *healthHandler) ServerHealth(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]any{
		"server": "live",
	})
}
