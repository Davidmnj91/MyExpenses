package controller

import (
	command2 "github.com/Davidmnj91/MyExpenses/modules/account/command"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Description close account
// @Tags Accounts
// @Accept json
// @Produce json
// @Param id path string true "Account ID"
// @Success 201 {object} domain.AccountAnemic
// @Router /accounts/close/{id} [put]
// @Security AccessToken
func (controller *Controller) close(ctx *gin.Context) {
	accountId := ctx.Param("id")

	if len(accountId) == 0 {
		httpError := controller.util.Error.HTTP.BadRequest()

		ctx.JSON(httpError.Code(), httpError.Message())
		return
	}

	var closeAccountDTO CloseAccountDTO

	if err := ctx.ShouldBindJSON(&closeAccountDTO); err != nil {
		httpError := controller.util.Error.HTTP.BadRequest(err.Error())
		ctx.JSON(httpError.Code(), httpError.Message())
		return
	}

	cmd := &command2.CloseAccountCommand{ID: accountId, Password: closeAccountDTO.Password}

	closedAccount, err := controller.commandBus.Handle(cmd)

	if err != nil {
		httpError := controller.util.Error.HTTP.InternalServerError(err.Error())

		ctx.JSON(httpError.Code(), httpError.Message())
		return
	}

	ctx.JSON(http.StatusCreated, closedAccount)
}

type CloseAccountDTO struct {
	Password string `json:"password" binding:"required" example:"userPassword1234"`
}
