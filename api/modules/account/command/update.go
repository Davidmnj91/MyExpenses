package command

import (
	"github.com/Davidmnj91/MyExpenses/account/domain"
)

type UpdateAccountCommand struct {
	ID       string
	Password string
	New string
}

func (bus *Bus) handleUpdateAccountCommand(command *UpdateAccountCommand) (*domain.Account, error) {
	foundAccount, foundError := bus.repository.FindByID(command.ID)

	if foundError != nil {
		return nil, foundError
	}

	updateError := foundAccount.UpdatePassword(command.New, command.Password)

	if updateError != nil {
		return nil, updateError
	}

	return &foundAccount, nil
}
