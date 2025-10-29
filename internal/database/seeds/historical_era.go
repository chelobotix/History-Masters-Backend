package seeds

import (
	"log"
	"myapp/internal/models"

	"gorm.io/gorm"
)

func HistoricalEraSeeds(db *gorm.DB) error {
	var count int64

	db.Model(&models.HistoricalEra{}).Count(&count)
	if count > 0 {
		log.Println("Skipping fixture loading: historical_era already exists.")
		return nil
	}

	var historicalEra = []models.HistoricalEra{
		{Name: "Classical Era"},
		{Name: "Early Modernd Era"},
		{Name: "Prehistory"},
		{Name: "The Middle Age"},
		{Name: "Modern Era"},
	}

	if err := db.Create(&historicalEra).Error; err != nil {
		return err
	} else {
		log.Println("Fixture loading: historical_era seed created.")
	}

	return nil
}
