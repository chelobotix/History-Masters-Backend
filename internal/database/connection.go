package database

import (
	"fmt"
	"myapp/internal/database/migrations"
	"myapp/internal/database/seeds"

	_ "github.com/joho/godotenv/autoload"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
}

func NewConnection() (*gorm.DB, error) {
	dbName := viper.GetString("database.name")
	dbUser := viper.GetString("database.user")
	dbPassword := viper.GetString("database.password")
	dbHost := viper.GetString("database.host")
	dbPort := viper.GetInt("database.port")

	// Connect to DB
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=UTC", dbHost, dbUser, dbPassword, dbName, dbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to Data Base Successfully!")

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
