package seeds

import (
	"encoding/json"
	"log"
	"myapp/internal/models"
	"os"

	"gorm.io/gorm"
)

func ProfessionsSeeds(db *gorm.DB) error {
	var count int64

	db.Model(&models.Profession{}).Count(&count)
	if count > 0 {
		log.Println("Skipping fixture loading: profession already exists.")
		return nil
	}

	data, err := os.ReadFile("internal/database/seeds/json/professions.json")
	if err != nil {
		panic(err)
	}

	var professions []models.Profession
	if err := json.Unmarshal(data, &professions); err != nil {
		panic(err)
	}

	if err := db.Create(&professions).Error; err != nil {
		return err
	} else {
		log.Println("Fixture loading: profession seed created.")
	}

	return nil
}
