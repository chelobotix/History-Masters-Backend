package routes

import (
	"myapp/internal/handlers"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func ConfigRoutes(e *echo.Echo, db *gorm.DB) {
	h := handlers.NewHandlersInit(db)
	api := e.Group("/api/v1")

	// Health
	api.GET("/health", h.Health.ServerHealth)

	// Event
	api.GET("/events", h.Event.GetAll)

}
