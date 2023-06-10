package app

import (
	"github.com/eduaravila/stori-challenge/pkg/summary/app/command"
	"github.com/eduaravila/stori-challenge/pkg/summary/app/query"
)

type Appliclation struct {
	query    Queries
	commands Commands
}

type Queries struct {
	TransactionsSummaryQuery query.TransactionsSummaryQuery
}

type Commands struct {
	AddUser             command.AddUser
	RegisterTransaction command.RegisterTransaction
}
