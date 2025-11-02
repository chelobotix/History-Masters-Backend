package repository

import (
	"context"
	"errors"
	"fmt"
	"myapp/internal/models"

	"gorm.io/gorm"
)

type HistoricalEraRepository interface {
	GetByID(ctx context.Context, db *gorm.DB, id uint) (models.HistoricalEra, error)
	GetByName(ctx context.Context, db *gorm.DB, name string) (models.HistoricalEra, error)
}

type historicalEraRepository struct{}

func NewHistoricalEraRepository() HistoricalEraRepository {
	return &historicalEraRepository{}
}

func (her *historicalEraRepository) GetByID(ctx context.Context, db *gorm.DB, id uint) (models.HistoricalEra, error) {
	var historicalEra models.HistoricalEra

	result := db.WithContext(ctx).Find(&historicalEra, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return historicalEra, fmt.Errorf("historical era with ID %d not found", id)
	}

	if result.Error != nil {
		return historicalEra, result.Error
	}

	return historicalEra, nil
}

func (her *historicalEraRepository) GetByName(ctx context.Context, db *gorm.DB, name string) (models.HistoricalEra, error) {
	var historicalEra models.HistoricalEra

	result := db.WithContext(ctx).Where("name = ?", name).First(&historicalEra)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return historicalEra, fmt.Errorf("historical era with name %s not found", name)
	}

	if result.Error != nil {
		return historicalEra, result.Error
	}

	return historicalEra, nil
}
