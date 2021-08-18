package query

import (
	"errors"
	entity2 "github.com/Davidmnj91/MyExpenses/modules/payment/entity"
	model2 "github.com/Davidmnj91/MyExpenses/modules/payment/model"
	repository2 "github.com/Davidmnj91/MyExpenses/modules/payment/repository"
)

// Bus file query bus
type Bus struct {
	repository repository2.Repository
}

// New create Bus instance
func New(repository repository2.Repository) *Bus {
	return &Bus{repository: repository}
}

// Handle handle query
func (bus *Bus) Handle(query interface{}) (*model2.Payment, error) {
	switch query := query.(type) {
	case *ReadPaymentByIDQuery:
		return bus.handleReadPaymentByIDQuery(query)
	default:
		return nil, errors.New("query can not handled")
	}
}

func (bus *Bus) entityToModel(entity entity2.Payment) (*model2.Payment, error) {
	paymentConfiguration := model2.PaymentConfiguration{}
	if err := entity.PaymentConfiguration.AssignTo(&paymentConfiguration); err != nil {
		return nil, err
	}
	return &model2.Payment{
		ID:                   entity.ID,
		Concept:              entity.Concept,
		Amount:               entity.Amount,
		PaymentConfiguration: paymentConfiguration,
		CreatedAt:            entity.CreatedAt,
		UpdatedAt:            entity.UpdatedAt,
	}, nil
}
