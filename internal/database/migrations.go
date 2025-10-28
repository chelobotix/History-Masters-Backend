package database

import (
	"myapp/internal/models"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&models.Figure{},
		&models.Year{},
		&models.Country{},
		&models.HistoricalEra{},
		&models.Area{},
		&models.FigureAreas{},
		&models.Profession{},
		&models.Achievement{},
		&models.FigureAchievements{},
	)
	if err != nil {
		return err
	}

	return nil
}
