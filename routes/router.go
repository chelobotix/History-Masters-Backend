package routes

import (
	"myapp/internal/handlers"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func ConfigRoutes(e *echo.Echo, db *gorm.DB, logger *zap.Logger) {
	h := handlers.NewHandlersInit(db)
	api := e.Group("/api/v1")

	// Health
	api.GET("/health", h.Health.ServerHealth)

	// Event
	api.GET("/events", h.Event.GetAll)

}
