package query

import "github.com/Davidmnj91/MyExpenses/group/model"

// ReadGroupByIDQuery read group data query by group id
type ReadGroupByIDQuery struct {
	GroupID string
}

func (bus *Bus) handleReadGroupByIDQuery(query *ReadGroupByIDQuery) (*model.Group, error) {
	froundGroup, findError := bus.repository.FindByID(query.GroupID)

	if findError != nil {
		return nil, findError
	}

	return bus.entityToModel(froundGroup), nil
}
