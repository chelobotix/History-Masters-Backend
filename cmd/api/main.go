package main

import (
	"myapp/routes"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	routes.ConfigRoutes(e)

	e.Logger.Fatal(e.Start(os.Getenv("SERVER_PORT")))
}
