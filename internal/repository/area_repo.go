package repository

import (
	"context"
	"errors"
	"fmt"
	"myapp/internal/models"

	"gorm.io/gorm"
)

type AreaRepository interface {
	GetByIDs(ctx context.Context, db *gorm.DB, id []uint) ([]models.Area, error)
	GetByNames(ctx context.Context, db *gorm.DB, names []string) ([]models.Area, error)
}

type areaRepository struct{}

func NewAreaRepository() AreaRepository {
	return &areaRepository{}
}

func (ar *areaRepository) GetByIDs(ctx context.Context, db *gorm.DB, ids []uint) ([]models.Area, error) {
	var areas []models.Area

	result := db.WithContext(ctx).Where("id IN (?)", ids).Find(&areas)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return areas, fmt.Errorf("areas with IDs %v not found", ids)
	}

	if result.Error != nil {
		return areas, result.Error
	}

	return areas, nil
}

func (ar *areaRepository) GetByNames(ctx context.Context, db *gorm.DB, names []string) ([]models.Area, error) {
	var areas []models.Area

	result := db.WithContext(ctx).Where("name IN (?)", names).Find(&areas)
	if result.Error != nil {
		return areas, result.Error
	}

	if len(areas) != len(names) {
		foundNames := make(map[string]bool)
		for _, a := range areas {
			foundNames[a.Name] = true
		}

		var missing []string
		for _, name := range names {
			if !foundNames[name] {
				missing = append(missing, name)
			}
		}

		return areas, fmt.Errorf("some areas not found: missing names %v", missing)
	}

	return areas, nil
}
