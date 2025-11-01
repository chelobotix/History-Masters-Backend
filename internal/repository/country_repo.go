package repository

import (
	"context"
	"errors"
	"fmt"
	"myapp/config"
	"myapp/internal/models"

	"gorm.io/gorm"
)

type CountryRepository interface {
	GetByID(ctx context.Context, id uint) (models.Country, error)
	GetByISOCode(ctx context.Context, isoCode string) (models.Country, error)
}

type countryRepository struct {
	DB *gorm.DB
}

func NewCountryRepository(mainDependencies *config.MainDependencies) CountryRepository {
	return &countryRepository{
		DB: mainDependencies.DB,
	}
}

func (cr *countryRepository) GetByID(ctx context.Context, id uint) (models.Country, error) {
	var country models.Country

	result := cr.DB.WithContext(ctx).Find(&country, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return country, fmt.Errorf("country with ID %d not found", id)
	}

	if result.Error != nil {
		return country, result.Error
	}

	return country, nil
}

func (cr *countryRepository) GetByISOCode(ctx context.Context, isoCode string) (models.Country, error) {
	var country models.Country

	result := cr.DB.WithContext(ctx).Where("iso_code = ?", isoCode).First(&country)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return country, fmt.Errorf("country with ISO code %s not found", isoCode)
	}

	if result.Error != nil {
		return country, result.Error
	}

	return country, nil
}
