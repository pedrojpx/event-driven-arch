package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	ID        string
	From      *Account
	To        *Account
	Amount    float64
	CreatedAt time.Time
}

func NewTransaction(from, to *Account, amount float64) (*Transaction, error) {
	t := &Transaction{
		ID:        uuid.New().String(),
		From:      from,
		To:        to,
		Amount:    amount,
		CreatedAt: time.Now(),
	}
	if err := t.Validate(); err != nil {
		return nil, err
	}
	t.Commit()
	return t, nil
}

func (t *Transaction) Commit() {
	t.From.Debit(t.Amount)
	t.To.Credit(t.Amount)
}

func (t *Transaction) Validate() error {
	if t.Amount <= 0 {
		return errors.New("amount must be greater than zero")
	}
	if t.From.Balance < t.Amount {
		return errors.New("insufficient funds")
	}
	return nil
}
