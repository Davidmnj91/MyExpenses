package database

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/Davidmnj91/MyExpenses/config"
)

type Database struct {
	*sql.DB
}

func SetupDatabase(databaseConfig config.DatabaseConfig) (*sql.DB, error) {
	switch databaseConfig.Type() {
	case "POSTGRES":
		return NewPostgres(databaseConfig)
	default:
		return nil, errors.New(fmt.Sprintf("Cannot configure database of type %s", databaseConfig.Type))
	}
}
