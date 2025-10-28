package seeds

import (
	"encoding/json"
	"log"
	"myapp/internal/models"
	"os"

	"gorm.io/gorm"
)

func AreaSeeds(db *gorm.DB) error {
	var count int64

	db.Model(&models.Area{}).Count(&count)
	if count > 0 {
		log.Println("Skipping fixture loading: area already exists.")
		return nil
	}

	data, err := os.ReadFile("internal/database/seeds/json/areas.json")
	if err != nil {
		panic(err)
	}

	var areas []models.Area
	if err := json.Unmarshal(data, &areas); err != nil {
		panic(err)
	}

	if err := db.Create(&areas).Error; err != nil {
		return err
	} else {
		log.Println("Fixture loading: area seed created.")
	}

	return nil
}
