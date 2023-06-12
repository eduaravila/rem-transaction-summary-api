package query

import (
	"context"

	"github.com/eduaravila/stori-challenge/internal/errors"
	"github.com/eduaravila/stori-challenge/pkg/summary/domain"
)

type TransactionsForUser struct {
	UserID string
}

type TransactionsSummaryHandler struct {
	summaryStorage      domain.Storage
	notificationService NotificationsService
}

type NotificationsService interface {
	Send(user *domain.User, transactionSummary *domain.TransactionsSummary) error
}

func NewTransactionsSummaryHandler(
	summaryStorage domain.Storage,
	notificationService NotificationsService,
) TransactionsSummaryHandler {
	return TransactionsSummaryHandler{
		summaryStorage:      summaryStorage,
		notificationService: notificationService,
	}
}

func (t *TransactionsSummaryHandler) Handle(
	ctx context.Context,
	query TransactionsForUser,
) (*domain.TransactionsSummary, error) {
	user, err := t.summaryStorage.GetUser(ctx, query.UserID)

	if err != nil {
		return nil, errors.NewSlugError(err.Error(), "user-not-found")
	}

	transactions, err := t.summaryStorage.GetUserTransactions(ctx, user)

	if err != nil {
		return nil, errors.NewSlugError(err.Error(), "transactions-not-found")
	}

	summary := domain.NewTransactionSummary(user, transactions)

	go func() {
		t.notificationService.Send(user, summary)
	}()

	return summary, nil
}
