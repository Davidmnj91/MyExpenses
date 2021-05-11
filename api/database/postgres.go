package database

import (
	"fmt"
	"github.com/Davidmnj91/MyBookmarks/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Connect to a database handle from a connection string.
func Connect(configuration *config.Database) (*gorm.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=disable", configuration.Host, configuration.Port, configuration.Database, configuration.User, configuration.Password)
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return db, nil
}
