package payment

import (
	"database/sql"
	command2 "github.com/Davidmnj91/MyExpenses/modules/payment/command"
	controller2 "github.com/Davidmnj91/MyExpenses/modules/payment/controller"
	repository2 "github.com/Davidmnj91/MyExpenses/modules/payment/repository"
	"github.com/Davidmnj91/MyExpenses/util"
	"github.com/gin-gonic/gin"
)

func Initialize(engine *gin.Engine, db *sql.DB, util *util.Util) {
	repo := repository2.New(db)

	commandBus := command2.New(repo)

	controller2.New(engine, commandBus, util)
}
