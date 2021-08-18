package database

import (
	"errors"
	"fmt"
	"github.com/Davidmnj91/MyExpenses/config"
	"gorm.io/gorm"
)

type Database interface {
	Connect(configuration *config.Database) (*gorm.DB, error)
}

func SetupDatabase(config *config.Database) (*gorm.DB, error) {
	switch config.Type() {
	case "POSTGRES":
		return Postgres{}.Connect(config)
	default:
		return nil, errors.New(fmt.Sprintf("Cannot configure database of type %s", config.Type))
	}
}
