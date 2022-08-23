package controllers

import (
	"fmt"
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
		var sellerBody SellerRegisterBody
		if bodyErr := ctx.ShouldBindBodyWith(&sellerBody, binding.JSON); bodyErr != nil {
			fmt.Println("body: ", bodyErr)
			ctx.Status(http.StatusInternalServerError)
			return
		}
		if sellerBody.Password != sellerBody.PasswordAgain {
			fmt.Println("passwords does not match")
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "passwords does not match"})
			return
		}

		sellerFromDB, getErr := postgresql.GetSeller(sellerBody.Email)
		if getErr != nil {
			fmt.Println("postgresql (get): ", getErr)
			ctx.Status(http.StatusInternalServerError)
			return
		}
		if sellerFromDB.Email == sellerBody.Email {
			fmt.Println("email already registered")
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "email already registered"})
			return
		}

		salt := dotEnv.GoDotEnvVariable("SALT")
		saltedPassword := sellerBody.Password + salt
		hash, _ := bcrypt.HashPassword(saltedPassword)

		insertErr := postgresql.Insert(sellerBody.CompanyName, sellerBody.Email, hash, sellerBody.Address, sellerBody.PhoneNumber)
		if insertErr != nil {
			fmt.Println("postgresql (insert): ", insertErr)
			ctx.Status(http.StatusInternalServerError)
			return
		}

		fmt.Println("Seller registered: ", sellerBody.Email)
		ctx.JSON(http.StatusOK, gin.H{"message": "OK"})
	}
}

func SellerLoginPostHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var sellerBody LoginBody
		if bodyErr := ctx.ShouldBindBodyWith(&sellerBody, binding.JSON); bodyErr != nil {
			fmt.Println("body: ", bodyErr)
			ctx.Status(http.StatusInternalServerError)
			return
		}

		seller, pqErr := postgresql.GetSeller(sellerBody.Email)
		if pqErr != nil {
			fmt.Println("postgresql (insert): ", pqErr)
			ctx.Status(http.StatusInternalServerError)
			return
		}
		if seller.Email != sellerBody.Email {
			fmt.Println("email not registered")
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "wrong credentials"})
			return
		}

		salt := dotEnv.GoDotEnvVariable("SALT")
		saltedPassword := sellerBody.Password + salt
		match := bcrypt.CheckPasswordHash(saltedPassword, seller.Password)
		if !match {
			fmt.Println("wrong password")
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "wrong credentials"})
			return
		}

		secretToken := dotEnv.GoDotEnvVariable("TOKEN")
		hmacSampleSecret := []byte(secretToken)
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"mail": sellerBody.Email,
			"type": sellerBody.Type,
			"nbf":  time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
		})
		tokenString, tokenErr := token.SignedString(hmacSampleSecret)
		if tokenErr != nil {
			fmt.Println("token: ", tokenErr)
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "bad token"})
			return
		}

		fmt.Println("Seller logged in: ", sellerBody.Email)
		ctx.JSON(http.StatusOK, gin.H{"message": "OK", "token": tokenString})
	}
}

func AddProduct() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var requestBody Product
		auth, authErr := middleware.UserAuth(ctx)

		if authErr != nil {
			fmt.Println("auth: ", authErr)
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "auth error"})
			return
		}
		if bodyErr := ctx.ShouldBindBodyWith(&requestBody, binding.JSON); bodyErr != nil {
			fmt.Println("body: ", bodyErr)
			ctx.Status(http.StatusInternalServerError)
			return
		}

		insertErr := postgresql.InsertProduct(auth.EMail, requestBody.Title, requestBody.Price, requestBody.Description, requestBody.Stock, requestBody.Stock)
		if insertErr != nil {
			fmt.Println("postgresql (insert): ", insertErr)
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "auth error"})
			return
		}

		fmt.Println("Seller: ", auth.EMail, " - Added new product:", requestBody.Title)
		ctx.JSON(http.StatusOK, gin.H{"message": "OK", "mail": auth.EMail, "type": auth.Type})
	}
}

func GetSellerProducts() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		auth, authErr := middleware.UserAuth(ctx)
		if authErr != nil {
			fmt.Println("authentication: ", authErr)
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "auth error"})
			return
		}

		products, pqErr := postgresql.GetSellerProducts(auth.EMail)
		if pqErr != nil {
			fmt.Println("postgresql (get): ", pqErr)
			ctx.Status(http.StatusInternalServerError)
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "OK", "products": products})
	}
}

func DeleteProduct() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		auth, authErr := middleware.UserAuth(ctx)
		if authErr != nil {
			fmt.Println("auth: ", authErr)
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "auth error"})
			return
		}
		var requestBody Product
		if bodyErr := ctx.ShouldBindBodyWith(&requestBody, binding.JSON); bodyErr != nil {
			fmt.Println("body: ", bodyErr)
			ctx.Status(http.StatusInternalServerError)
			return
		}

		delErr := postgresql.DeleteProduct(requestBody.Id)
		if delErr != nil {
			fmt.Println("postgresql (delete): ", delErr)
			ctx.Status(http.StatusInternalServerError)
			return
		}

		fmt.Println("Seller: ", auth.EMail, " - Delete a product:", requestBody.Id)
		ctx.JSON(http.StatusOK, gin.H{"message": "OK"})
	}
}

func GetAllProducts() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		products, err := postgresql.GetAllProducts()
		if err != nil {
			fmt.Println(err)
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"products": products})
	}
}
