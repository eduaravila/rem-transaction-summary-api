package command

import (
	"context"

	"github.com/eduaravila/stori-challenge/pkg/summary/domain"
)

type AddUser struct {
	ID    string
	Name  string
	Email string
}

type AddUserHandler struct {
	summaryStorage domain.Storage
}

func NewAddUserHandler(summaryStorage domain.Storage) *AddUserHandler {
	return &AddUserHandler{
		summaryStorage: summaryStorage,
	}
}

func (a *AddUserHandler) Handle(cxt context.Context, cmd AddUser) error {
	user, err := domain.NewUser(cmd.ID, cmd.Name, cmd.Email)

	if err != nil {
		return err
	}

	return a.summaryStorage.AddUser(cxt, user)
}
