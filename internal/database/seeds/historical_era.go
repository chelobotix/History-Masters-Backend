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
		{Name: "Classical Era", NameEs: "Era Clásica"},
		{Name: "Early Modernd Era", NameEs: "Era Temprana Moderna"},
		{Name: "Prehistory", NameEs: "Prehistoria"},
		{Name: "The Middle Age", NameEs: "Edad Media"},
		{Name: "Modern Era", NameEs: "Era Moderna"},
	}

	if err := db.Create(&historicalEra).Error; err != nil {
		return err
	} else {
		log.Println("Fixture loading: historical_era seed created.")
	}

	return nil
}
