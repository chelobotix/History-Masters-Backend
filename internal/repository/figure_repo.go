package repository

import (
	"errors"
	"myapp/internal/models"

	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type FigureRepository interface {
	GetAll() ([]models.Figure, error)
	GetByID(id string) (models.Figure, error)
}

type figureRepository struct {
	DB     *gorm.DB
	Logger *zap.Logger
}

func NewFigureRepository(db *gorm.DB, logger *zap.Logger) FigureRepository {
	return &figureRepository{
		DB:     db,
		Logger: logger,
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

func (fr *figureRepository) Create(figure models.Figure) (models.Figure, error) {
	result := fr.DB.Create(&figure)
	if result.Error != nil {
		return figure, result.Error
	}

	return figure, nil
}
