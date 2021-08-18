package command

import (
	model2 "github.com/Davidmnj91/MyExpenses/modules/payment/model"
)

type UpdatePaymentCommand struct {
	ID                   string
	Concept              string
	Total                float32
	PaymentConfiguration model2.PaymentConfiguration
}

func (bus *Bus) handleUpdatePaymentCommand(command *UpdatePaymentCommand) (*model2.Payment, error) {
	updatedPayment, updateError := bus.repository.Update(command.ID, command.Concept, command.Total, command.PaymentConfiguration)

	if updateError != nil {
		return nil, updateError
	}

	return bus.entityToModel(updatedPayment)
}
