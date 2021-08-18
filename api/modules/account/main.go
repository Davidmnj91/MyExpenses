package account

import (
	"database/sql"
	command2 "github.com/Davidmnj91/MyExpenses/modules/account/command"
	controller2 "github.com/Davidmnj91/MyExpenses/modules/account/controller"
	infrastructure2 "github.com/Davidmnj91/MyExpenses/modules/account/infrastructure"
	"github.com/Davidmnj91/MyExpenses/util"
	"github.com/gin-gonic/gin"
)

func Initialize(engine *gin.Engine, db *sql.DB, util *util.Util) {
	repo := infrastructure2.NewRepository(db)

	commandBus := command2.New(repo)

	controller2.New(engine, commandBus, util)
}
