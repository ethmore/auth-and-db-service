package controllers

import (
	"auth-and-db-service/repositories/postgresql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type PaymentBody struct {
	Token      string
	BuyerID    string
	AddressID  string
	TotalPrice string
}

type UpdatePayment struct {
	Token     string
	PaymentID int
	Status    string
}

func CreatePayment(paymentRepo postgresql.IPaymentRepo) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var body PaymentBody
		if err := ctx.ShouldBindBodyWith(&body, binding.JSON); err != nil {
			fmt.Println("body err: ", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{})
			return
		}

		paymentID, err := paymentRepo.InsertPayment(body.BuyerID, body.AddressID, body.TotalPrice)
		if err != nil {
			fmt.Println("posgresql (insert): ", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"paymentID": paymentID})
	}
}

func UpdatePaymentStatus(paymentRepo postgresql.IPaymentRepo) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var payment UpdatePayment
		if err := ctx.ShouldBindBodyWith(&payment, binding.JSON); err != nil {
			fmt.Println("body err: ", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{})
			return
		}

		updateErr := paymentRepo.UpdatePaymentStatus(payment.Status, payment.PaymentID)
		if updateErr != nil {
			fmt.Println("posgresql (update payment status): ", updateErr)
			ctx.JSON(http.StatusInternalServerError, gin.H{})
			return
		}

		ctx.Status(http.StatusOK)
	}
}
