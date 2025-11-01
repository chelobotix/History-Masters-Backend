package repository

import (
	"context"
	"errors"
	"fmt"
	"myapp/config"
	"myapp/internal/models"

	"gorm.io/gorm"
)

type AchievementRepository interface {
	GetByIDs(ctx context.Context, id []uint) ([]models.Achievement, error)
	GetByNames(ctx context.Context, names []string) ([]models.Achievement, error)
	GetByName(ctx context.Context, name string) (models.Achievement, error)
	Create(ctx context.Context, achievement *models.Achievement) error
}

type achievementRepository struct {
	DB *gorm.DB
}

func NewAchievementRepository(mainDependencies *config.MainDependencies) AchievementRepository {
	return &achievementRepository{
		DB: mainDependencies.DB,
	}
}

func (ar *achievementRepository) GetByIDs(ctx context.Context, ids []uint) ([]models.Achievement, error) {
	var achievements []models.Achievement

	result := ar.DB.WithContext(ctx).Where("id IN (?)", ids).Find(&achievements)
	if result.Error != nil {
		return achievements, result.Error
	}

	if len(achievements) != len(ids) {
		foundIDs := make(map[uint]bool)
		for _, a := range achievements {
			foundIDs[a.ID] = true
		}

		var missing []uint
		for _, id := range ids {
			if !foundIDs[id] {
				missing = append(missing, id)
			}
		}

		return achievements, fmt.Errorf("some achievements not found: missing IDs %v", missing)
	}

	return achievements, nil
}

func (ar *achievementRepository) GetByNames(ctx context.Context, names []string) ([]models.Achievement, error) {
	var achievements []models.Achievement

	result := ar.DB.WithContext(ctx).Where("name IN (?)", names).Find(&achievements)
	if result.Error != nil {
		return achievements, result.Error
	}

	if len(achievements) != len(names) {
		foundNames := make(map[string]bool)
		for _, a := range achievements {
			foundNames[a.Name] = true
		}

		var missing []string
		for _, name := range names {
			if !foundNames[name] {
				missing = append(missing, name)
			}
		}

		return achievements, fmt.Errorf("some achievements not found: missing names %v", missing)
	}

	return achievements, nil
}

func (ar *achievementRepository) GetByName(ctx context.Context, name string) (models.Achievement, error) {
	var achievement models.Achievement

	result := ar.DB.WithContext(ctx).Where("name = ?", name).First(&achievement)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return achievement, errors.New("achievement not found")
	}

	if result.Error != nil {
		return achievement, result.Error
	}

	return achievement, nil
}

func (ar *achievementRepository) Create(ctx context.Context, achievement *models.Achievement) error {
	err := ar.DB.WithContext(ctx).Create(achievement).Error
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return errors.New("achievement already exists")
		}

		return err
	}

	return nil
}
