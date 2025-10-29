package repository

import (
	"myapp/internal/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type FigureRepository interface {
	GetAll() ([]models.Figure, error)
}

type figureRepository struct {
	DB *gorm.DB
}

func NewFigureRepository(db *gorm.DB) FigureRepository {
	return &figureRepository{
		DB: db,
	}
}

func (fr *figureRepository) GetAll() ([]models.Figure, error) {
	var figures []models.Figure

	err := fr.DB.Preload(clause.Associations).Find(&figures).Error
	if err != nil {
		return figures, err
	}

	return figures, nil
}
