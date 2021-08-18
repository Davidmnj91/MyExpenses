package config

import "strconv"

// DatabaseConfig database config interface
type DatabaseConfig interface {
	Type() string
	Host() string
	Port() string
	Name() string
	User() string
	Password() string
}

// Database database config struct
type Database struct {
	dbType   string
	port     string
	host     string
	name     string
	user     string
	password string
}

// NewDatabaseConfig create database instance
func NewDatabaseConfig() *Database {
	dbType := envString("DB_TYPE", defaultDBType)
	host := envString("DB_HOST", defaultHost)
	port := envInt("DB_PORT", defaultDBPort)
	name := envString("DB_NAME", defaultDBName)
	user := envString("DB_USER", defaultDBUser)
	password := envString("DB_PASSWORD", defaultDBPassword)

	return &Database{
		dbType:   dbType,
		port:     strconv.Itoa(port),
		host:     host,
		name:     name,
		user:     user,
		password: password,
	}
}

// Type get database type
func (database *Database) Type() string {
	return database.dbType
}

// Host get database host
func (database *Database) Host() string {
	return database.host
}

// Port get database port number
func (database *Database) Port() string {
	return database.port
}

// Name get database name
func (database *Database) Name() string {
	return database.name
}

// User get database user name
func (database *Database) User() string {
	return database.user
}

// Password get database user password
func (database *Database) Password() string {
	return database.password
}
