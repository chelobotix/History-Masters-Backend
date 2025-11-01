package routes

import (
	"myapp/config"
	"myapp/internal/handlers"
)

func ConfigRoutes(mainDependencies *config.MainDependencies) {
	h := handlers.NewHandlersInit(mainDependencies)
	api := mainDependencies.Echo.Group("/api/v1")

	// Root
	mainDependencies.Echo.GET("/", h.Health.ServerHealth)

	// Health
	api.GET("/health", h.Health.ServerHealth)

	// Figure
	api.GET("/figures", h.Figure.GetAll)
	api.GET("/figures/:id", h.Figure.GetById)
	api.POST("/figures", h.Figure.Create)

}
