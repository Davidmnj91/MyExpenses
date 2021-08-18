package query

import (
	"errors"
	"github.com/Davidmnj91/MyExpenses/account/repository"
	entity2 "github.com/Davidmnj91/MyExpenses/modules/group/entity"
	model2 "github.com/Davidmnj91/MyExpenses/modules/group/model"
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
func (bus *Bus) Handle(query interface{}) (*model2.Group, error) {
	switch query := query.(type) {
	case *ReadGroupByIDQuery:
		return bus.handleReadGroupByIDQuery(query)
	default:
		return nil, errors.New("query can not handled")
	}
}

func (bus *Bus) entityToModel(entity entity2.Group) *model2.Group {
	return &model2.Group{
		ID:        entity.ID,
		Name:      entity.Name,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
		ClosedAt:  entity.ClosedAt,
	}
}
