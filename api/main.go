package main

import (
	"github.com/Davidmnj91/MyExpenses/config"
	"github.com/Davidmnj91/MyExpenses/database"
	"github.com/Davidmnj91/MyExpenses/modules/account"
	"github.com/Davidmnj91/MyExpenses/modules/group"
	"github.com/Davidmnj91/MyExpenses/modules/payment"
	"github.com/Davidmnj91/MyExpenses/util"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
)

// @title MyExpenses API
// @version 1.0
// @description This is My expenses api server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:5000
// @BasePath /api/v1

// @securityDefinitions.apikey AccessToken
// @in header
// @name Authorization
func main() {
	configuration := config.Initialize()
	util := util.Initialize()

	// establish DB connection
	db, err := database.SetupDatabase(configuration.Database())
	if err != nil {
		panic(err)
	}

	gin.SetMode(configuration.Server().Mode())
	route := gin.Default()

	account.Initialize(route, db, util)
	group.Initialize(route, db, util)
	payment.Initialize(route, db, util)

	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	log.Fatal(route.Run(":" + configuration.Server().Port()))
}
