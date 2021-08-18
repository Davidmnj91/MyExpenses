package command

import "github.com/Davidmnj91/MyExpenses/group/model"

type CloseGroupCommand struct {
	ID string
}

func (bus *Bus) handleCloseGroupCommand(command *CloseGroupCommand) (*model.Group, error) {
	closedGroup, closeError := bus.repository.Close(command.ID)

	if closeError != nil {
		return nil, closeError
	}

	return bus.entityToDomain(closedGroup), nil
}
