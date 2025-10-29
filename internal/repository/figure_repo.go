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
	var figure []models.Figure

	err := fr.DB.Preload(clause.Associations).Find(&figure).Error
	if err != nil {
		return figure, err
	}

	return figure, err
}
