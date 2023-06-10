package domain

import "context"

type Storage interface {
	GetUserTransactions(context.Context, string) ([]Transaction, error)
	AddUser(context.Context, *User) error
	AddTransaction(context.Context, *Transaction) error
}
