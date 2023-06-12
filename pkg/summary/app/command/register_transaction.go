package command

import (
	"context"
	"time"

	"github.com/eduaravila/stori-challenge/pkg/summary/domain"
)

type RegisterTransaction struct {
	UserID string
	ID     string
	Amount float64
	Date   time.Time
}
type RegisterTransactionHandler struct {
	summaryStorage domain.Storage
}

func NewRegisterTransactionHandler(
	summaryStorage domain.Storage,
) RegisterTransactionHandler {
	return RegisterTransactionHandler{
		summaryStorage: summaryStorage,
	}
}

func (r *RegisterTransactionHandler) Handle(cxt context.Context, cmd RegisterTransaction) error {
	transaction, err := domain.NewTransaction(cmd.ID, cmd.Amount, cmd.Date)

	if err != nil {
		return err
	}

	user, err := r.summaryStorage.GetUser(cxt, cmd.UserID)

	if err != nil {
		return err
	}

	return r.summaryStorage.AddTransaction(cxt, transaction, user)
}
