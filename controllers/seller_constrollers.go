package controllers

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"e-comm/authService/bcrypt"
	"e-comm/authService/dotEnv"

	"e-comm/authService/middleware"

	"e-comm/authService/repositories/postgresql"

	"github.com/golang-jwt/jwt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type SellerRegisterBody struct {
	CompanyName   string
	Email         string
	Password      string
	PasswordAgain string
	Address       string
	PhoneNumber   string
}

type Product struct {
	Token       string
	Id          string
	Title       string
	Description string
	Price       string
	Stock       string
	Photo       string
}

func SellerRegisterPostHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var requestBody SellerRegisterBody

		if err := ctx.ShouldBindBodyWith(&requestBody, binding.JSON); err != nil {
			log.Printf("%+v", err)
		}

		companyName := requestBody.CompanyName
		email := requestBody.Email
		password := requestBody.Password
		passwordAgain := requestBody.PasswordAgain
		address := requestBody.Address
		phonenumber := requestBody.PhoneNumber

		salt := dotEnv.GoDotEnvVariable("SALT")

		if password == passwordAgain {
			seller, err := postgresql.GetSeller(email)
			if err != nil {
				fmt.Println(err)
			}

			if seller.Email == email {
				fmt.Println("email already registered")
				ctx.JSON(400, gin.H{"message": "email already registered"})
			} else {
				saltedPassword := password + salt
				hash, _ := bcrypt.HashPassword(saltedPassword)

				err := postgresql.Insert(companyName, email, hash, address, phonenumber)
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

func SellerLoginPostHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var requestBody LoginBody
		if err := ctx.ShouldBindBodyWith(&requestBody, binding.JSON); err != nil {
			log.Printf("%+v", err)
		}

		email := requestBody.Email
		password := requestBody.Password

		salt := dotEnv.GoDotEnvVariable("SALT")

		seller, err := postgresql.GetSeller(email)
		if err != nil {
			fmt.Println(err)
		}

		if seller.Email == email {
			saltedPassword := password + salt
			match := bcrypt.CheckPasswordHash(saltedPassword, seller.Password)

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

func SellerDashboard() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var mailAuth, loginType = middleware.UserAuth(ctx)
		if mailAuth != "" {
			ctx.JSON(http.StatusOK, gin.H{"message": "OK", "mail": mailAuth, "type": loginType})
		}
	}
}

func AddProduct() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var mailAuth, loginType = middleware.UserAuth(ctx)
		if mailAuth != "" {
			var requestBody Product
			if err := ctx.ShouldBindBodyWith(&requestBody, binding.JSON); err != nil {
				log.Printf("%+v", err)
			}

			fmt.Println(requestBody.Title, requestBody.Description, requestBody.Price, requestBody.Stock)
			postgresql.InsertProduct(mailAuth, requestBody.Title, requestBody.Price, requestBody.Description, requestBody.Stock, requestBody.Stock)
			ctx.JSON(http.StatusOK, gin.H{"message": "OK", "mail": mailAuth, "type": loginType})
		}

	}
}

func GetSellerProducts() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var mailAuth, _ = middleware.UserAuth(ctx)
		if mailAuth != "" {
			fmt.Println("a")
			products, err := postgresql.GetSellerProducts(mailAuth)
			if err != nil {
				log.Fatal(err)
			}
			ctx.JSON(http.StatusOK, gin.H{"message": "OK", "products": products})
		}
	}
}

func DeleteProduct() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var mailAuth, _ = middleware.UserAuth(ctx)
		if mailAuth != "" {

			var requestBody Product
			if err := ctx.ShouldBindBodyWith(&requestBody, binding.JSON); err != nil {
				log.Printf("%+v", err)
			}

			postgresql.DeleteProduct(requestBody.Id)
			fmt.Println(requestBody.Id)

			ctx.JSON(http.StatusOK, gin.H{"message": "OK"})
		}
	}
}

func GetAllProducts() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		products, err := postgresql.GetAllProducts()
		if err != nil {
			log.Fatal(err)
		}
		ctx.JSON(http.StatusOK, gin.H{"products": products})
	}
}
