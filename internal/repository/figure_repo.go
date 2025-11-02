package repository

import (
	"context"
	"errors"
	"myapp/internal/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type FigureRepository interface {
	GetAll(ctx context.Context, db *gorm.DB) ([]models.Figure, error)
	GetByID(ctx context.Context, db *gorm.DB, id string) (models.Figure, error)
	Create(ctx context.Context, db *gorm.DB, figure *models.Figure) error
}

type figureRepository struct{}

func NewFigureRepository() FigureRepository {
	return &figureRepository{}
}

func (fr *figureRepository) GetAll(ctx context.Context, db *gorm.DB) ([]models.Figure, error) {
	var figures []models.Figure

	result := db.WithContext(ctx).Preload(clause.Associations).Find(&figures)
	if result.Error != nil {
		return figures, result.Error
	}

	return figures, nil
}

func (fr *figureRepository) GetByID(ctx context.Context, db *gorm.DB, id string) (models.Figure, error) {
	var figure models.Figure

	result := db.WithContext(ctx).Preload(clause.Associations).Where("id = ?", id).First(&figure)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return figure, errors.New("no figure found with this id")
		}

		return figure, result.Error
	}

	return figure, nil
}

func (fr *figureRepository) Create(ctx context.Context, db *gorm.DB, figure *models.Figure) error {
	err := db.WithContext(ctx).Create(figure).Error
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return errors.New("figure already exists")
		}

		return err
	}

	return nil
}
