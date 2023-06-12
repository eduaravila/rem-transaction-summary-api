package domain

import (
	"errors"
	"strconv"
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

func DecodeTransactionFromCSV(record []string) (*Transaction, error) {
	id := record[0]
	amountToFloat, err := strconv.ParseFloat(record[1], 64)

	if err != nil {
		return nil, err
	}

	timeToTime, err := time.Parse(time.RFC3339, record[2])

	if err != nil {
		return nil, errors.New("error parsing date")
	}

	return NewTransaction(id, amountToFloat, timeToTime)
}
