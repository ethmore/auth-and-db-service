package controllers

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"e-comm/authService/bcrypt"
	"e-comm/authService/dotEnv"

	"e-comm/authService/middleware"

	"e-comm/authService/repositories/mongodb"

	"github.com/golang-jwt/jwt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type UserRegisterBody struct {
	Name          string
	Surname       string
	Email         string
	Password      string
	PasswordAgain string
}

type LoginBody struct {
	Email    string
	Password string
	Type     string
}

func UserRegisterPostHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var requestBody UserRegisterBody
		if err := ctx.ShouldBindBodyWith(&requestBody, binding.JSON); err != nil {
			log.Printf("%+v", err)
		}

		name := requestBody.Name
		surname := requestBody.Surname
		email := requestBody.Email
		password := requestBody.Password
		passwordAgain := requestBody.PasswordAgain

		salt := dotEnv.GoDotEnvVariable("SALT")

		if password == passwordAgain {
			user, err := mongodb.FindOneUser(email)
			if err != nil {
				fmt.Println(err)
			}
			if user.Email == email {
				fmt.Println("email already registered")
				ctx.JSON(400, gin.H{"message": "email already registered"})
			} else {
				saltedPassword := password + salt
				hash, _ := bcrypt.HashPassword(saltedPassword)

				err := mongodb.InsertOneUser(name, surname, email, hash)
				if err == nil {
					ctx.JSON(http.StatusOK, gin.H{"message": "OK"})
				}

			}

		} else {
			fmt.Println("passwords does not match")
			ctx.JSON(400, gin.H{"message": "passwords does not match"})
		}
	}
}

func UserLoginPostHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var requestBody LoginBody
		if err := ctx.ShouldBindBodyWith(&requestBody, binding.JSON); err != nil {
			log.Printf("%+v", err)
		}

		salt := dotEnv.GoDotEnvVariable("SALT")

		user, err := mongodb.FindOneUser(requestBody.Email)
		if err != nil {
			fmt.Println(err)
		}

		if user.Email == requestBody.Email {
			saltedPassword := requestBody.Password + salt
			match := bcrypt.CheckPasswordHash(saltedPassword, user.Password)

			if match {
				secretToken := dotEnv.GoDotEnvVariable("TOKEN")
				hmacSampleSecret := []byte(secretToken)

				token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
					"mail": requestBody.Email,
					"type": requestBody.Type,
					"nbf":  time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
				})
				tokenString, err := token.SignedString(hmacSampleSecret)
				if err != nil {
					fmt.Println(err)
				}

				fmt.Println("OK")
				ctx.JSON(http.StatusOK, gin.H{"message": "OK", "token": tokenString})
			} else {
				fmt.Println("wrong password")
				ctx.JSON(400, gin.H{"message": "wrong password"})
			}
		} else {
			fmt.Println("email not registered")
			ctx.JSON(400, gin.H{"message": "email not registered"})
		}
	}
}

func UserProfile() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var mailAuth, loginType = middleware.UserAuth(ctx)
		if mailAuth != "" {
			ctx.JSON(http.StatusOK, gin.H{"message": "OK", "mail": mailAuth, "type": loginType})
		}
	}
}
