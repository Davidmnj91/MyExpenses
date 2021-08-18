package command

import "github.com/Davidmnj91/MyExpenses/payment/model"

type UpdatePaymentCommand struct {
	ID                   string
	Concept              string
	Total                float32
	PaymentConfiguration model.PaymentConfiguration
}

func (bus *Bus) handleUpdatePaymentCommand(command *UpdatePaymentCommand) (*model.Payment, error) {
	updatedPayment, updateError := bus.repository.Update(command.ID, command.Concept, command.Total, command.PaymentConfiguration)

	if updateError != nil {
		return nil, updateError
	}

	return bus.entityToModel(updatedPayment)
}
