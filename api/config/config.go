package config

import (
	"fmt"
	"os"
	"strconv"
)

type Database struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
}

type Config struct {
	Host     string
	Port     int
	Database *Database
}

const (
	defaultHost       = "localhost"
	defaultPort       = 5000
	defaultDBPort     = 5432
	defaultDBUser     = "admin"
	defaultDBPassword = "admin"
	defaultDBName     = "mybookmarks"
)

func NewConfig() *Config {
	return &Config{
		Host: envString("API_HOST", defaultHost),
		Port: envInt("API_PORT", defaultPort),
		Database: &Database{
			Host:     envString("DB_HOST", defaultHost),
			Port:     envInt("DB_PORT", defaultDBPort),
			User:     envString("DB_USER", defaultDBUser),
			Password: envString("DB_PASSWORD", defaultDBPassword),
			Database: envString("DB_NAME", defaultDBName),
		},
	}
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
