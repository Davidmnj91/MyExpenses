package command

import "github.com/Davidmnj91/MyExpenses/payment/model"

type CreatePaymentCommand struct {
	Concept              string
	Total                float32
	PaymentConfiguration model.PaymentConfiguration
}

func (bus *Bus) handleCreatePaymentCommand(command *CreatePaymentCommand) (*model.Payment, error) {
	createdPayment, createError := bus.repository.Create("", command.Concept, command.Total, command.PaymentConfiguration)

	if createError != nil {
		return nil, createError
	}

	return bus.entityToModel(createdPayment)
}
