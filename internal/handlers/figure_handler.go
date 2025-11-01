package handlers

import (
	"context"
	"encoding/json"
	"myapp/config"
	"myapp/internal/models"
	"myapp/internal/repository"
	services "myapp/internal/services/figures"
	"myapp/internal/services/model_validator"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type FigureHandler interface {
	GetAll(c echo.Context) error
	GetById(c echo.Context) error
	Create(c echo.Context) error
}

type figureHandler struct {
	DB             *gorm.DB
	Logger         *zap.Logger
	Repository     repository.FigureRepository
	ModelValidator model_validator.ModelValidator
	Services       Services
}

type Services struct {
	CreateFigureService services.CreateFigureService
}

func NewFigureHandler(mainDependencies *config.MainDependencies) FigureHandler {
	return &figureHandler{
		DB:             mainDependencies.DB,
		Logger:         mainDependencies.Logger,
		Repository:     repository.NewFigureRepository(mainDependencies),
		ModelValidator: mainDependencies.ModelValidator,
		Services: Services{
			CreateFigureService: services.NewCreateFigureService(mainDependencies),
		},
	}
}

func (fh *figureHandler) GetAll(c echo.Context) error {
	figures, err := fh.Repository.GetAll()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"error":   true,
			"details": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"figures": figures,
	})
}

func (fh *figureHandler) GetById(c echo.Context) error {
	id := c.Param("id")

	figure, err := fh.Repository.GetByID(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"error":   true,
			"details": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"figure": figure,
	})
}

func (fh *figureHandler) Create(c echo.Context) error {
	var ctx context.Context
	var cancel context.CancelFunc

	if viper.GetString("server.mode") == "production" {
		queryTimeout := viper.GetInt("server.query_timeout")
		ctx, cancel = context.WithTimeout(c.Request().Context(), time.Duration(queryTimeout)*time.Second)
		defer cancel()
	} else {
		ctx = c.Request().Context()
	}

	figureInput := &models.FigureModelInput{}

	decoder := json.NewDecoder(c.Request().Body)
	if err := decoder.Decode(&figureInput); err != nil {
		return ErrorHandler(c, err, http.StatusBadRequest)
	}

	err := fh.ModelValidator.Validate(figureInput)
	if err != nil {
		return ErrorHandler(c, err, http.StatusBadRequest)
	}

	newFigure, err := fh.Services.CreateFigureService.CreateFigure(ctx, figureInput)

	if err != nil {
		return ErrorHandler(c, err, http.StatusBadRequest)
	}

	return c.JSON(http.StatusCreated, newFigure)
}
