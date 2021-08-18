package controller

import (
	command2 "github.com/Davidmnj91/MyExpenses/modules/payment/command"
	"github.com/Davidmnj91/MyExpenses/util"
	"github.com/gin-gonic/gin"
)

// Controller payment controller
type Controller struct {
	route      *gin.Engine
	commandBus *command2.Bus
	util       *util.Util
}

// New create payment controller instance
func New(route *gin.Engine, bus *command2.Bus, util *util.Util) *Controller {
	controller := &Controller{
		route:      route,
		commandBus: bus,
		util: util,
	}
	controller.SetupRoutes()
	return controller
}

// SetupRoutes setup payment router
func (controller *Controller) SetupRoutes() {
	controller.route.POST("payments", controller.create)
	controller.route.PUT("payments/:id", controller.update)
}
