package repository

import (
	"context"
	"errors"
	"fmt"
	"myapp/config"
	"myapp/internal/models"

	"gorm.io/gorm"
)

type ProfessionRepository interface {
	GetByID(ctx context.Context, id uint) (models.Profession, error)
	GetByName(ctx context.Context, name string) (models.Profession, error)
}

type professionRepository struct {
	DB *gorm.DB
}

func NewProfessionRepository(mainDependencies *config.MainDependencies) ProfessionRepository {
	return &professionRepository{
		DB: mainDependencies.DB,
	}
}

func (pr *professionRepository) GetByID(ctx context.Context, id uint) (models.Profession, error) {
	var profession models.Profession

	result := pr.DB.WithContext(ctx).First(&profession, id)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return profession, fmt.Errorf("profession with ID %d not found", id)
	}

	if result.Error != nil {
		return profession, result.Error
	}

	return profession, nil
}

func (pr *professionRepository) GetByName(ctx context.Context, name string) (models.Profession, error) {
	var profession models.Profession

	result := pr.DB.WithContext(ctx).Where("name = ?", name).First(&profession)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return profession, fmt.Errorf("profession with name %s not found", name)
	}

	if result.Error != nil {
		return profession, result.Error
	}

	return profession, nil
}
