package eventservice

import (
	"sync"

	"github.com/XOEF/go-event-broker/internal/domain"
)

type EventService interface {
	Subscribe(eventType domain.EventType, subscriber chan<- domain.Event)
	Publish(domain.Event)
}

type eventService struct {
	subscribers map[domain.EventType][]chan<- domain.Event
	mutex       sync.Mutex
}

// NewEventService creates a new instance of the event bus
func NewEventService() EventService {
	return &eventService{
		subscribers: make(map[domain.EventType][]chan<- domain.Event),
	}
}

// Subscribe adds a new subscriber for a given event type
func (eb *eventService) Subscribe(eventType domain.EventType, subscriber chan<- domain.Event) {
	eb.mutex.Lock()
	defer eb.mutex.Unlock()

	eb.subscribers[eventType] = append(eb.subscribers[eventType], subscriber)
}

// Publish sends an event to all subscribers of a given event type
func (eb *eventService) Publish(event domain.Event) {
	eb.mutex.Lock()
	defer eb.mutex.Unlock()

	subscribers := eb.subscribers[event.Type]
	for _, subscriber := range subscribers {
		subscriber <- event
	}
}
