package query

import (
	"errors"
	"github.com/Davidmnj91/MyExpenses/payment/entity"
	"github.com/Davidmnj91/MyExpenses/payment/model"
	"github.com/Davidmnj91/MyExpenses/payment/repository"
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
func (bus *Bus) Handle(query interface{}) (*model.Payment, error) {
	switch query := query.(type) {
	case *ReadPaymentByIDQuery:
		return bus.handleReadPaymentByIDQuery(query)
	default:
		return nil, errors.New("query can not handled")
	}
}

func (bus *Bus) entityToModel(entity entity.Payment) (*model.Payment, error) {
	paymentConfiguration := model.PaymentConfiguration{}
	if err := entity.PaymentConfiguration.AssignTo(&paymentConfiguration); err != nil {
		return nil, err
	}
	return &model.Payment{
		ID:                   entity.ID,
		Concept:              entity.Concept,
		Amount:               entity.Amount,
		PaymentConfiguration: paymentConfiguration,
		CreatedAt:            entity.CreatedAt,
		UpdatedAt:            entity.UpdatedAt,
	}, nil
}
