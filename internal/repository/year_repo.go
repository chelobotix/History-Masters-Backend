package repository

import (
	"context"
	"errors"
	"fmt"
	"myapp/config"
	"myapp/internal/models"

	"gorm.io/gorm"
)

type YearRepository interface {
	GetByID(ctx context.Context, id uint) (models.Year, error)
	GetByYear(ctx context.Context, yearTarget uint) (models.Year, error)
}

type yearRepository struct {
	DB *gorm.DB
}

func NewYearRepository(mainDependencies *config.MainDependencies) YearRepository {
	return &yearRepository{
		DB: mainDependencies.DB,
	}
}

func (yr *yearRepository) GetByID(ctx context.Context, id uint) (models.Year, error) {
	var year models.Year

	result := yr.DB.WithContext(ctx).Find(&year, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return year, fmt.Errorf("year with ID %d not found", id)
	}

	if result.Error != nil {
		return year, result.Error
	}

	return year, nil
}

func (yr *yearRepository) GetByYear(ctx context.Context, yearTarget uint) (models.Year, error) {
	var year models.Year

	result := yr.DB.WithContext(ctx).Where("year = ?", yearTarget).First(&year)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return year, fmt.Errorf("year %d not found", yearTarget)
	}

	if result.Error != nil {
		return year, result.Error
	}

	return year, nil
}
