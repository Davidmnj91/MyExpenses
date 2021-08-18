package command

import (
	"errors"
	"github.com/Davidmnj91/MyExpenses/group/entity"
	"github.com/Davidmnj91/MyExpenses/group/model"
	"github.com/Davidmnj91/MyExpenses/group/repository"
)

type Bus struct {
	repository repository.Repository
}

func New(groupRepository repository.Repository) *Bus {
	return &Bus{repository: groupRepository}
}

func (bus *Bus) Handle(command interface{}) (*model.Group, error) {
	switch command := command.(type) {
	case *CreateGroupCommand:
		return bus.handleCreateGroupCommand(command)
	case *UpdateGroupCommand:
		return bus.handleUpdateGroupCommand(command)
	case *CloseGroupCommand:
		return bus.handleCloseGroupCommand(command)
	default:
		return nil, errors.New("invalid command type")
	}
}

func (bus *Bus) entityToDomain(entity entity.Group) *model.Group {
	return &model.Group{
		ID:        entity.ID,
		Name:      entity.Name,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
		ClosedAt:  entity.ClosedAt,
	}
}
