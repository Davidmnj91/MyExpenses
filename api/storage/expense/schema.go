package expense

import (
	"github.com/gofrs/uuid"
	"time"
)

type Schema struct {
	ID        			 uuid.UUID `gorm:"primary_key;type:uuid;default:gen_random_uuid()"`
	Concept              string
	TotalPrice           float32
	Date                 time.Time
	PaymentConfiguration PaymentConfiguration
	GroupId              string
}
