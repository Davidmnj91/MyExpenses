package command

import (
	"errors"
	"github.com/Davidmnj91/MyExpenses/modules/account/domain"
)

type Bus struct {
	repository domain.AccountRepository
}

func New(accountRepository domain.AccountRepository) *Bus {
	return &Bus{repository: accountRepository}
}

func (bus *Bus) Handle(command interface{}) (*domain.Account, error) {
	switch command := command.(type) {
	case *CreateAccountCommand:
		return bus.handleCreateAccountCommand(command)
	case *UpdateAccountCommand:
		return bus.handleUpdateAccountCommand(command)
	case *CloseAccountCommand:
		return bus.handleCloseAccountCommand(command)
	default:
		return nil, errors.New("invalid command type")
	}
}
