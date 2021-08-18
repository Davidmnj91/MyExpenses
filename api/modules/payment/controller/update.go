package controller

import (
	"github.com/Davidmnj91/MyExpenses/payment/command"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Description update payment
// @Tags Payments
// @Accept json
// @Produce json
// @Param id path string true "Payment ID"
// @Param payment body controller.UpdatePaymentDTO true "New payment name"
// @Success 201 {object} model.Payment
// @Router /payments/{id} [put]
// @Security AccessToken
func (controller *Controller) update(ctx *gin.Context) {
	paymentId := ctx.Param("id")

	if len(paymentId) == 0 {
		httpError := controller.util.Error.HTTP.BadRequest()

		ctx.JSON(httpError.Code(), httpError.Message())
		return
	}

	var paymentDto UpdatePaymentDTO

	if err := ctx.ShouldBindJSON(&paymentDto); err != nil {
		httpError := controller.util.Error.HTTP.BadRequest(err.Error())
		ctx.JSON(httpError.Code(), httpError.Message())
		return
	}

	cmd := &command.UpdatePaymentCommand{ID: paymentId, NewName: paymentDto.Name}

	updatedPayment, err := controller.commandBus.Handle(cmd)

	if err != nil {
		httpError := controller.util.Error.HTTP.InternalServerError(err.Error())

		ctx.JSON(httpError.Code(), httpError.Message())
		return
	}

	ctx.JSON(http.StatusCreated, updatedPayment)
}

type UpdatePaymentDTO struct {
	Name string `json:"name" binding:"required" example:"MyChangedExpensesPayment"`
}
