package events_test

import (
	"time"
)

type EventTest struct {
	Name    string
	Payload interface{}
}

func (event *EventTest) GetName() string {
	return event.Name
}

func (event *EventTest) GetPayload() interface{} {
	return event.Payload
}

func (event *EventTest) GetTime() time.Time {
	return time.Now()
}

func NewEventTest(name string, payload interface{}) *EventTest {
	return &EventTest{
		Name:    name,
		Payload: payload,
	}
}
