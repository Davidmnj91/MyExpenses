package controller

import (
	"github.com/Davidmnj91/MyExpenses/account/command"
	"github.com/Davidmnj91/MyExpenses/util"
	"github.com/gin-gonic/gin"
)

// Controller group controller
type Controller struct {
	route      *gin.Engine
	commandBus *command.Bus
	util       *util.Util
}

// New create new http router
func New(route *gin.Engine, bus *command.Bus, util *util.Util) *Controller {
	controller := &Controller{
		route:      route,
		commandBus: bus,
		util:       util,
	}

	controller.SetupRoutes()
	return controller
}

func (controller *Controller) SetupRoutes() {
	controller.route.POST("/accounts", controller.create)
	controller.route.PUT("/accounts/:accountID/password", controller.updateAccountPassword)
	controller.route.DELETE("/accounts/:accountID", controller.close)
}
