package events

import (
	"errors"
	"sync"
)

type EventDispatcherImpl struct {
	handlers map[string][]EventHandler
}

func NewEventDispatcher() EventDispatcher {
	return &EventDispatcherImpl{
		handlers: make(map[string][]EventHandler),
	}
}

func (dispatch *EventDispatcherImpl) AddHandler(eventName string, handler EventHandler) error {

	if _, ok := dispatch.handlers[eventName]; ok {
		for _, handlerIterated := range dispatch.handlers[eventName] {
			if handlerIterated == handler {
				return errors.New("event converter already exists")
			}
		}
	}

	dispatch.handlers[eventName] = append(dispatch.handlers[eventName], handler)

	return nil
}

func (dispatch *EventDispatcherImpl) Remove(eventName string, handler EventHandler) error {

	if handlers, ok := dispatch.handlers[eventName]; ok {
		for index, handlerIterated := range handlers {
			if handlerIterated == handler {
				handlers = append(handlers[:index], handlers[index+1:]...)
				return nil
			}
		}
	}

	return errors.New("converter not exists in event")
}

func (dispatch *EventDispatcherImpl) Has(eventName string, handler EventHandler) error {

	if _, ok := dispatch.handlers[eventName]; ok {
		for _, handlerIterated := range dispatch.handlers[eventName] {
			if handlerIterated == handler {
				return nil
			}
		}
	}

	return errors.New("event converter not exists")
}

func (dispatch *EventDispatcherImpl) Dispatch(event Event) error {

	waitGroup := &sync.WaitGroup{}

	if handlers, ok := dispatch.handlers[event.GetName()]; ok {
		for _, handlerIterated := range handlers {
			waitGroup.Add(1)
			handlerIterated.Handle(event, waitGroup)
		}
		waitGroup.Wait()
	}

	return nil
}

func (dispatch *EventDispatcherImpl) Clear() error {

	dispatch.handlers = make(map[string][]EventHandler)
	return nil
}
