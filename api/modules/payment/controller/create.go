package controller

import (
	command2 "github.com/Davidmnj91/MyExpenses/modules/payment/command"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Description create payment
// @Tags Payments
// @Accept json
// @Produce json
// @Param CreatePayment body controller.CreatePaymentDTO true "Create Payment data"
// @Success 201 {object} model.Payment
// @Router /payments [post]
// @Security AccessToken
func (controller *Controller) create(ctx *gin.Context)  {
	var paymentDto CreatePaymentDTO

	if err := ctx.ShouldBindJSON(&paymentDto); err != nil {
		httpError := controller.util.Error.HTTP.BadRequest(err.Error())
		ctx.JSON(httpError.Code(), httpError.Message())
		return
	}

	cmd := &command2.CreatePaymentCommand{Concept: paymentDto.Name}

	createdPayment, err := controller.commandBus.Handle(cmd)

	if err != nil {
		httpError := controller.util.Error.HTTP.InternalServerError(err.Error())

		ctx.JSON(httpError.Code(), httpError.Message())
		return
	}

	ctx.JSON(http.StatusCreated, createdPayment)
}

type CreatePaymentDTO struct {
	Name string `json:"name" binding:"required" example:"MyExpensesPayment"`
}