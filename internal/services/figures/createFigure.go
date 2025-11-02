package services

import (
	"context"
	"myapp/config"
	"myapp/internal/models"
	"myapp/internal/repository"

	"gorm.io/gorm"
)

type CreateFigureService interface {
	CreateFigure(ctx context.Context, fmi *models.FigureInput) (*models.Figure, error)
	handleAchievements(ctx context.Context, db *gorm.DB, achievements []string) ([]models.Achievement, error)
}

type createFigureService struct {
	DB                      *gorm.DB
	FigureRepository        repository.FigureRepository
	YearRepository          repository.YearRepository
	CountryRepository       repository.CountryRepository
	HistoricalEraRepository repository.HistoricalEraRepository
	AreaRepository          repository.AreaRepository
	ProfessionRepository    repository.ProfessionRepository
	AchievementRepository   repository.AchievementRepository
}

func NewCreateFigureService(mainDependencies *config.MainDependencies) CreateFigureService {
	return &createFigureService{
		DB:                      mainDependencies.DB,
		YearRepository:          repository.NewYearRepository(),
		CountryRepository:       repository.NewCountryRepository(),
		HistoricalEraRepository: repository.NewHistoricalEraRepository(),
		AreaRepository:          repository.NewAreaRepository(),
		ProfessionRepository:    repository.NewProfessionRepository(),
		AchievementRepository:   repository.NewAchievementRepository(),
		FigureRepository:        repository.NewFigureRepository(),
	}
}

func (s *createFigureService) CreateFigure(ctx context.Context, fmi *models.FigureInput) (*models.Figure, error) {
	var newFigure *models.Figure

	err := s.DB.Transaction(func(tx *gorm.DB) error {
		yearOfBirth, err := s.YearRepository.GetByYear(ctx, tx, fmi.YearOfBirth)
		if err != nil {
			return err
		}

		yearOfDeath, err := s.YearRepository.GetByYear(ctx, tx, fmi.YearOfDeath)
		if err != nil {
			return err
		}

		country, err := s.CountryRepository.GetByName(ctx, tx, fmi.CountryISOCode)
		if err != nil {
			return err
		}

		historicalEra, err := s.HistoricalEraRepository.GetByName(ctx, tx, fmi.HistoricalEra)
		if err != nil {
			return err
		}

		areas, err := s.AreaRepository.GetByNames(ctx, tx, fmi.Areas)
		if err != nil {
			return err
		}

		profession, err := s.ProfessionRepository.GetByName(ctx, tx, fmi.Profession)
		if err != nil {
			return err
		}

		achievements, err := s.handleAchievements(ctx, tx, fmi.Achievements)
		if err != nil {
			return err
		}

		newFigure = &models.Figure{
			FirstName:     fmi.FirstName,
			LastName:      fmi.LastName,
			YearOfBirth:   yearOfBirth,
			YearOfDeath:   yearOfDeath,
			Country:       country,
			HistoricalEra: historicalEra,
			Areas:         areas,
			Profession:    profession,
			Achievements:  achievements,
		}

		err = s.FigureRepository.Create(ctx, tx, newFigure)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return newFigure, nil
}

func (s *createFigureService) handleAchievements(ctx context.Context, db *gorm.DB, achievementNames []string) ([]models.Achievement, error) {
	var foundAchievements []models.Achievement

	for _, achievementName := range achievementNames {
		achievement, err := s.AchievementRepository.GetByName(ctx, db, achievementName)

		if err != nil {
			if err.Error() == "achievement not found" {
				newAchievement := &models.Achievement{
					Name: achievementName,
				}

				err = s.AchievementRepository.Create(ctx, db, newAchievement)
				if err != nil {
					return foundAchievements, err
				}

				foundAchievements = append(foundAchievements, *newAchievement)
				continue
			} else {
				return foundAchievements, err
			}
		}

		foundAchievements = append(foundAchievements, achievement)
	}

	return foundAchievements, nil
}
