package handlers

import "myapp/config"

type HandlersInit struct {
	Health HealthHandler
	Figure FigureHandler
}

func NewHandlersInit(mainDependencies *config.MainDependencies) *HandlersInit {
	return &HandlersInit{
		Health: NewHealthHandler(mainDependencies),
		Figure: NewFigureHandler(mainDependencies),
	}
}
