package group

import (
	"github.com/Davidmnj91/MyExpenses/group/command"
	"github.com/Davidmnj91/MyExpenses/group/controller"
	"github.com/Davidmnj91/MyExpenses/group/repository"
	"github.com/Davidmnj91/MyExpenses/util"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Initialize(engine *gin.Engine, db *gorm.DB, util *util.Util) {
	repo := repository.New(db)

	commandBus := command.New(repo)

	controller.New(engine, commandBus, util)
}
