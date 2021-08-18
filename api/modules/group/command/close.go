package command

import (
	model2 "github.com/Davidmnj91/MyExpenses/modules/group/model"
)

type CloseGroupCommand struct {
	ID string
}

func (bus *Bus) handleCloseGroupCommand(command *CloseGroupCommand) (*model2.Group, error) {
	closedGroup, closeError := bus.repository.Close(command.ID)

	if closeError != nil {
		return nil, closeError
	}

	return bus.entityToDomain(closedGroup), nil
}
