package domain

import (
	"errors"
	"math"
	"strconv"
	"time"
)

type Transaction struct {
	id     string
	date   string
	amount float64
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
		date:   date.Format(time.RFC3339),    // format to RFC3339
		amount: math.Floor(amount*100) / 100, // round to 2 decimals
	}, nil
}

func (t *Transaction) ID() string {
	return t.id
}

func (t *Transaction) Amount() float64 {
	return t.amount
}

func (t *Transaction) Date() string {
	return t.date
}

func DecodeTransactionFromCSV(record []string) (*Transaction, error) {
	id := record[0]
	amountToFloat, err := strconv.ParseFloat(record[2], 64)

	if err != nil {
		return nil, err
	}

	timeToTime, err := time.Parse(time.RFC3339, record[1])

	if err != nil {
		return nil, errors.New("error parsing date")
	}

	return NewTransaction(id, amountToFloat, timeToTime)
}
