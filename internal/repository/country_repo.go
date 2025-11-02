package repository

import (
	"context"
	"errors"
	"fmt"
	"myapp/internal/models"

	"gorm.io/gorm"
)

type CountryRepository interface {
	GetByID(ctx context.Context, db *gorm.DB, id uint) (models.Country, error)
	GetByName(ctx context.Context, db *gorm.DB, name string) (models.Country, error)
}

type countryRepository struct{}

func NewCountryRepository() CountryRepository {
	return &countryRepository{}
}

func (cr *countryRepository) GetByID(ctx context.Context, db *gorm.DB, id uint) (models.Country, error) {
	var country models.Country

	result := db.WithContext(ctx).Find(&country, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return country, fmt.Errorf("country with ID %d not found", id)
	}

	if result.Error != nil {
		return country, result.Error
	}

	return country, nil
}

func (cr *countryRepository) GetByName(ctx context.Context, db *gorm.DB, name string) (models.Country, error) {
	var country models.Country

	result := db.WithContext(ctx).Where("name = ?", name).First(&country)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return country, fmt.Errorf("country with name %s not found", name)
	}

	if result.Error != nil {
		return country, result.Error
	}

	return country, nil
}
