package database

import (
	"fmt"
	"log"
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
		dbPort = "5432" // Default port if not specified
	}

	// Connect to DB
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC", dbHost, dbUser, dbPassword, dbName, dbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func AutoMigrate(db *gorm.DB) error {
	// err := db.AutoMigrate(
	// 	&models.Bank{},
	// 	&models.InvestmentType{},
	// 	&models.InterestType{},
	// 	&models.EconomyIndexer{},
	// 	&models.Investments{},
	// 	&models.Users{},
	// 	&models.Favorite{},
	// )
	// if err != nil {
	// 	return err
	// }

	return nil
}

func LoadFixtures(db *gorm.DB) error {
	// var count int64
	// db.Model(&models.EconomyIndexer{}).Count(&count)
	// if count > 0 {
	// 	log.Println("Skipping fixture loading: economy_indexers already exists.")
	// 	return nil
	// }

	// economyIndexers := []models.EconomyIndexer{
	// 	{Name: "CDI", Value: decimal.NewFromFloat(0.1088)},
	// 	{Name: "CDI_ACUMULADO", Value: decimal.NewFromFloat(0.1088)},
	// 	{Name: "SELIC", Value: decimal.NewFromFloat(0.1225)},
	// 	{Name: "IPCA", Value: decimal.NewFromFloat(0.0483)},
	// }

	// if err := db.Create(&economyIndexers).Error; err != nil {
	// 	return err
	// }

	return nil
}
