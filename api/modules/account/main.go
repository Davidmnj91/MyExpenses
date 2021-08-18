package account

import (
	"github.com/Davidmnj91/MyExpenses/account/command"
	"github.com/Davidmnj91/MyExpenses/account/controller"
	"github.com/Davidmnj91/MyExpenses/account/infrastructure"
	"github.com/Davidmnj91/MyExpenses/util"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Initialize(engine *gin.Engine, db *gorm.DB, util *util.Util) {
	repo := infrastructure.NewRepository(db)

	commandBus := command.New(repo)

	controller.New(engine, commandBus, util)
}
