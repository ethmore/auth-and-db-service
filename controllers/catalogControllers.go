package controllers

import (
	"e-comm/authService/postgresql"
	"fmt"
	"log"

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
		ctx.JSON(200, gin.H{"products": products})
	}
}
