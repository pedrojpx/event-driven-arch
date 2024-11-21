package events

import (
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type TestEvent struct {
	Name    string
	Payload interface{}
}

func (t *TestEvent) GetName() string {
	return t.Name
}
func (t *TestEvent) GetDateTime() time.Time {
	return time.Now()
}
func (t *TestEvent) GetPayload() interface{} {
	return t.Payload
}
func (t *TestEvent) SetPayload(p interface{}) {
	t.Payload = p
}

type TestEventHandler struct {
	ID int
}

func (h *TestEventHandler) Handle(event EventInterface) {

}

type EventDispatcherTestSuite struct {
	suite.Suite
	event      TestEvent
	event2     TestEvent
	handler    TestEventHandler
	handler2   TestEventHandler
	handler3   TestEventHandler
	dispatcher *EventDispatcher
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(EventDispatcherTestSuite))
}

func (s *EventDispatcherTestSuite) SetupTest() {
	s.dispatcher = NewEventDispatcher()
	s.handler = TestEventHandler{ID: 1}
	s.handler2 = TestEventHandler{ID: 2}
	s.handler3 = TestEventHandler{ID: 3}
	s.event = TestEvent{Name: "test1", Payload: "pay1"}
	s.event2 = TestEvent{Name: "test2", Payload: "pay2"}
}

func (s *EventDispatcherTestSuite) TestEventDispatcher_Register() {
	err := s.dispatcher.Register(s.event.GetName(), &s.handler)
	s.Nil(err)
	s.Equal(1, len(s.dispatcher.handlers[s.event.Name]))

	err = s.dispatcher.Register(s.event.Name, &s.handler2)
	s.Nil(err)
	s.Equal(2, len(s.dispatcher.handlers[s.event.Name]))

	s.Equal(&s.handler, s.dispatcher.handlers[s.event.GetName()][0])
	s.Equal(&s.handler2, s.dispatcher.handlers[s.event.GetName()][1])
}

func (s *EventDispatcherTestSuite) TestEventDispatcher_Register_WithSameHandler() {
	err := s.dispatcher.Register(s.event.GetName(), &s.handler)
	s.Nil(err)
	s.Equal(1, len(s.dispatcher.handlers[s.event.Name]))

	err = s.dispatcher.Register(s.event.GetName(), &s.handler)
	s.Equal(ErrHandlerAlreadyRegistered, err)
	s.Equal(1, len(s.dispatcher.handlers[s.event.Name]))

}
func (s *EventDispatcherTestSuite) TestEventDispatcher_Clear() {
	err := s.dispatcher.Register(s.event.GetName(), &s.handler)
	s.Nil(err)
	s.Equal(1, len(s.dispatcher.handlers[s.event.Name]))

	err = s.dispatcher.Register(s.event.GetName(), &s.handler2)
	s.Nil(err)
	s.Equal(2, len(s.dispatcher.handlers[s.event.Name]))

	err = s.dispatcher.Register(s.event2.GetName(), &s.handler3)
	s.Nil(err)
	s.Equal(2, len(s.dispatcher.handlers[s.event.Name]))
	s.Equal(1, len(s.dispatcher.handlers[s.event2.Name]))

	s.dispatcher.Clear()
	s.Equal(0, len(s.dispatcher.handlers))
}

func (s *EventDispatcherTestSuite) TestEventDispatcher_Has() {
	err := s.dispatcher.Register(s.event.GetName(), &s.handler)
	s.Nil(err)
	s.Equal(1, len(s.dispatcher.handlers[s.event.Name]))

	err = s.dispatcher.Register(s.event.GetName(), &s.handler2)
	s.Nil(err)
	s.Equal(2, len(s.dispatcher.handlers[s.event.Name]))

	err = s.dispatcher.Register(s.event2.GetName(), &s.handler3)
	s.Nil(err)
	s.Equal(2, len(s.dispatcher.handlers[s.event.Name]))
	s.Equal(1, len(s.dispatcher.handlers[s.event2.Name]))

	has := s.dispatcher.Has(s.event.GetName(), &s.handler)
	s.True(has)

	has = s.dispatcher.Has(s.event.GetName(), &s.handler2)
	s.True(has)

	hasNot := s.dispatcher.Has(s.event.GetName(), &s.handler3)
	s.False(hasNot)
}

type MockHandler struct {
	mock.Mock
}

func (m *MockHandler) Handle(event EventInterface) {
	m.Called(event)

}

func (s *EventDispatcherTestSuite) TestEventDispatch_Dispatch() {
	eh := &MockHandler{}
	eh.On("Handle", &s.event)

	s.dispatcher.Register(s.event.GetName(), eh)
	s.dispatcher.Dispatch(&s.event)
	eh.AssertExpectations(s.T())
	eh.AssertNumberOfCalls(s.T(), "Handle", 1)
}

func (s *EventDispatcherTestSuite) TestEventDispatcher_Unregister() {
	err := s.dispatcher.Register(s.event.GetName(), &s.handler)
	s.Nil(err)
	s.Equal(1, len(s.dispatcher.handlers[s.event.Name]))

	err = s.dispatcher.Register(s.event.GetName(), &s.handler2)
	s.Nil(err)
	s.Equal(2, len(s.dispatcher.handlers[s.event.Name]))

	err = s.dispatcher.Register(s.event2.GetName(), &s.handler3)
	s.Nil(err)
	s.Equal(2, len(s.dispatcher.handlers[s.event.Name]))
	s.Equal(1, len(s.dispatcher.handlers[s.event2.Name]))

	err = s.dispatcher.Unregister(s.event.Name, &s.handler)
	s.Nil(err)
	s.Equal(1, len(s.dispatcher.handlers[s.event.Name]))
	err = s.dispatcher.Unregister(s.event.Name, &s.handler2)
	s.Nil(err)
	s.Equal(0, len(s.dispatcher.handlers[s.event.Name]))

	err = s.dispatcher.Unregister(s.event2.Name, &s.handler3)
	s.Nil(err)
	s.Equal(0, len(s.dispatcher.handlers[s.event2.Name]))
}
