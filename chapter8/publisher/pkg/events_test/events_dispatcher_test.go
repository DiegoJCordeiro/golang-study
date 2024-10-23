package events_test

import (
	"github.com/DiegoJCordeiro/golang-study/chapter8/publisher/pkg/events"
	"github.com/stretchr/testify/suite"
	"testing"
)

type TestEventDispatcherSuite struct {
	suite.Suite
	firstEvent      events.Event
	secondEvent     events.Event
	firstHandler    events.EventHandler
	secondHandler   events.EventHandler
	eventDispatcher events.EventDispatcher
}

func (suite *TestEventDispatcherSuite) SetupSuite() {

	eventDispatcher := events.NewEventDispatcher()

	suite.firstEvent = NewEventTest("Event 1", "")
	suite.secondEvent = NewEventTest("Event 2", "")
	suite.firstHandler = NewEventsHandlerTest()
	suite.secondHandler = NewEventsHandlerTest()
	suite.eventDispatcher = eventDispatcher
}

func (suite *TestEventDispatcherSuite) TestEventDispatcher_Has() {

	errAddHandlerFirst := suite.eventDispatcher.AddHandler(suite.firstEvent.GetName(), suite.firstHandler)
	errHasHandlerFirst := suite.eventDispatcher.Has(suite.firstEvent.GetName(), suite.firstHandler)

	suite.Nil(errAddHandlerFirst)
	suite.Nil(errHasHandlerFirst)
}

func (suite *TestEventDispatcherSuite) TestEventDispatcher_AddHandler() {

	errAddHandlerFirst := suite.eventDispatcher.AddHandler(suite.firstEvent.GetName(), suite.firstHandler)
	errHasHandlerFirst := suite.eventDispatcher.Has(suite.firstEvent.GetName(), suite.firstHandler)

	errAddHandlerSecond := suite.eventDispatcher.AddHandler(suite.secondEvent.GetName(), suite.secondHandler)
	errHasHandlerSecond := suite.eventDispatcher.Has(suite.secondEvent.GetName(), suite.secondHandler)

	suite.Nil(errAddHandlerFirst)
	suite.Nil(errHasHandlerFirst)

	suite.Nil(errAddHandlerSecond)
	suite.Nil(errHasHandlerSecond)
}

func (suite *TestEventDispatcherSuite) TestEventDispatcher_Remove() {

	errRemoveHandlerFirst := suite.eventDispatcher.Remove(suite.firstEvent.GetName(), suite.firstHandler)

	suite.Nil(errRemoveHandlerFirst)
}

func (suite *TestEventDispatcherSuite) TestEventDispatcher_Clear() {

	errClearHandlers := suite.eventDispatcher.Clear()

	suite.Nil(errClearHandlers)
}

func (suite *TestEventDispatcherSuite) TestEventDispatcher_Dispatch() {

	eventHandler := NewEventsHandlerTest()

	eventHandler.On("Handle")

	errAddHandler := suite.eventDispatcher.AddHandler(suite.firstEvent.GetName(), eventHandler)
	errDispatchHandler := suite.eventDispatcher.Dispatch(suite.firstEvent)

	eventHandler.AssertExpectations(suite.T())
	eventHandler.AssertNumberOfCalls(suite.T(), "Handle", 1)

	suite.Nil(errAddHandler)
	suite.Nil(errDispatchHandler)
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(TestEventDispatcherSuite))
}
