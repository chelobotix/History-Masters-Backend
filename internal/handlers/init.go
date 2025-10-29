package handlers

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type HandlersInit struct {
	Health HealthHandler
	Figure FigureHandler
}

func NewHandlersInit(db *gorm.DB, logger *zap.Logger) *HandlersInit {
	return &HandlersInit{
		Health: NewHealthHandler(logger),
		Figure: NewFigureHandler(db, logger),
	}
}
