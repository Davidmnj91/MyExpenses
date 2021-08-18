package controller

import (
	command2 "github.com/Davidmnj91/MyExpenses/modules/group/command"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Description close group
// @Tags Groups
// @Accept json
// @Produce json
// @Param id path string true "Group ID"
// @Success 201 {object} model.Group
// @Router /groups/close/{id} [put]
// @Security AccessToken
func (controller *Controller) close(ctx *gin.Context) {
	groupId := ctx.Param("id")

	if len(groupId) == 0 {
		httpError := controller.util.Error.HTTP.BadRequest()

		ctx.JSON(httpError.Code(), httpError.Message())
		return
	}

	cmd := &command2.CloseGroupCommand{ID: groupId}

	closedGroup, err := controller.commandBus.Handle(cmd)

	if err != nil {
		httpError := controller.util.Error.HTTP.InternalServerError(err.Error())

		ctx.JSON(httpError.Code(), httpError.Message())
		return
	}

	ctx.JSON(http.StatusCreated, closedGroup)
}
