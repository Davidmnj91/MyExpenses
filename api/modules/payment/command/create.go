package command

import (
	model2 "github.com/Davidmnj91/MyExpenses/modules/payment/model"
)

type CreatePaymentCommand struct {
	Concept              string
	Total                float32
	PaymentConfiguration model2.PaymentConfiguration
}

func (bus *Bus) handleCreatePaymentCommand(command *CreatePaymentCommand) (*model2.Payment, error) {
	createdPayment, createError := bus.repository.Create("", command.Concept, command.Total, command.PaymentConfiguration)

	if createError != nil {
		return nil, createError
	}

	return bus.entityToModel(createdPayment)
}
