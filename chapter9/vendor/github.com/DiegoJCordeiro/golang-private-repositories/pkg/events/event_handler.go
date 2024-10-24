package events

import "sync"

type EventHandlerImpl struct{}

func NewEventHandler() EventHandler {
	return &EventHandlerImpl{}
}

func (eventHandler *EventHandlerImpl) Handle(event Event, waitGroup *sync.WaitGroup) {
	waitGroup.Done()
}
