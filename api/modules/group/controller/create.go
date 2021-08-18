package controller

import (
	command2 "github.com/Davidmnj91/MyExpenses/modules/group/command"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Description create group
// @Tags Groups
// @Accept json
// @Produce json
// @Param CreateGroup body controller.CreateGroupDTO true "Create Group data"
// @Success 201 {object} model.Group
// @Router /groups [post]
// @Security AccessToken
func (controller *Controller) create(ctx *gin.Context)  {
	var groupDto CreateGroupDTO

	if err := ctx.ShouldBindJSON(&groupDto); err != nil {
		httpError := controller.util.Error.HTTP.BadRequest(err.Error())
		ctx.JSON(httpError.Code(), httpError.Message())
		return
	}

	cmd := &command2.CreateGroupCommand{Name: groupDto.Name}

	createdGroup, err := controller.commandBus.Handle(cmd)

	if err != nil {
		httpError := controller.util.Error.HTTP.InternalServerError(err.Error())

		ctx.JSON(httpError.Code(), httpError.Message())
		return
	}

	ctx.JSON(http.StatusCreated, createdGroup)
}

type CreateGroupDTO struct {
	Name string `json:"name" binding:"required" example:"MyExpensesGroup"`
}