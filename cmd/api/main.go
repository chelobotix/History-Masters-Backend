package main

import (
	"myapp/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	routes.ConfigRoutes(e)

	e.Logger.Fatal(e.Start(":4000"))
}
