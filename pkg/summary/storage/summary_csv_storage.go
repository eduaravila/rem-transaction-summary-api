package storage

import (
	"context"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
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

var userCSVHeaders = []string{"Id", "Name", "Email"}
var transactionsCSVHeaders = []string{"Id", "Date", "Transaction"}

func (p *Path) Transactions() string {
	path := fmt.Sprintf("%s/%s", p.app, p.transactions)
	p.createIfNotExists(path)

	return fmt.Sprintf(path)
}

func (p *Path) createIfNotExists(path string) {
	err := os.MkdirAll(path, os.ModePerm)

	if err != nil {
		panic(fmt.Sprintf("error creating transactions dir: %s", err))
	}
}

func (p *Path) Users() string {
	path := fmt.Sprintf("%s/%s", p.app, p.users)
	p.createIfNotExists(path)

	return fmt.Sprintf(path)
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

	path := filepath.Join(appDir, "summary")

	return &SummaryCSVStorage{path: Path{
		app:          path,
		transactions: "transactions",
		users:        "users",
	}}, nil
}

func (s *SummaryCSVStorage) AddTransaction(ctx context.Context, t *domain.Transaction, user *domain.User) error {
	return s.WriteToFile(ctx, transactionsCSVHeaders, structToCSVArray(*t), s.getTransactionsPath(user.ID()))
}

func (s *SummaryCSVStorage) AddUser(ctx context.Context, user *domain.User) error {
	if _, userExist := s.getUserWithEmail(user.Email()); userExist == nil {
		return &domain.UserAlreadyExistsError{Email: user.Email()}
	}

	return s.WriteToFile(ctx, userCSVHeaders, structToCSVArray(*user), s.getUsersPath())
}

func (s *SummaryCSVStorage) GetUserTransactions(ctx context.Context, user *domain.User) ([]domain.Transaction, error) {
	file, err := os.Open(s.getTransactionsPath(user.ID()))

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

	for _, record := range records[1:] {
		transaction, err := domain.DecodeTransactionFromCSV(record)
		if err != nil {
			return nil, err
		}

		transactions = append(transactions, *transaction)

	}

	return transactions, nil
}

func (s *SummaryCSVStorage) getUsersPath() string {
	return fmt.Sprintf("%s/%s.csv", s.path.Users(), "index")
}

func (s *SummaryCSVStorage) getTransactionsPath(userID string) string {
	return fmt.Sprintf("%s/%s.csv", s.path.Transactions(), userID)
}

func (s *SummaryCSVStorage) WriteToFile(ctx context.Context, headers []string, data []string, path string) error {
	_, err := os.Stat(path)
	fileExists := !os.IsNotExist(err)

	fileWriter, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, os.ModePerm)

	if err != nil {
		return err
	}
	defer fileWriter.Close()

	csvWritter := csv.NewWriter(fileWriter)

	if !fileExists {
		err = csvWritter.Write(headers)
		csvWritter.Flush()

		if err != nil {
			return err
		}
	}

	fileWriter.Seek(0, io.SeekEnd)

	err = csvWritter.Write(data)

	if err != nil {
		return err
	}

	csvWritter.Flush()

	return nil
}

func (s *SummaryCSVStorage) GetUser(ctx context.Context, userID string) (*domain.User, error) {
	records, err := s.getFileRecords(s.getUsersPath())

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

func (s *SummaryCSVStorage) getUserWithEmail(email string) (*domain.User, error) {
	records, err := s.getFileRecords(s.getUsersPath())

	if err != nil {
		return nil, err
	}

	var userRecords []string

	for _, record := range records {
		recordUserID := record[2]
		if recordUserID == email {
			userRecords = record

			break
		}
	}

	if userRecords == nil {
		return nil, errors.New("user not found")
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
	value := reflect.ValueOf(t)

	if value.Kind() == reflect.Ptr {
		value = value.Elem() // Dereference the pointer
	}

	values := make([]string, value.NumField())

	for i := 0; i < value.NumField(); i++ {
		field := value.Field(i)
		values[i] = fmt.Sprintf("%v", field)
	}

	return values
}
