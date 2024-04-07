package tests

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/XOEF/go-event-broker/internal/domain"
	eventservice0 "github.com/XOEF/go-event-broker/internal/services"
)

func TestEventServicePublish(t *testing.T) {
	eventService := eventservice0.NewEventService()

	subscriber := make(chan domain.Event)
	defer close(subscriber)

	eventService.Subscribe(domain.EventType1, subscriber)

	event := domain.Event{Type: domain.EventType1, Data: "event 1 occured"}

	eventReceived := make(chan struct{})

	go func() {
		receivedEvent := <-subscriber
		assert.Equal(t, "event 1 occured", receivedEvent.Data, "Received event data should match")
		close(eventReceived)
	}()

	eventService.Publish(event)

	select {
	case <-eventReceived:
	case <-time.After(time.Second):
		t.Error("Timeout waiting for event to be received")
	}
}

func TestEventServiceMultipleEvents(t *testing.T) {
	eventService := eventservice0.NewEventService()

	subscriber1 := make(chan domain.Event)
	subscriber2 := make(chan domain.Event)

	eventService.Subscribe(domain.EventType1, subscriber1)
	eventService.Subscribe(domain.EventType2, subscriber2)

	event1Received := make(chan struct{})
	event2Received := make(chan struct{})

	go func() {
		receivedEvent := <-subscriber1
		assert.Equal(t, "event 1 occured", receivedEvent.Data, "Received event data should match")
		close(event1Received)
	}()

	go func() {
		receivedEvent := <-subscriber2
		assert.Equal(t, "event 2 occured", receivedEvent.Data, "Received event data should match")
		close(event2Received)
	}()

	event1 := domain.Event{Type: domain.EventType1, Data: "event 1 occured"}
	event2 := domain.Event{Type: domain.EventType2, Data: "event 2 occured"}
	eventService.Publish(event1)
	eventService.Publish(event2)

	select {
	case <-event1Received:
	case <-time.After(time.Second):
		t.Error("Timeout waiting for event 1 to be received")
	}

	select {
	case <-event2Received:
	case <-time.After(time.Second):
		t.Error("Timeout waiting for event 2 to be received")
	}
}

//	func BenchmarkEventServiceSubscribe(b *testing.B) {
//		service := eventservice0.NewEventService()
//		subscriber := make(chan domain.Event)
//
//		b.ResetTimer()
//		for i := 0; i < b.N; i++ {
//			service.Subscribe(domain.EventType1, subscriber)
//		}
//	}
func BenchmarkEventServicePublish(b *testing.B) {
	eventService := eventservice0.NewEventService()

	for i := 0; i < b.N; i++ {
		subscriber := make(chan domain.Event)
		eventService.Subscribe(domain.EventType1, subscriber)

		event := domain.Event{Type: domain.EventType1, Data: "event 1 occurred"}

		eventReceived := make(chan struct{})

		go func() {
			<-subscriber
			close(eventReceived)
		}()

		eventService.Publish(event)

		select {
		case <-eventReceived:
		case <-time.After(time.Second):
			b.Error("Timeout waiting for event to be received")
		}

		close(subscriber)
	}
}
