package repository

import (
	"context"
	"errors"
	"fmt"
	"myapp/internal/models"

	"gorm.io/gorm"
)

type YearRepository interface {
	GetByID(ctx context.Context, db *gorm.DB, id uint) (models.Year, error)
	GetByYear(ctx context.Context, db *gorm.DB, yearTarget uint) (models.Year, error)
}

type yearRepository struct{}

func NewYearRepository() YearRepository {
	return &yearRepository{}
}

func (yr *yearRepository) GetByID(ctx context.Context, db *gorm.DB, id uint) (models.Year, error) {
	var year models.Year

	result := db.WithContext(ctx).Find(&year, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return year, fmt.Errorf("year with ID %d not found", id)
	}

	if result.Error != nil {
		return year, result.Error
	}

	return year, nil
}

func (yr *yearRepository) GetByYear(ctx context.Context, db *gorm.DB, yearTarget uint) (models.Year, error) {
	var year models.Year

	result := db.WithContext(ctx).Where("year = ?", yearTarget).First(&year)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return year, fmt.Errorf("year %d not found", yearTarget)
	}

	if result.Error != nil {
		return year, result.Error
	}

	return year, nil
}
