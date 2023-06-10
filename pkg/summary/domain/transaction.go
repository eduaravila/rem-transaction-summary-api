package domain

import (
	"errors"
	"time"
)

type Transaction struct {
	id     string
	amount float64
	date   time.Time
}

func NewTransaction(id string, amount float64, date time.Time) (*Transaction, error) {
	if id == "" {
		return nil, errors.New("id cannot be empty")
	}

	if amount == 0 {
		return nil, errors.New("amount cannot be empty")
	}

	if date.IsZero() {
		return nil, errors.New("date cannot be empty")
	}

	return &Transaction{
		id:     id,
		amount: amount,
		date:   date,
	}, nil
}

func (t *Transaction) ID() string {
	return t.id
}

func (t *Transaction) Amount() float64 {
	return t.amount
}

func (t *Transaction) Date() time.Time {
	return t.date
}
