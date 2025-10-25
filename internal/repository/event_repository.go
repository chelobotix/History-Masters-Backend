package repository

type EventRepository interface {
	Save()
}

type eventRepository struct{}

func NewEventRepository() EventRepository {
	return &eventRepository{}
}

func (e *eventRepository) Save() {

}
