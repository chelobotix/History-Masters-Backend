package routes

import (
	"myapp/internal/handlers"

	"github.com/labstack/echo"
)

type HandlerRegistry struct {
	Event handlers.EventHandler
}

func newHandlers() *handlers_struct {
	return &handlers_struct{
		Event: handlers.NewEventHandler(),
	}
}

func ConfigRoutes(e *echo.Echo) {
	h := newHandlers()
	api := e.Group("/api/v1")

	// Event
	api.GET("/events", h.Event.GetEvents)

}
