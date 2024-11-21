package event

import "time"

type TransactionCreatedEvent struct {
	Name    string
	Payload interface{}
	Date    time.Time
}

func NewTransactionCreatedEvent() *TransactionCreatedEvent {
	return &TransactionCreatedEvent{
		Name: "TransactionCreated",
		Date: time.Now(),
	}
}

func (e *TransactionCreatedEvent) GetName() string {
	return e.Name
}

func (e *TransactionCreatedEvent) GetPayload() interface{} {
	return e.GetPayload()
}

func (e *TransactionCreatedEvent) GetDateTime() time.Time {
	return e.Date
}

func (e *TransactionCreatedEvent) SetPayload(p interface{}) {
	e.Payload = p
}
