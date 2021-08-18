package command

import (
	"errors"
	"github.com/Davidmnj91/MyExpenses/payment/entity"
	"github.com/Davidmnj91/MyExpenses/payment/model"
	"github.com/Davidmnj91/MyExpenses/payment/repository"
)

type Bus struct {
	repository repository.Repository
}

func New(paymentRepository repository.Repository) *Bus {
	return &Bus{repository: paymentRepository}
}

func (bus *Bus) Handle(command interface{}) (*model.Payment, error) {
	switch command := command.(type) {
	case *CreatePaymentCommand:
		return bus.handleCreatePaymentCommand(command)
	case *UpdatePaymentCommand:
		return bus.handleUpdatePaymentCommand(command)
	default:
		return nil, errors.New("invalid command type")
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
