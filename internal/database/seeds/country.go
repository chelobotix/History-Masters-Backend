package seeds

import (
	"encoding/json"
	"log"
	"myapp/internal/models"
	"os"

	"gorm.io/gorm"
)

func CountrySeeds(db *gorm.DB) {
	var count int64

	db.Model(&models.Country{}).Count(&count)
	if count > 0 {
		log.Println("Skipping fixture loading: country already exists.")
		return
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
		panic(err)
	} else {
		log.Println("Fixture loading: country seed created.")
	}
}
