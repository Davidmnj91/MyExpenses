package controller

import (
	command2 "github.com/Davidmnj91/MyExpenses/modules/account/command"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Description create account
// @Tags Accounts
// @Accept json
// @Produce json
// @Param CreateAccount body controller.CreateAccountDTO true "Create Account data"
// @Success 201 {object} domain.AccountAnemic
// @Router /accounts [post]
// @Security AccessToken
func (controller *Controller) create(ctx *gin.Context) {
	var accountDto CreateAccountDTO

	if err := ctx.ShouldBindJSON(&accountDto); err != nil {
		httpError := controller.util.Error.HTTP.BadRequest(err.Error())
		ctx.JSON(httpError.Code(), httpError.Message())
		return
	}

	cmd := command2.CreateAccountCommand{Name: accountDto.Name, Password: accountDto.Password}

	createdAccount, err := controller.commandBus.Handle(cmd)

	if err != nil {
		httpError := controller.util.Error.HTTP.InternalServerError(err.Error())

		ctx.JSON(httpError.Code(), httpError.Message())
		return
	}

	ctx.JSON(http.StatusCreated, createdAccount)
}

type CreateAccountDTO struct {
	Name     string `json:"name" binding:"required" example:"MyExpensesAccount"`
	Password string `json:"password" binding:"required" example:"mypassword1234"`
}
