package main

import (
	"fmt"

	"e-comm/authService/dotEnv"
	"e-comm/authService/mongodb"
	"e-comm/authService/postgresql"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	go postgresql.Connect()
	go mongodb.Connect()

	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3001/seller-register", "http://localhost:3001"}
	router.Use(cors.New(config))

	type SellerRegisterBody struct {
		CompanyName   string
		Email         string
		Password      string
		PasswordAgain string
		Address       string
		PhoneNumber   string
	}

	router.POST("/sellerRegister", func(ctx *gin.Context) {
		var requestBody SellerRegisterBody

		if err := ctx.BindJSON(&requestBody); err != nil {
			fmt.Println(err)
		}

		companyName := requestBody.CompanyName
		email := requestBody.Email
		password := requestBody.Password
		passwordAgain := requestBody.PasswordAgain
		address := requestBody.Address
		phonenumber := requestBody.PhoneNumber

		salt := dotEnv.GoDotEnvVariable("SALT")
		fmt.Println("Credentials:   ", email, " : ", password, " : ", passwordAgain, " : ", companyName, "", address, "", phonenumber)

		if password == passwordAgain {
			checkedMail := postgresql.GetSeller(email)
			if checkedMail == email {
				fmt.Println("User Already Registered")
			} else {
				saltedPassword := password + salt
				hash, _ := HashPassword(saltedPassword)

				fmt.Println("Salted Password:", saltedPassword)
				fmt.Println("Hash:    ", hash)

				match := CheckPasswordHash(saltedPassword, hash)
				fmt.Println("Match:   ", match)

				postgresql.Insert(companyName, email, hash, address, phonenumber)
			}

		} else {
			fmt.Println("Passwords does not match")
		}

	})

	type UserRegisterBody struct {
		Name          string
		Surname       string
		Email         string
		Password      string
		PasswordAgain string
	}

	router.POST("/userRegister", func(ctx *gin.Context) {
		var requestBody UserRegisterBody

		if err := ctx.BindJSON(&requestBody); err != nil {
			fmt.Println(err)
		}

		name := requestBody.Name
		surname := requestBody.Surname
		email := requestBody.Email
		password := requestBody.Password
		passwordAgain := requestBody.PasswordAgain

		salt := dotEnv.GoDotEnvVariable("SALT")
		fmt.Println("Credentials:   ", email, " : ", password, " : ", passwordAgain, " : ", name, "", surname, salt)

		ctx.JSON(200, `message: "OK-Success"`)
		if password == passwordAgain {
			checkedMail := mongodb.FindOneUser(email)
			if checkedMail == email {
				fmt.Println("User Already Registered")
			} else {
				saltedPassword := password + salt
				hash, _ := HashPassword(saltedPassword)

				fmt.Println("Salted Password:", saltedPassword)
				fmt.Println("Hash:    ", hash)

				match := CheckPasswordHash(saltedPassword, hash)
				fmt.Println("Match:   ", match)

				mongodb.InsertOneUser(name, surname, email, hash)
			}

		} else {
			fmt.Println("Passwords does not match")
		}

	})

	router.Run(":3002")
}
