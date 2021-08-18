package query

import (
	model2 "github.com/Davidmnj91/MyExpenses/modules/group/model"
)

// ReadGroupByIDQuery read group data query by group id
type ReadGroupByIDQuery struct {
	GroupID string
}

func (bus *Bus) handleReadGroupByIDQuery(query *ReadGroupByIDQuery) (*model2.Group, error) {
	froundGroup, findError := bus.repository.FindByID(query.GroupID)

	if findError != nil {
		return nil, findError
	}

	return bus.entityToModel(froundGroup), nil
}
