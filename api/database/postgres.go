package database

import (
	"database/sql"
	"fmt"
	"github.com/Davidmnj91/MyExpenses/config"
	"github.com/golang-migrate/migrate"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
)

// NewPostgres connect to a postgres database handle from a database configuration.
func NewPostgres(dbConfig config.DatabaseConfig) (*sql.DB, error) {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbConfig.User(), dbConfig.Password(), dbConfig.Host(), dbConfig.Port(), dbConfig.Name())
	connConfig, err := pgx.ParseConfig(connStr)

	if err != nil {
		return nil, err
	}

	db := stdlib.OpenDB(*connConfig)

	if err = runMigrations(connStr); err != nil {
		return nil, err
	}

	return db, nil
}

// runMigrations trigger migrations process for postgres database
func runMigrations(connStr string) error {
	m, err := migrate.New("file://database/migrations/", connStr)

	if err != nil {
		return err
	}

	if err = m.Up(); err != nil {
		return err
	}

	return nil
}