package main

import (
	"log"
	"myapp/config"
	"myapp/routes"

	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// Config
	db, logger, err := config.NewConfig(e)
	if err != nil {
		log.Panic(err)
		return
	}

	routes.ConfigRoutes(e, db, logger)

	e.Logger.Fatal(e.Start(":4000"))
}
