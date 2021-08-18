package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	swagger  *Swagger
	server   ServerConfig
	database DatabaseConfig
}

type ConfigInterface interface {
	Swagger() *Swagger
	Server() ServerConfig
	Database() DatabaseConfig
}

const (
	defaultHost       = "localhost"
	defaultPort       = 5000
	defaultDBPort     = 5432
	defaultDBUser     = "admin"
	defaultDBPassword = "admin"
	defaultDBName     = "myexpenses"
)

func Initialize() ConfigInterface {
	return &Config{
		swagger:  NewSwagger(),
		server:   NewServerConfig(),
		database: NewDatabaseConfig(),
	}
}

func (config *Config) Swagger() *Swagger {
	return config.swagger
}

func (config *Config) Server() ServerConfig {
	return config.server
}

func (config *Config) Database() DatabaseConfig {
	return config.database
}

func envString(env, fallback string) string {
	e := os.Getenv(env)
	if e == "" {
		return fallback
	}
	return e
}

func envInt(env string, fallback int) int {
	e := os.Getenv(env)
	if e == "" {
		return fallback
	}
	i, err := strconv.Atoi(e)

	if err != nil {
		msg := fmt.Sprintf("Cannot read %s environment variable, %d", env, err)
		panic(msg)
	}
	return i
}
