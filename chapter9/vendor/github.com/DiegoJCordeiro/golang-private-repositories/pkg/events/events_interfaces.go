package events

import (
	"sync"
	"time"
)

type Event interface {
	GetName() string
	GetPayload() interface{}
	GetTime() time.Time
}

type EventHandler interface {
	Handle(event Event, waitGroup *sync.WaitGroup)
}

type EventDispatcher interface {
	AddHandler(eventName string, handler EventHandler) error
	Remove(eventName string, handler EventHandler) error
	Has(eventName string, handler EventHandler) error
	Dispatch(event Event) error
	Clear() error
}
