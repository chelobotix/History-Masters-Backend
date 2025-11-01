package repository

import (
	"context"
	"errors"
	"myapp/config"
	"myapp/internal/models"

	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type FigureRepository interface {
	GetAll() ([]models.Figure, error)
	GetByID(id string) (models.Figure, error)
	Create(ctx context.Context, figure *models.Figure) error
}

type figureRepository struct {
	DB     *gorm.DB
	Logger *zap.Logger
}

func NewFigureRepository(mainDependencies *config.MainDependencies) FigureRepository {
	return &figureRepository{
		DB:     mainDependencies.DB,
		Logger: mainDependencies.Logger,
	}
}

func (fr *figureRepository) GetAll() ([]models.Figure, error) {
	var figures []models.Figure

	result := fr.DB.Preload(clause.Associations).Find(&figures)
	if result.Error != nil {
		return figures, result.Error
	}

	return figures, nil
}

func (fr *figureRepository) GetByID(id string) (models.Figure, error) {
	var figure models.Figure

	result := fr.DB.Preload(clause.Associations).Where("id = ?", id).First(&figure)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return figure, errors.New("no figure found with this id")
		}

		return figure, result.Error
	}

	return figure, nil
}

func (fr *figureRepository) Create(ctx context.Context, figure *models.Figure) error {
	err := fr.DB.WithContext(ctx).Create(figure).Error
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return errors.New("figure already exists")
		}

		return err
	}

	return nil
}
