package config

import "github.com/Davidmnj91/MyExpenses/docs"

// Swagger swagger config struct
type Swagger struct{}

// NewSwagger create swagger configuration instance
func NewSwagger() *Swagger {
	docs.SwaggerInfo.Title = "My Expenses REST api"
	docs.SwaggerInfo.Description = "This is My Expenses REST api"
	docs.SwaggerInfo.Version = "1.0.0"
	docs.SwaggerInfo.Host = "localhost:5000"
	return &Swagger{}
}
