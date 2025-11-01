package main

import (
	"log"
	"myapp/config"
	"myapp/routes"

	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

func main() {
	e := echo.New()

	// Initialize Config
	mainDependencies, err := config.NewConfig(e)
	if err != nil {
		log.Panic(err)
		return
	}

	routes.ConfigRoutes(mainDependencies)

	e.Logger.Fatal(e.Start(viper.GetString("server.port")))
}
