package command

import (
	model2 "github.com/Davidmnj91/MyExpenses/modules/group/model"
)

type CreateGroupCommand struct {
	Name string
}

func (bus *Bus) handleCreateGroupCommand(command *CreateGroupCommand) (*model2.Group, error)  {
	createdGroup, createError := bus.repository.Create("", command.Name)

	if createError != nil {
		return nil, createError
	}

	return bus.entityToDomain(createdGroup), nil
}
