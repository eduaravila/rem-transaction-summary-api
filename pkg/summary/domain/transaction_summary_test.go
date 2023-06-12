package domain_test

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/eduaravila/stori-challenge/pkg/summary/domain"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

type MockTransaction struct {
	Total                        float64 // add debit and credit
	AverageCredit                float64 // + (plus) add to credit
	AvarageDebit                 float64 // - (minus) add to debit
	NumberOfTransactionsPerMonth []int   // 12 months
}

func TestTransactionsSummary(t *testing.T) {
	// Create a mock user and transaction summary
	user, _ := domain.NewUser(uuid.NewString(), "John Doe", "john.doe@example.com")

	transaction1, _ := domain.NewTransaction(uuid.NewString(), 60.5, time.Date(2020, time.July, 1, 0, 0, 0, 0, time.UTC))
	transaction2, _ := domain.NewTransaction(uuid.NewString(), -10.3, time.Date(2020, time.July, 1, 0, 0, 0, 0, time.UTC))
	transaction3, _ := domain.NewTransaction(uuid.NewString(), -20.46, time.Date(2020, time.August, 1, 0, 0, 0, 0, time.UTC))
	transaction4, _ := domain.NewTransaction(uuid.NewString(), 10, time.Date(2020, time.August, 1, 0, 0, 0, 0, time.UTC))

	transactionSummary := domain.NewTransactionSummary(user, []domain.Transaction{
		*transaction1,
		*transaction2,
		*transaction3,
		*transaction4,
	})

	testCases := []struct {
		name     string
		input    *domain.TransactionsSummary
		expected MockTransaction
	}{
		{
			name:  "Test 1",
			input: transactionSummary,
			expected: MockTransaction{
				Total:                        39.74,
				AverageCredit:                35.25,
				AvarageDebit:                 -15.38,
				NumberOfTransactionsPerMonth: []int{0, 0, 0, 0, 0, 0, 0, 2, 2, 0, 0, 0},
			},
		},
	}

	for _, tc := range testCases {
		require.Equal(t, tc.expected.Total, tc.input.Total())
		require.Equal(t, tc.expected.AverageCredit, tc.input.AvarageCredit())
		require.Equal(t, tc.expected.AvarageDebit, tc.input.AvarageDebit())

		fmt.Println(tc.expected.NumberOfTransactionsPerMonth, tc.input.NumberOfTransactionsPerMonth())
		require.True(t, reflect.DeepEqual(tc.expected.NumberOfTransactionsPerMonth, tc.input.NumberOfTransactionsPerMonth()))
	}

}
