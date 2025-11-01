package handlers

import (
	"github.com/labstack/echo/v4"
)

func ErrorHandler(c echo.Context, err error, status int) error {
	return c.JSON(status, map[string]any{
		"error":   true,
		"details": err.Error(),
	})
}
