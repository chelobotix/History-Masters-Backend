package repository

import (
	"context"
	"errors"
	"fmt"
	"myapp/internal/models"

	"gorm.io/gorm"
)

type ProfessionRepository interface {
	GetByID(ctx context.Context, db *gorm.DB, id uint) (models.Profession, error)
	GetByName(ctx context.Context, db *gorm.DB, name string) (models.Profession, error)
}

type professionRepository struct{}

func NewProfessionRepository() ProfessionRepository {
	return &professionRepository{}
}

func (pr *professionRepository) GetByID(ctx context.Context, db *gorm.DB, id uint) (models.Profession, error) {
	var profession models.Profession

	result := db.WithContext(ctx).First(&profession, id)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return profession, fmt.Errorf("profession with ID %d not found", id)
	}

	if result.Error != nil {
		return profession, result.Error
	}

	return profession, nil
}

func (pr *professionRepository) GetByName(ctx context.Context, db *gorm.DB, name string) (models.Profession, error) {
	var profession models.Profession

	result := db.WithContext(ctx).Where("name = ?", name).First(&profession)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return profession, fmt.Errorf("profession with name %s not found", name)
	}

	if result.Error != nil {
		return profession, result.Error
	}

	return profession, nil
}
