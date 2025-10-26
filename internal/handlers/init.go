package handlers

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type HandlersInit struct {
	Health HealthHandler
	Event  EventHandler
}

func NewHandlersInit(db *gorm.DB, logger *zap.Logger) *HandlersInit {
	return &HandlersInit{
		Health: NewHealthHandler(logger),
		Event:  NewEventHandler(),
	}
}
