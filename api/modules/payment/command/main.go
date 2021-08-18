package command

import (
	"errors"
	entity2 "github.com/Davidmnj91/MyExpenses/modules/payment/entity"
	model2 "github.com/Davidmnj91/MyExpenses/modules/payment/model"
	repository2 "github.com/Davidmnj91/MyExpenses/modules/payment/repository"
)

type Bus struct {
	repository repository2.Repository
}

func New(paymentRepository repository2.Repository) *Bus {
	return &Bus{repository: paymentRepository}
}

func (bus *Bus) Handle(command interface{}) (*model2.Payment, error) {
	switch command := command.(type) {
	case *CreatePaymentCommand:
		return bus.handleCreatePaymentCommand(command)
	case *UpdatePaymentCommand:
		return bus.handleUpdatePaymentCommand(command)
	default:
		return nil, errors.New("invalid command type")
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
