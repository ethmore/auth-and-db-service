package controllers

import (
	"fmt"
	"log"
	"time"

	"e-comm/authService/bcrypt"
	"e-comm/authService/dotEnv"

	"e-comm/authService/middleware"

	"e-comm/authService/mongodb"

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
			_, checkedMail, _ := mongodb.FindOneUser(email)
			if checkedMail == email {
				fmt.Println("email already registered")
				ctx.JSON(400, gin.H{"message": "email already registered"})
			} else {
				saltedPassword := password + salt
				hash, _ := bcrypt.HashPassword(saltedPassword)

				res := mongodb.InsertOneUser(name, surname, email, hash)
				if res == 200 {
					ctx.JSON(200, gin.H{"message": "OK"})
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

		_, checkedMail, checkedPassword := mongodb.FindOneUser(requestBody.Email)

		if checkedMail == requestBody.Email {
			saltedPassword := requestBody.Password + salt
			match := bcrypt.CheckPasswordHash(saltedPassword, checkedPassword)

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
				ctx.JSON(200, gin.H{"message": "OK", "token": tokenString})
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
			ctx.JSON(200, gin.H{"message": "OK", "mail": mailAuth, "type": loginType})
		}
	}
}
