package ports

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/eduaravila/stori-challenge/pkg/summary/app"
	"github.com/eduaravila/stori-challenge/pkg/summary/app/command"
	"github.com/eduaravila/stori-challenge/pkg/summary/app/query"
	"github.com/google/uuid"
	"golang.org/x/exp/slog"
)

type HTTPServer struct {
	app app.Application
}

func NewHTTPServer(app app.Application) *HTTPServer {
	return &HTTPServer{
		app: app,
	}
}

func (h *HTTPServer) CreateUser(w http.ResponseWriter, r *http.Request) {
	userBody := &User{}

	err := json.NewDecoder(r.Body).Decode(&userBody)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)

		return
	}

	userMail := userBody.Email

	if err != nil {
		slog.Error(err.Error())
		return
	}

	id := uuid.NewString()

	if err := h.app.Commands.AddUser.Handle(r.Context(), command.AddUser{
		ID:    id,
		Name:  userBody.Name,
		Email: string(userMail),
	}); err != nil {
		slog.Error(err.Error())
		http.Error(w, "user already exists", http.StatusBadRequest)

		return
	}

	userResponse := &UserResponse{
		Id:    &id,
		Name:  userBody.Name,
		Email: userMail,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(userResponse)

}

func (h *HTTPServer) CreateTransaction(w http.ResponseWriter, r *http.Request, userID string) {
	transactionBody := &CreateTransactionJSONBody{}

	err := json.NewDecoder(r.Body).Decode(&transactionBody)

	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	date, err := time.Parse("2006-01-02", transactionBody.Date.String())

	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)

		return
	}

	h.app.Commands.RegisterTransaction.Handle(r.Context(), command.RegisterTransaction{
		UserID: userID,
		ID:     uuid.NewString(),
		Amount: float64(transactionBody.Amount),
		Date:   date,
	})
}

func (h *HTTPServer) GetTransactions(w http.ResponseWriter, r *http.Request, userID string) {
	summary, err := h.app.Queries.TransactionsSummaryQuery.Handle(r.Context(), query.TransactionsForUser{
		UserID: userID,
	})

	summaryID := uuid.NewString()

	if err != nil {
		slog.Error(err.Error())
		http.Error(w, "Bad Request", http.StatusBadRequest)

		return
	}

	summaryResponse := &TransactionSummaryResponse{
		Status: Completed,
		Data: &struct {
			NotificationId string    `json:"notificationId"`
			Recipient      string    `json:"recipient"`
			Timestamp      time.Time `json:"timestamp"`
		}{
			NotificationId: summaryID,
			Recipient:      userID,
			Timestamp:      time.Now(),
		},
		SummaryId: summaryID,
	}

	summaryR := map[string]any{
		"averageCredit":          summary.AvarageCredit(),
		"averageDebit":           summary.AvarageDebit(),
		"transactions per month": summary.NumberOfTransactionsPerMonth(),
	}

	fmt.Println(summary.AvarageCredit(), summaryResponse, summaryR, summary.AvarageDebit())

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(summaryR)

}
