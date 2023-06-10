// Package ports provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.12.4 DO NOT EDIT.
package ports

import (
	"time"

	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
)

// Defines values for TransactionSummaryResponseStatus.
const (
	Completed TransactionSummaryResponseStatus = "completed"
	Pending   TransactionSummaryResponseStatus = "pending"
)

// Error defines model for Error.
type Error struct {
	Error string `json:"error"`
}

// TransactionSummaryResponse defines model for TransactionSummaryResponse.
type TransactionSummaryResponse struct {
	Data *struct {
		NotificationId *string    `json:"notificationId,omitempty"`
		Recipient      *string    `json:"recipient,omitempty"`
		Timestamp      *time.Time `json:"timestamp,omitempty"`
	} `json:"data,omitempty"`
	Status    *TransactionSummaryResponseStatus `json:"status,omitempty"`
	SummaryId *string                           `json:"summaryId,omitempty"`
}

// TransactionSummaryResponseStatus defines model for TransactionSummaryResponse.Status.
type TransactionSummaryResponseStatus string

// User defines model for User.
type User struct {
	Email openapi_types.Email `json:"email"`
	Name  string              `json:"name"`
}

// UserResponse defines model for UserResponse.
type UserResponse struct {
	Email openapi_types.Email `json:"email"`
	Id    *string             `json:"id,omitempty"`
	Name  string              `json:"name"`
}

// CreateTransactionJSONBody defines parameters for CreateTransaction.
type CreateTransactionJSONBody struct {
	Amount float32   `json:"amount"`
	Date   time.Time `json:"date"`
}

// CreateUserJSONRequestBody defines body for CreateUser for application/json ContentType.
type CreateUserJSONRequestBody = User

// CreateTransactionJSONRequestBody defines body for CreateTransaction for application/json ContentType.
type CreateTransactionJSONRequestBody CreateTransactionJSONBody