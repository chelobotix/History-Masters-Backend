package handlers

import "gorm.io/gorm"

type HandlersInit struct {
	Health HealthHandler
	Event  EventHandler
}

func NewHandlersInit(db *gorm.DB) *HandlersInit {
	return &HandlersInit{
		Health: NewHealthHandler(),
		Event:  NewEventHandler(),
	}
}
