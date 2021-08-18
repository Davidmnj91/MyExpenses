package controller

import (
	command2 "github.com/Davidmnj91/MyExpenses/modules/group/command"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Description update group
// @Tags Groups
// @Accept json
// @Produce json
// @Param id path string true "Group ID"
// @Param group body controller.UpdateGroupDTO true "New group name"
// @Success 201 {object} model.Group
// @Router /groups/{id} [put]
// @Security AccessToken
func (controller *Controller) update(ctx *gin.Context) {
	groupId := ctx.Param("id")

	if len(groupId) == 0 {
		httpError := controller.util.Error.HTTP.BadRequest()

		ctx.JSON(httpError.Code(), httpError.Message())
		return
	}

	var groupDto UpdateGroupDTO

	if err := ctx.ShouldBindJSON(&groupDto); err != nil {
		httpError := controller.util.Error.HTTP.BadRequest(err.Error())
		ctx.JSON(httpError.Code(), httpError.Message())
		return
	}

	cmd := &command2.UpdateGroupCommand{ID: groupId, NewName: groupDto.Name}

	updatedGroup, err := controller.commandBus.Handle(cmd)

	if err != nil {
		httpError := controller.util.Error.HTTP.InternalServerError(err.Error())

		ctx.JSON(httpError.Code(), httpError.Message())
		return
	}

	ctx.JSON(http.StatusCreated, updatedGroup)
}

type UpdateGroupDTO struct {
	Name string `json:"name" binding:"required" example:"MyChangedExpensesGroup"`
}
