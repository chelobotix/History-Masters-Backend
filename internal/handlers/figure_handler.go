package handlers

import (
	"myapp/internal/models"
	"myapp/internal/repository"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type FigureHandler interface {
	GetAll(c echo.Context) error
	GetById(c echo.Context) error
}

type figureHandler struct {
	DB         *gorm.DB
	Logger     *zap.Logger
	Repository repository.FigureRepository
}

func NewFigureHandler(db *gorm.DB, logger *zap.Logger) FigureHandler {
	return &figureHandler{
		DB:         db,
		Logger:     logger,
		Repository: repository.NewFigureRepository(db, logger),
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
	var figure models.Figure

	if err := c.Bind(&figure); err != nil {
		return ErrorHandler(err, http.StatusBadRequest, c)
	}

	return nil
}
