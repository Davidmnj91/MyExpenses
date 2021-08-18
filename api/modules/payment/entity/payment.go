package entity

import (
	"github.com/jackc/pgtype"
	"time"
)

type Payment struct {
	ID                   string `gorm:"primary_key"`
	Concept              string
	Amount               float32
	PaymentConfiguration pgtype.JSONB
	CreatedAt            time.Time
	UpdatedAt            time.Time
}
