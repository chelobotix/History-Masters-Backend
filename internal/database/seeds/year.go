package seeds

import (
	"log"
	"myapp/internal/models"
	"time"

	"gorm.io/gorm"
)

func YearSeeds(db *gorm.DB) error {
	var count int64

	db.Model(&models.Year{}).Count(&count)
	if count > 0 {
		log.Println("Skipping fixture loading: year already exists.")
		return nil
	}

	var years []models.Year

	for i := 2700; i >= 1; i-- {
		year := models.Year{
			Year: i,
			BC:   true,
		}

		years = append(years, year)
	}

	for i := 0; i <= time.Now().Year(); i++ {
		year := models.Year{
			Year: i,
			BC:   false,
		}

		years = append(years, year)
	}

	if err := db.Create(&years).Error; err != nil {
		return err
	} else {
		log.Println("Fixture loading: year seed created.")
	}

	return nil
}
