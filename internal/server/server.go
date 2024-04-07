package server

import (
	"fmt"
	"net/http"
	"time"

	eventservice0 "github.com/XOEF/go-event-broker/internal/event-service-0"
)

type Server struct {
	port          int
	EventService0 eventservice0.EventService
}

func NewServer() *http.Server {
	NewServer := &Server{
		port:          5000,
		EventService0: eventservice0.NewEventService(),
	}

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
