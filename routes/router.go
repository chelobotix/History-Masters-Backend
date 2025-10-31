package routes

import (
	"myapp/config"
	"myapp/internal/handlers"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func ConfigRoutes(e *echo.Echo, db *gorm.DB, cfg *config.Config, logger *zap.Logger) {
	h := handlers.NewHandlersInit(db, logger)
	api := e.Group("/api/v1")

	// Root
	e.GET("/", h.Health.ServerHealth)

	// Health
	api.GET("/health", h.Health.ServerHealth)

	// Figure
	api.GET("/figures", h.Figure.GetAll)
	api.GET("/figures/:id", h.Figure.GetById)

}
