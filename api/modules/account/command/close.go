package command

import (
	"github.com/Davidmnj91/MyExpenses/account/domain"
)

type CloseAccountCommand struct {
	ID       string
	Password string
}

func (bus *Bus) handleCloseAccountCommand(command *CloseAccountCommand) (*domain.Account, error) {
	foundAccount, foundError := bus.repository.FindByID(command.ID)

	if foundError != nil {
		return nil, foundError
	}

	closeError := foundAccount.Close(command.Password)
	if closeError != nil {
		return nil, closeError
	}

	return &foundAccount, nil
}
