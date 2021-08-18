package controller

import (
	"github.com/Davidmnj91/MyExpenses/account/command"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Description update account
// @Tags Accounts
// @Accept json
// @Produce json
// @Param id path string true "Account ID"
// @Param group body controller.UpdateAccountPasswordDTO true "Update Password"
// @Success 201 {object} domain.AccountAnemic
// @Router /accounts/{id} [put]
// @Security AccessToken
func (controller *Controller) updateAccountPassword(ctx *gin.Context) {
	accountId := ctx.Param("id")

	if len(accountId) == 0 {
		httpError := controller.util.Error.HTTP.BadRequest()

		ctx.JSON(httpError.Code(), httpError.Message())
		return
	}

	var accountPasswordDTO UpdateAccountPasswordDTO

	if err := ctx.ShouldBindJSON(&accountPasswordDTO); err != nil {
		httpError := controller.util.Error.HTTP.BadRequest(err.Error())
		ctx.JSON(httpError.Code(), httpError.Message())
		return
	}

	cmd := &command.UpdateAccountCommand{
		ID:       accountId,
		Password: accountPasswordDTO.Password,
		New:      accountPasswordDTO.New,
	}

	updatedAccountPassword, err := controller.commandBus.Handle(cmd)

	if err != nil {
		httpError := controller.util.Error.HTTP.InternalServerError(err.Error())

		ctx.JSON(httpError.Code(), httpError.Message())
		return
	}

	ctx.JSON(http.StatusCreated, updatedAccountPassword)
}

// UpdateAccountPasswordDTO dto for update account password
type UpdateAccountPasswordDTO struct {
	Password string `json:"password" binding:"required" example:"oldPassword"`
	New      string `json:"new" binding:"required" example:"newPassword"`
}
