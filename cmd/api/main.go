package main

import (
	"myapp/routes"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	routes.ConfigRoutes(e)

	e.GET("/", home)

	e.Logger.Fatal(e.Start(":4000"))
}

func home(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]any{
		"state": "token generated",
	})
}
