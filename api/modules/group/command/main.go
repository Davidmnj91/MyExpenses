package command

import (
	"errors"
	entity2 "github.com/Davidmnj91/MyExpenses/modules/group/entity"
	model2 "github.com/Davidmnj91/MyExpenses/modules/group/model"
	repository2 "github.com/Davidmnj91/MyExpenses/modules/group/repository"
)

type Bus struct {
	repository repository2.Repository
}

func New(groupRepository repository2.Repository) *Bus {
	return &Bus{repository: groupRepository}
}

func (bus *Bus) Handle(command interface{}) (*model2.Group, error) {
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

func (bus *Bus) entityToDomain(entity entity2.Group) *model2.Group {
	return &model2.Group{
		ID:        entity.ID,
		Name:      entity.Name,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
		ClosedAt:  entity.ClosedAt,
	}
}
