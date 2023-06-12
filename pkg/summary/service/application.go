package service

import (
	"github.com/eduaravila/stori-challenge/pkg/summary/adapters"
	"github.com/eduaravila/stori-challenge/pkg/summary/app"
	"github.com/eduaravila/stori-challenge/pkg/summary/app/command"
	"github.com/eduaravila/stori-challenge/pkg/summary/app/query"
	"github.com/eduaravila/stori-challenge/pkg/summary/domain"
	"github.com/eduaravila/stori-challenge/pkg/summary/storage"
)

func NewApplication(storageType string) app.Application {
	emailNotification := adapters.NewEmailNotification()
	return newApplication(storageType, emailNotification)

}

func newApplication(storageType string, notificationService query.NotificationsService) app.Application {
	var storageSummary domain.Storage

	switch storageType {
	case "csv":
		csvStorage, err := storage.NewSummaryCSVStorageWithDefaultPath()
		if err != nil {
			panic(err)
		}

		storageSummary = csvStorage
		// case "postgres":
		// storage = adapters.NewPostgresStorage()
	}

	return app.Application{
		Queries: app.Queries{
			TransactionsSummaryQuery: query.NewTransactionsSummaryHandler(storageSummary, notificationService),
		},
		Commands: app.Commands{
			AddUser:             command.NewAddUserHandler(storageSummary),
			RegisterTransaction: command.NewRegisterTransactionHandler(storageSummary),
		},
	}

}
