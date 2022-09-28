package controllers

import (
	"fmt"
	"net/http"

	"auth-and-db-service/middleware"
	"auth-and-db-service/repositories/mongodb"
	"auth-and-db-service/services"

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

func Test() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "OK"})
	}
}

func UserRegisterPostHandler(ur mongodb.IUsersRepo) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var userBody services.UserRegisterBody
		if bodyErr := ctx.ShouldBindBodyWith(&userBody, binding.JSON); bodyErr != nil {
			fmt.Println("body: ", bodyErr)
			ctx.Status(http.StatusInternalServerError)
			return
		}

		if registerErr := services.UserRegister(ur, userBody); registerErr != nil {
			fmt.Println("UserRegister: ", registerErr)
			ctx.Status(http.StatusInternalServerError)
			return
		}

		fmt.Println("User registered: ", userBody.Email)
		ctx.JSON(http.StatusOK, gin.H{"message": "OK"})
	}
}

func UserLoginPostHandler(ur mongodb.IUsersRepo) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var userBody services.LoginBody
		if bodyErr := ctx.ShouldBindBodyWith(&userBody, binding.JSON); bodyErr != nil {
			fmt.Println("body: ", bodyErr)
			ctx.Status(http.StatusBadRequest)
			return
		}

		token, loginErr := services.UserLogin(ur, userBody)
		if loginErr != nil {
			fmt.Println("UserLogin: ", loginErr)
			ctx.Status(http.StatusInternalServerError)
			return
		}

		fmt.Println("User logged in: ", userBody.Email)
		ctx.JSON(http.StatusOK, gin.H{"message": "OK", "token": token})
	}
}

func UserProfile(authenticator middleware.IUserAuthenticator) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		auth, err := authenticator.UserAuth(ctx)
		if err != nil {
			fmt.Println("authentication: ", err)
			ctx.JSON(http.StatusOK, gin.H{"message": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "OK", "mail": auth.EMail, "type": auth.Type})
	}
}

func GetUserInfo(authenticator middleware.IUserAuthenticator, ur mongodb.IUsersRepo) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var body Body
		if bodyErr := ctx.ShouldBindBodyWith(&body, binding.JSON); bodyErr != nil {
			fmt.Println("body: ", bodyErr)
			ctx.Status(http.StatusInternalServerError)
			return
		}

		auth, authErr := authenticator.UserAuth(ctx)
		if authErr != nil {
			fmt.Println("auth:", authErr)
			ctx.Status(http.StatusInternalServerError)
			return
		}

		user, findErr := ur.FindOneUser(auth.EMail)
		if findErr != nil {
			fmt.Println("mongodb (find): ", findErr)
			ctx.Status(http.StatusInternalServerError)
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"message": "OK", "userInfo": user})
	}
}

func ChangeUserPassword(authenticator middleware.IUserAuthenticator, ur mongodb.IUsersRepo) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var passBody services.ChangePassword
		if bodyErr := ctx.ShouldBindBodyWith(&passBody, binding.JSON); bodyErr != nil {
			fmt.Println("body: ", bodyErr)
			ctx.Status(http.StatusInternalServerError)
			return
		}

		auth, authErr := authenticator.UserAuth(ctx)
		if authErr != nil {
			fmt.Println("auth:", authErr)
			ctx.Status(http.StatusInternalServerError)
			return
		}

		if changePassErr := services.ChangeUserPassword(ur, passBody, auth.EMail); changePassErr != nil {
			fmt.Println("ChangeUserPassword:", changePassErr)
			ctx.Status(http.StatusInternalServerError)
			return
		}

		fmt.Println("User changed password: ", auth.EMail)
		ctx.JSON(http.StatusOK, gin.H{"message": "OK"})
	}
}

func NewUserAddress(authenticator middleware.IUserAuthenticator, ur mongodb.IUsersRepo, uar mongodb.IUserAddressesRepo) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		auth, authErr := authenticator.UserAuth(ctx)
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

		insertErr := uar.InsertUserAddress(ur, auth.EMail, addressBody.Title, addressBody.Name, addressBody.Surname, addressBody.PhoneNumber, addressBody.Province, addressBody.County, addressBody.DetailedAddress)
		if insertErr != nil {
			fmt.Println("mongodb (insertAddress): ", insertErr)
			ctx.Status(http.StatusInternalServerError)
			return
		}

		fmt.Println("User: ", auth.EMail, " - Added new address: ", addressBody.Title)
		ctx.JSON(http.StatusOK, gin.H{"message": "OK"})
	}
}

func GetUserAddressById(authenticator middleware.IUserAuthenticator, uar mongodb.IUserAddressesRepo) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		_, authErr := authenticator.UserAuth(ctx)
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

		address, getErr := uar.FindUserAddress(addressBody.AddressId)
		if getErr != nil {
			fmt.Println("mongodb (getAddresses): ", getErr)
			ctx.Status(http.StatusInternalServerError)
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "OK", "address": address})
	}
}

func GetUserAddresses(authenticator middleware.IUserAuthenticator, ur mongodb.IUsersRepo, uar mongodb.IUserAddressesRepo) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		auth, authErr := authenticator.UserAuth(ctx)
		if authErr != nil {
			fmt.Println("auth: ", authErr)
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "auth error"})
			return
		}

		addresses, getErr := uar.FindAllUserAddresses(ur, auth.EMail)
		if getErr != nil {
			fmt.Println("mongodb (getAddresses): ", getErr)
			ctx.Status(http.StatusInternalServerError)
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"message": "OK", "addresses": addresses})
	}
}
