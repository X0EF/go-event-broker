package domain

import "time"

type EventType string

type Event struct {
	Type      EventType
	Timestamp time.Time
	Data      interface{}
}

const (
	EventType1 EventType = "Event1"
	EventType2 EventType = "Event2"
)
