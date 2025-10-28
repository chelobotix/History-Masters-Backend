package seeds

import (
	"encoding/json"
	"log"
	"myapp/internal/models"
	"os"

	"gorm.io/gorm"
)

func CountrySeeds(db *gorm.DB) error {
	var count int64

	db.Model(&models.Country{}).Count(&count)
	if count > 0 {
		log.Println("Skipping fixture loading: country already exists.")
		return nil
	}

	data, err := os.ReadFile("internal/database/seeds/json/countries.json")
	if err != nil {
		panic(err)
	}

	var countries []models.Country
	if err := json.Unmarshal(data, &countries); err != nil {
		panic(err)
	}

	if err := db.Create(&countries).Error; err != nil {
		return err
	} else {
		log.Println("Fixture loading: country seed created.")
	}

	return nil
}
