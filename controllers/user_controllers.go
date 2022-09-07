package controllers

import (
	"fmt"
	"net/http"

	"e-comm/authService/middleware"
	"e-comm/authService/repositories/mongodb"
	"e-comm/authService/services"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type AddressBody struct {
	Title           string
	Name            string
	Surname         string
	PhoneNumber     string
	Province        string
	County          string
	DetailedAddress string
}

type PaymentAddress struct {
	Token     string
	AddressId string
}

func UserRegisterPostHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var userBody services.UserRegisterBody
		if bodyErr := ctx.ShouldBindBodyWith(&userBody, binding.JSON); bodyErr != nil {
			fmt.Println("body: ", bodyErr)
			ctx.Status(http.StatusInternalServerError)
			return
		}

		if registerErr := services.UserRegister(userBody); registerErr != nil {
			fmt.Println("UserRegister: ", registerErr)
			ctx.Status(http.StatusInternalServerError)
			return
		}

		fmt.Println("User registered: ", userBody.Email)
		ctx.JSON(http.StatusOK, gin.H{"message": "OK"})
	}
}

func UserLoginPostHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var userBody services.LoginBody
		if bodyErr := ctx.ShouldBindBodyWith(&userBody, binding.JSON); bodyErr != nil {
			fmt.Println("body: ", bodyErr)
			ctx.Status(http.StatusInternalServerError)
			return
		}

		token, loginErr := services.UserLogin(userBody)
		if loginErr != nil {
			fmt.Println("UserLogin: ", loginErr)
			ctx.Status(http.StatusInternalServerError)
			return
		}

		fmt.Println("User logged in: ", userBody.Email)
		ctx.JSON(http.StatusOK, gin.H{"message": "OK", "token": token})
	}
}

func UserProfile() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		auth, err := middleware.UserAuth(ctx)
		if err != nil {
			fmt.Println("authentication: ", err)
			ctx.JSON(http.StatusOK, gin.H{"message": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "OK", "mail": auth.EMail, "type": auth.Type})
	}
}

func GetUserInfo() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var body Body
		if bodyErr := ctx.ShouldBindBodyWith(&body, binding.JSON); bodyErr != nil {
			fmt.Println("body: ", bodyErr)
			ctx.Status(http.StatusInternalServerError)
			return
		}

		auth, authErr := middleware.UserAuth(ctx)
		if authErr != nil {
			fmt.Println("auth:", authErr)
			ctx.Status(http.StatusInternalServerError)
			return
		}

		user, findErr := mongodb.FindOneUser(auth.EMail)
		if findErr != nil {
			fmt.Println("mongodb (find): ", findErr)
			ctx.Status(http.StatusInternalServerError)
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"message": "OK", "userInfo": user})
	}
}

func ChangeUserPassword() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var passBody services.ChangePassword
		if bodyErr := ctx.ShouldBindBodyWith(&passBody, binding.JSON); bodyErr != nil {
			fmt.Println("body: ", bodyErr)
			ctx.Status(http.StatusInternalServerError)
			return
		}

		auth, authErr := middleware.UserAuth(ctx)
		if authErr != nil {
			fmt.Println("auth:", authErr)
			ctx.Status(http.StatusInternalServerError)
			return
		}

		if changePassErr := services.ChangeUserPassword(passBody, auth.EMail); changePassErr != nil {
			fmt.Println("ChangeUserPassword:", changePassErr)
			ctx.Status(http.StatusInternalServerError)
			return
		}

		fmt.Println("User changed password: ", auth.EMail)
		ctx.JSON(http.StatusOK, gin.H{"message": "OK"})
	}
}

func NewUserAddress() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		auth, authErr := middleware.UserAuth(ctx)
		if authErr != nil {
			fmt.Println("auth: ", authErr)
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "auth error"})
			return
		}

		var addressBody AddressBody
		if bodyErr := ctx.ShouldBindBodyWith(&addressBody, binding.JSON); bodyErr != nil {
			fmt.Println("body: ", bodyErr)
			ctx.Status(http.StatusInternalServerError)
			return
		}

		insertErr := mongodb.InsertUserAddress(auth.EMail, addressBody.Title, addressBody.Name, addressBody.Surname, addressBody.PhoneNumber, addressBody.Province, addressBody.County, addressBody.DetailedAddress)
		if insertErr != nil {
			fmt.Println("mongodb (insertAddress): ", insertErr)
			ctx.Status(http.StatusInternalServerError)
			return
		}

		fmt.Println("User: ", auth.EMail, " - Added new address: ", addressBody.Title)
		ctx.JSON(http.StatusOK, gin.H{"message": "OK"})
	}
}

func GetUserAddressById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		_, authErr := middleware.UserAuth(ctx)
		if authErr != nil {
			fmt.Println("auth: ", authErr)
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "auth error"})
			return
		}

		var addressBody PaymentAddress
		if bodyErr := ctx.ShouldBindBodyWith(&addressBody, binding.JSON); bodyErr != nil {
			fmt.Println("body: ", bodyErr)
			ctx.Status(http.StatusInternalServerError)
			return
		}

		address, getErr := mongodb.FindUserAddress(addressBody.AddressId)
		if getErr != nil {
			fmt.Println("mongodb (getAddresses): ", getErr)
			ctx.Status(http.StatusInternalServerError)
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "OK", "address": address})
	}
}

func GetUserAddresses() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		auth, authErr := middleware.UserAuth(ctx)
		if authErr != nil {
			fmt.Println("auth: ", authErr)
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "auth error"})
			return
		}

		addresses, getErr := mongodb.FindAllUserAddresses(auth.EMail)
		if getErr != nil {
			fmt.Println("mongodb (getAddresses): ", getErr)
			ctx.Status(http.StatusInternalServerError)
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"message": "OK", "addresses": addresses})
	}
}
