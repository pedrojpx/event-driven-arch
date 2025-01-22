package event

import "time"

type BalanceUpdatedEvent struct {
	Name    string
	Payload interface{}
	Date    time.Time
}

func NewBalanceUpdatedEvent() *BalanceUpdatedEvent {
	return &BalanceUpdatedEvent{
		Name: "BalanceUpdated",
		Date: time.Now(),
	}
}

func (e *BalanceUpdatedEvent) GetName() string {
	return e.Name
}

func (e *BalanceUpdatedEvent) GetPayload() interface{} {
	return e.GetPayload()
}

func (e *BalanceUpdatedEvent) GetDateTime() time.Time {
	return e.Date
}

func (e *BalanceUpdatedEvent) SetPayload(p interface{}) {
	e.Payload = p
}
