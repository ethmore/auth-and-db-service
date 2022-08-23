package controllers

import (
	"e-comm/authService/repositories/postgresql"
	"fmt"
	"log"
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
		if err := ctx.ShouldBindBodyWith(&requestBody, binding.JSON); err != nil {
			log.Printf("%+v", err)
		}

		fmt.Println(requestBody.Id)
		products, err := postgresql.GetProduct(requestBody.Id)
		if err != nil {
			log.Fatal(err)
		}
		ctx.JSON(http.StatusOK, gin.H{"products": products})
	}
}
