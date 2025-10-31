package handlers

import (
	"github.com/labstack/echo/v4"
)

func ErrorHandler(err error, status int, c echo.Context) error {
	return c.JSON(status, map[string]any{
		"error":   true,
		"details": err.Error(),
	})
}
