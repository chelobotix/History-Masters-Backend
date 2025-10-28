package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type FigureHandler interface {
	GetAll(c echo.Context) error
}

type figureHandler struct{}

func NewFigureHandler() FigureHandler {
	return &figureHandler{}
}

func (fh *figureHandler) GetAll(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]any{
		"status": "ok",
	})
}
