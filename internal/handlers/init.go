package handlers

type HandlersInit struct {
	Health HealthHandler
	Event  EventHandler
}

func NewHandlersInit() *HandlersInit {
	return &HandlersInit{
		Health: NewHealthHandler(),
		Event:  NewEventHandler(),
	}
}
