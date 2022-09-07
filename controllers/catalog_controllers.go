package controllers

import (
	"e-comm/authService/repositories/postgresql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type ProductResponse struct {
	Token string
	Id    string
}

func GetProduct() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var requestBody ProductResponse
		if bodyErr := ctx.ShouldBindBodyWith(&requestBody, binding.JSON); bodyErr != nil {
			fmt.Println("body: ", bodyErr)
			ctx.Status(http.StatusInternalServerError)
			return
		}

		products, pqErr := postgresql.GetProduct(requestBody.Id)
		if pqErr != nil {
			fmt.Println("postgresql (get)", pqErr, "req id: ", requestBody.Id)
			ctx.Status(http.StatusInternalServerError)
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"products": products})
	}
}

func GetAllProducts() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		products, err := postgresql.GetAllProducts()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(products)
		ctx.JSON(http.StatusOK, gin.H{"products": products})
	}
}
