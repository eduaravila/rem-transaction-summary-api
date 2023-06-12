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

type UserAlreadyExistsError struct {
	Email string
}

func (e *UserNotFoundError) Error() string {
	return fmt.Sprintf("user with id %s not found", e.ID)
}

func (e *TransactionNotFoundError) Error() string {
	return fmt.Sprintf("transaction with id %s not found", e.ID)
}

func (e *UserAlreadyExistsError) Error() string {
	return fmt.Sprintf("user with mail %s already exists", e.Email)
}

type Storage interface {
	GetUserTransactions(context.Context, *User) ([]Transaction, error)
	AddUser(context.Context, *User) error
	AddTransaction(context.Context, *Transaction, *User) error
	GetUser(context.Context, string) (*User, error)
}
