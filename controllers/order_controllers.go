package controllers

import (
	"e-comm/authService/middleware"
	"e-comm/authService/repositories/mongodb"
	"e-comm/authService/repositories/postgresql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func InsertOrder(postgresqlRepo postgresql.IPostgreSQL) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var orderBody postgresql.Order
		if err := ctx.ShouldBindBodyWith(&orderBody, binding.JSON); err != nil {
			fmt.Println("body err: ", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{})
			return
		}

		auth, authErr := middleware.UserAuth(ctx)
		if authErr != nil {
			fmt.Println("auth: ", authErr)
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "auth error"})
			return
		}

		user, userErr := mongodb.FindOneUser(auth.EMail)
		if userErr != nil {
			fmt.Println("userErr: ", userErr)
			ctx.JSON(http.StatusInternalServerError, gin.H{})
			return
		}

		id, err := postgresqlRepo.InsertOrder(user.Id, orderBody)
		if err != nil {
			fmt.Println("insert order: ", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{})
			return
		}
		strID := strconv.Itoa(id)

		ctx.JSON(http.StatusOK, gin.H{"message": "OK", "orderID": strID})
	}
}

func GetAllOrders(postgresqlRepo postgresql.IPostgreSQL) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		auth, authErr := middleware.UserAuth(ctx)
		if authErr != nil {
			fmt.Println("auth: ", authErr)
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "auth error"})
			return
		}

		user, userErr := mongodb.FindOneUser(auth.EMail)
		if userErr != nil {
			fmt.Println("userErr: ", userErr)
			ctx.JSON(http.StatusInternalServerError, gin.H{})
			return
		}

		strID := user.Id.Hex()
		orders, err := postgresqlRepo.GetAllOrders(strID)
		if err != nil {
			fmt.Println("postgres (get order): ", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{})
			return
		}

		var newOrders []postgresql.Order
		for _, j := range orders {
			orderProducts, ordErr := postgresqlRepo.GetAllOrderProducts(j.ID)
			if ordErr != nil {
				fmt.Println("postgres (get order products): ", err)
				ctx.JSON(http.StatusInternalServerError, gin.H{})
				return
			}

			j.Products = orderProducts
			newOrders = append(newOrders, j)
		}
		fmt.Println(newOrders)
		ctx.JSON(http.StatusOK, gin.H{"message": "OK", "orders": newOrders})
	}
}
