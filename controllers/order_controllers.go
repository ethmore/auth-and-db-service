package controllers

import (
	"auth-and-db-service/middleware"
	"auth-and-db-service/repositories/mongodb"
	"auth-and-db-service/repositories/postgresql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func InsertOrder(authenticator middleware.IUserAuthenticator, orderRepo postgresql.IOrderRepo, userRepo mongodb.IUsersRepo) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var orderBody postgresql.Order
		if err := ctx.ShouldBindBodyWith(&orderBody, binding.JSON); err != nil {
			fmt.Println("body err: ", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{})
			return
		}

		auth, authErr := authenticator.UserAuth(ctx)
		if authErr != nil {
			fmt.Println("auth: ", authErr)
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "auth error"})
			return
		}

		user, userErr := userRepo.FindOneUser(auth.EMail)
		if userErr != nil {
			fmt.Println("userErr: ", userErr)
			ctx.JSON(http.StatusInternalServerError, gin.H{})
			return
		}

		id, err := orderRepo.InsertOrder(user.Id, orderBody)
		if err != nil {
			fmt.Println("insert order: ", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{})
			return
		}
		strID := strconv.Itoa(id)

		ctx.JSON(http.StatusOK, gin.H{"message": "OK", "orderID": strID})
	}
}

func GetAllOrders(authenticator middleware.IUserAuthenticator, orderRepo postgresql.IOrderRepo, userRepo mongodb.IUsersRepo) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		auth, authErr := authenticator.UserAuth(ctx)
		if authErr != nil {
			fmt.Println("auth: ", authErr)
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "auth error"})
			return
		}

		user, userErr := userRepo.FindOneUser(auth.EMail)
		if userErr != nil {
			fmt.Println("userErr: ", userErr)
			ctx.JSON(http.StatusInternalServerError, gin.H{})
			return
		}

		strID := user.Id.Hex()
		orders, err := orderRepo.GetAllOrders(strID)
		if err != nil {
			fmt.Println("postgres (get order): ", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{})
			return
		}

		var newOrders []postgresql.Order
		for _, j := range orders {
			orderProducts, ordErr := orderRepo.GetAllOrderProducts(j.ID)
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
