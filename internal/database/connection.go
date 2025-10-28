package database

import (
	"fmt"
	"log"
	"myapp/internal/database/migrations"
	"myapp/internal/database/seeds"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewConnection() (*gorm.DB, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Warning: .env file not found, skipping load")
	}

	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	if dbPort == "" {
		dbPort = "5432"
	}

	// Connect to DB
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC", dbHost, dbUser, dbPassword, dbName, dbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Migrations
	err = migrations.Migrate(db)
	if err != nil {
		return nil, err
	}

	// Seeds
	loadSeeds(db)

	return db, nil
}

func loadSeeds(db *gorm.DB) {
	seeds.CountrySeeds(db)
	seeds.YearSeeds(db)
	seeds.HistoricalEraSeeds(db)
	seeds.ProfessionsSeeds(db)
	seeds.AreaSeeds(db)
}
