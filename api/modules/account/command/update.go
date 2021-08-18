package command

import (
	"github.com/Davidmnj91/MyExpenses/modules/account/domain"
)

type UpdateAccountCommand struct {
	ID       string
	Password string
	New string
}

func (bus *Bus) handleUpdateAccountCommand(command *UpdateAccountCommand) (*domain.Account, error) {
	account, foundError := bus.repository.FindByID(command.ID)

	if foundError != nil {
		return nil, foundError
	}

	updateError := account.UpdatePassword(command.New, command.Password)

	if updateError != nil {
		return nil, updateError
	}

	if err := bus.repository.Update(account); err != nil {
		return nil, err
	}

	return &account, nil
}
