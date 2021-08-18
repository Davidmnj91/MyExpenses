package controller

import (
	"github.com/Davidmnj91/MyExpenses/group/command"
	"github.com/Davidmnj91/MyExpenses/util"
	"github.com/gin-gonic/gin"
)

// Controller group controller
type Controller struct {
	route      *gin.Engine
	commandBus *command.Bus
	util       *util.Util
}

// New create group controller instance
func New(route *gin.Engine, bus *command.Bus, util *util.Util) *Controller {
	controller := &Controller{
		route:      route,
		commandBus: bus,
		util: util,
	}
	controller.SetupRoutes()
	return controller
}

// SetupRoutes setup group router
func (controller *Controller) SetupRoutes() {
	controller.route.POST("groups", controller.create)
	controller.route.PUT("groups/:id", controller.update)
	controller.route.PUT("groups/close/:id", controller.close)
}
