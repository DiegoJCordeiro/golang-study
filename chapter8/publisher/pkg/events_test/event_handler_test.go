package events_test

import (
	"github.com/DiegoJCordeiro/golang-study/chapter8/publisher/pkg/events"
	"github.com/stretchr/testify/mock"
	"sync"
)

type EventsHandlerTest struct {
	mock.Mock
}

func NewEventsHandlerTest() *EventsHandlerTest {
	return &EventsHandlerTest{}
}

func (eventsDispatcherTest *EventsHandlerTest) Handle(_ events.Event, waitGroup *sync.WaitGroup) {
	eventsDispatcherTest.Called()
	waitGroup.Done()
}
