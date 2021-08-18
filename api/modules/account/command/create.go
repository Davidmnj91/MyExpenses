package command

import (
	"github.com/Davidmnj91/MyExpenses/modules/account/domain"
	"github.com/google/uuid"
)

type CreateAccountCommand struct {
	Name     string
	Password string
}

func (bus *Bus) handleCreateAccountCommand(command *CreateAccountCommand) (*domain.Account, error) {
	accountOptions := domain.AccountOptions{
		ID:       uuid.NewString(),
		Name:     command.Name,
		Password: command.Password,
	}

	account := domain.NewAccount(accountOptions)

	if err := bus.repository.Save(account); err != nil {
		return nil, err
	}

	return &account, nil
}
