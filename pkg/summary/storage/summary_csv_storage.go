package storage

import (
	"context"
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"reflect"

	"github.com/eduaravila/stori-challenge/pkg/summary/domain"
)

type csvTransaction struct {
	ID     string  `csv:"id"`
	Amount float64 `csv:"amount"`
	Date   string  `csv:"date"`
}

type Path struct {
	app          string
	transactions string
	users        string
}

func (p *Path) Transactions() string {
	p.createIfNotExists(p.transactions)

	return fmt.Sprintf("%s/%s", p.app, p.transactions)
}

func (p *Path) createIfNotExists(path string) {
	err := os.MkdirAll(path, os.ModePerm)

	if err != nil {
		panic(fmt.Sprintf("error creating transactions dir: %s", err))
	}
}

func (p *Path) Users() string {
	p.createIfNotExists(p.users)

	return fmt.Sprintf("%s/%s", p.app, p.users)
}

type SummaryCSVStorage struct {
	path Path
}

func NewSummaryCSVStorageWithDefaultPath() (*SummaryCSVStorage, error) {
	executablePath, err := os.Executable()
	if err != nil {
		panic(fmt.Sprintf("error getting executable path: %s", err))
	}

	appDir := filepath.Dir(executablePath)

	path := filepath.Join(appDir, "transactions")

	return &SummaryCSVStorage{path: Path{
		app:          path,
		transactions: "transactions",
		users:        "users",
	}}, nil
}

func (s *SummaryCSVStorage) AddTransaction(ctx context.Context, t *domain.Transaction, user *domain.User) error {
	return s.WriteToFile(ctx, structToCSVArray(t), s.getTransationsPath(user.ID()))
}

func (s *SummaryCSVStorage) AddUser(ctx context.Context, user *domain.User) error {
	return s.WriteToFile(ctx, structToCSVArray(user), s.getUsersPath(user.ID()))
}

func (s *SummaryCSVStorage) GetUserTransactions(ctx context.Context, user *domain.User) ([]domain.Transaction, error) {
	file, err := os.Open(s.getTransationsPath(user.ID()))

	if err != nil {
		return nil, err
	}

	defer file.Close()

	csvReader := csv.NewReader(file)

	records, err := csvReader.ReadAll()

	if err != nil {
		return nil, err
	}

	transactions := []domain.Transaction{}

	for _, record := range records {
		transaction, err := domain.DecodeTransactionFromCSV(record)
		if err != nil {
			return nil, err
		}

		transactions = append(transactions, *transaction)

	}

	return transactions, nil
}

func (s *SummaryCSVStorage) getUsersPath(userID string) string {
	return fmt.Sprintf("%s/%s.csv", s.path.Users(), userID)
}

func (s *SummaryCSVStorage) getTransationsPath(userID string) string {
	return fmt.Sprintf("%s/%s.csv", s.path.Transactions(), userID)
}

func (s *SummaryCSVStorage) WriteToFile(ctx context.Context, data []string, path string) error {
	fileWriter, err := os.Open(path)

	if err != nil {
		return err
	}

	defer fileWriter.Close()
	csvWritter := csv.NewWriter(fileWriter)

	err = csvWritter.Write(data)

	if err != nil {
		return err
	}

	csvWritter.Flush()

	return nil
}

func (s *SummaryCSVStorage) GetUser(ctx context.Context, userID string) (*domain.User, error) {
	records, err := s.getFileRecords(s.getTransationsPath(userID))

	if err != nil {
		return nil, err
	}

	var userRecords []string

	for _, record := range records {
		recordUserID := record[0]
		if recordUserID == userID {
			userRecords = record

			break
		}
	}

	if userRecords == nil {
		return nil, &domain.UserNotFoundError{ID: userID}
	}

	return domain.DecodeUserFromCSV(userRecords)
}

func (s *SummaryCSVStorage) getFileRecords(path string) ([][]string, error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	csvReader := csv.NewReader(file)

	records, err := csvReader.ReadAll()

	if err != nil {
		return nil, err
	}

	return records, nil
}

func structToCSVArray(t any) []string {
	keys := reflect.ValueOf(t)
	values := make([]string, keys.NumField())

	for i := 0; i < keys.NumField(); i++ {
		field := keys.Field(i)
		values[i] = field.String()
	}

	return values
}
