package domain

type TransactionsSummary struct {
	user                         *User
	total                        float64 // add debit and credit
	avarageCredit                float64 // + (plus) add to credit
	avarageDebit                 float64 // - (minus) add to debit
	numberOfTransactionsPerMonth []int   // 12 months
}

const MONTHS = 12

func NewTransactionSummary(user *User, transactions []Transaction) *TransactionsSummary {
	summary := &TransactionsSummary{
		user:                         user,
		total:                        0,
		avarageCredit:                0,
		avarageDebit:                 0,
		numberOfTransactionsPerMonth: make([]int, MONTHS),
	}

	summary.CalculateAverageCredit(transactions)
	summary.CalculateAverageDebit(transactions)
	summary.CalculateTotal(transactions)
	summary.CalculateNumberofTransactionsPerMonth(transactions)

	return summary
}

func (t *TransactionsSummary) User() *User {
	return t.user
}

func (t *TransactionsSummary) Total() float64 {
	return t.total
}

func (t *TransactionsSummary) AvarageCredit() float64 {
	return t.avarageCredit
}

func (t *TransactionsSummary) AvarageDebit() float64 {
	return t.avarageDebit
}

func (t *TransactionsSummary) NumberOfTransactionsPerMonth() []int {
	return t.numberOfTransactionsPerMonth
}

func (t *TransactionsSummary) CalculateNumberofTransactionsPerMonth(transactions []Transaction) []int {
	for _, transaction := range transactions {
		t.numberOfTransactionsPerMonth[transaction.Date().Month()]++
	}

	return t.numberOfTransactionsPerMonth
}

func (t *TransactionsSummary) CalculateTotal(transactions []Transaction) float64 {
	for _, transaction := range transactions {
		t.total += transaction.Amount()
	}
	return t.total
}

func (t *TransactionsSummary) CalculateAverageCredit(transactions []Transaction) float64 {
	for _, transaction := range transactions {
		if transaction.Amount() > 0 {
			t.avarageCredit += transaction.Amount()
		}
	}

	return t.avarageCredit
}

func (t *TransactionsSummary) CalculateAverageDebit(transactions []Transaction) float64 {
	for _, transaction := range transactions {
		if transaction.Amount() < 0 {
			t.avarageDebit += transaction.Amount()
		}
	}

	return t.avarageDebit
}
