package storage

type csvTransaction struct {
	ID     string  `csv:"id"`
	Amount float64 `csv:"amount"`
	Date   string  `csv:"date"`
}

type SummaryCSVStorage struct {
	path string
}
