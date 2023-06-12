package domain

import (
	"time"
)

type TransactionsSummary struct {
	user                         *User
	total                        float64 // add debit and credit
	averageCredit                float64 // + (plus) add to credit
	avarageDebit                 float64 // - (minus) add to debit
	numberOfTransactionsPerMonth []int   // 12 months
}

const MONTHS = 12

func NewTransactionSummary(user *User, transactions []Transaction) *TransactionsSummary {
	summary := &TransactionsSummary{
		user:                         user,
		total:                        0,
		averageCredit:                0,
		avarageDebit:                 0,
		numberOfTransactionsPerMonth: make([]int, MONTHS),
	}

	summary.calculateAverageCredit(transactions)
	summary.calculateAverageDebit(transactions)
	summary.calculateTotal(transactions)
	summary.calculateNumberofTransactionsPerMonth(transactions)
	return summary
}

func (t *TransactionsSummary) User() *User {
	return t.user
}

func (t *TransactionsSummary) Total() float64 {
	return t.total
}

func (t *TransactionsSummary) AvarageCredit() float64 {
	return t.averageCredit
}

func (t *TransactionsSummary) AvarageDebit() float64 {
	return t.avarageDebit
}

func (t *TransactionsSummary) NumberOfTransactionsPerMonth() []int {
	return t.numberOfTransactionsPerMonth
}

func (t *TransactionsSummary) calculateNumberofTransactionsPerMonth(transactions []Transaction) []int {
	for _, transaction := range transactions {
		month, _ := time.Parse(time.RFC3339, transaction.Date())
		t.numberOfTransactionsPerMonth[month.Month()]++
	}

	return t.numberOfTransactionsPerMonth
}

func (t *TransactionsSummary) calculateTotal(transactions []Transaction) float64 {
	for _, transaction := range transactions {
		t.total += transaction.Amount()
	}

	return t.total
}

func (t *TransactionsSummary) calculateAverageCredit(transactions []Transaction) float64 {
	var totalTransactions int

	for _, transaction := range transactions {
		if transaction.Amount() > 0 {
			totalTransactions++

			t.averageCredit += transaction.Amount()
		}
	}

	if totalTransactions < 1 {
		return 0
	}

	t.averageCredit = t.averageCredit / float64(totalTransactions)

	return t.averageCredit
}

func (t *TransactionsSummary) calculateAverageDebit(transactions []Transaction) float64 {
	var totalTransactions int

	for _, transaction := range transactions {
		if transaction.Amount() < 0 {
			totalTransactions++

			t.avarageDebit += transaction.Amount()
		}
	}

	if totalTransactions < 1 {
		return 0
	}

	t.avarageDebit = t.avarageDebit / float64(totalTransactions)

	return t.avarageDebit
}
