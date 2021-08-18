package query

import (
	model2 "github.com/Davidmnj91/MyExpenses/modules/payment/model"
)

// ReadPaymentByIDQuery read payment data query by payment id
type ReadPaymentByIDQuery struct {
	PaymentID string
}

func (bus *Bus) handleReadPaymentByIDQuery(query *ReadPaymentByIDQuery) (*model2.Payment, error) {
	froundPayment, findError := bus.repository.FindByID(query.PaymentID)

	if findError != nil {
		return nil, findError
	}

	return bus.entityToModel(froundPayment)
}
