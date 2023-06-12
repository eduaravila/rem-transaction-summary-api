package domain

import (
	"context"
	"fmt"
)

type TransactionNotFoundError struct {
	ID string
}
type UserNotFoundError struct {
	ID string
}

func (e *UserNotFoundError) Error() string {
	return fmt.Sprintf("user with id %s not found", e.ID)
}

func (e *TransactionNotFoundError) Error() string {
	return fmt.Sprintf("transaction with id %s not found", e.ID)
}

type Storage interface {
	GetUserTransactions(context.Context, *User) ([]Transaction, error)
	AddUser(context.Context, *User) error
	AddTransaction(context.Context, *Transaction, *User) error
	GetUser(context.Context, string) (*User, error)
}
