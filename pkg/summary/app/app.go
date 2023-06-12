package app

import (
	"github.com/eduaravila/stori-challenge/pkg/summary/app/command"
	"github.com/eduaravila/stori-challenge/pkg/summary/app/query"
)

type Application struct {
	Queries  Queries
	Commands Commands
}

type Queries struct {
	TransactionsSummaryQuery query.TransactionsSummaryHandler
}

type Commands struct {
	AddUser             command.AddUserHandler
	RegisterTransaction command.RegisterTransactionHandler
}
