package command

import "github.com/Davidmnj91/MyExpenses/group/model"

type CreateGroupCommand struct {
	Name string
}

func (bus *Bus) handleCreateGroupCommand(command *CreateGroupCommand) (*model.Group, error)  {
	createdGroup, createError := bus.repository.Create("", command.Name)

	if createError != nil {
		return nil, createError
	}

	return bus.entityToDomain(createdGroup), nil
}
