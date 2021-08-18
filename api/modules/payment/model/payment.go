package model

import "time"

type Payment struct {
	ID                   string               `json:"id" example:"c6a52781-c927-42ec-bc9e-254ddde5f904"`
	Concept              string               `json:"concept" example:"Shopping"`
	Amount               float32              `json:"amount" example:"134.45"`
	PaymentConfiguration PaymentConfiguration `json:"payment_configuration" example:"{\"1\": 32.3, \"4\": 45.5}"`
	CreatedAt            time.Time            `json:"createdAt" example:"2019-12-23 12:27:37"`
	UpdatedAt            time.Time            `json:"updatedAt" example:"2019-12-23 12:27:37"`
}

type PaymentConfiguration struct {
	AccountPaymentBalance map[string]float32
}
