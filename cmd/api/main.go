package main

import (
	"log"
	"myapp/config"
	"myapp/routes"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	// Config
	db, err := config.NewConfig(e)
	if err != nil {
		log.Panic(err)
		return
	}

	routes.ConfigRoutes(e, db)

	e.Logger.Fatal(e.Start(os.Getenv("SERVER_PORT")))
}
