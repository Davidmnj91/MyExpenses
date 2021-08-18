package query

import (
	"errors"
	"github.com/Davidmnj91/MyExpenses/account/repository"
	"github.com/Davidmnj91/MyExpenses/group/entity"
	"github.com/Davidmnj91/MyExpenses/group/model"
)

// Bus file query bus
type Bus struct {
	repository repository.Repository
}

// New create Bus instance
func New(repository repository.Repository) *Bus {
	return &Bus{repository: repository}
}

// Handle handle query
func (bus *Bus) Handle(query interface{}) (*model.Group, error) {
	switch query := query.(type) {
	case *ReadGroupByIDQuery:
		return bus.handleReadGroupByIDQuery(query)
	default:
		return nil, errors.New("query can not handled")
	}
}

func (bus *Bus) entityToModel(entity entity.Group) *model.Group {
	return &model.Group{
		ID:        entity.ID,
		Name:      entity.Name,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
		ClosedAt:  entity.ClosedAt,
	}
}
