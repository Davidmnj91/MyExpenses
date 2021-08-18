package command

import "github.com/Davidmnj91/MyExpenses/group/model"

type UpdateGroupCommand struct {
	ID      string
	NewName string
}

func (bus *Bus) handleUpdateGroupCommand(command *UpdateGroupCommand) (*model.Group, error) {
	updatedGroup, updateError := bus.repository.Update(command.ID, command.NewName)

	if updateError != nil {
		return nil, updateError
	}

	return bus.entityToDomain(updatedGroup), nil
}
