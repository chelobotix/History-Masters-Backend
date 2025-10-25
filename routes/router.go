package routes

import (
	"myapp/internal/handlers"

	"github.com/labstack/echo/v4"
)

func ConfigRoutes(e *echo.Echo) {
	h := handlers.NewHandlersInit()
	api := e.Group("/api/v1")

	// Health
	api.GET("/health", h.Health.ServerHealth)

	// Event
	api.GET("/events", h.Event.GetAll)

}
