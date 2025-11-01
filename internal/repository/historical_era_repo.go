package repository

import (
	"context"
	"errors"
	"fmt"
	"myapp/config"
	"myapp/internal/models"

	"gorm.io/gorm"
)

type HistoricalEraRepository interface {
	GetByID(ctx context.Context, id uint) (models.HistoricalEra, error)
	GetByName(ctx context.Context, name string) (models.HistoricalEra, error)
}

type historicalEraRepository struct {
	DB *gorm.DB
}

func NewHistoricalEraRepository(mainDependencies *config.MainDependencies) HistoricalEraRepository {
	return &historicalEraRepository{
		DB: mainDependencies.DB,
	}
}

func (her *historicalEraRepository) GetByID(ctx context.Context, id uint) (models.HistoricalEra, error) {
	var historicalEra models.HistoricalEra

	result := her.DB.WithContext(ctx).Find(&historicalEra, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return historicalEra, fmt.Errorf("historical era with ID %d not found", id)
	}

	if result.Error != nil {
		return historicalEra, result.Error
	}

	return historicalEra, nil
}

func (her *historicalEraRepository) GetByName(ctx context.Context, name string) (models.HistoricalEra, error) {
	var historicalEra models.HistoricalEra

	result := her.DB.WithContext(ctx).Where("name = ?", name).First(&historicalEra)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return historicalEra, fmt.Errorf("historical era with name %s not found", name)
	}

	if result.Error != nil {
		return historicalEra, result.Error
	}

	return historicalEra, nil
}
