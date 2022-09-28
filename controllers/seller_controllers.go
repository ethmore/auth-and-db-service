package controllers

import (
	"fmt"
	"net/http"

	"auth-and-db-service/middleware"
	"auth-and-db-service/repositories/postgresql"
	"auth-and-db-service/services"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type Product struct {
	Token       string
	Id          string
	Title       string
	Description string
	Price       string
	Stock       string
	Photo       string
}

type ProductsSeller struct {
	Token      string
	ProductIDs []string
}

type ProductAndSeller struct {
	ProductID string
	Seller    string
}

func SellerRegisterPostHandler(sellerRepo postgresql.ISellerRepo) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var sellerBody services.SellerRegisterBody
		if bodyErr := ctx.ShouldBindBodyWith(&sellerBody, binding.JSON); bodyErr != nil {
			fmt.Println("body: ", bodyErr)
			ctx.Status(http.StatusInternalServerError)
			return
		}

		if registerErr := services.SellerRegister(sellerRepo, sellerBody); registerErr != nil {
			fmt.Println("SellerRegister: ", registerErr)
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": registerErr})
			return
		}

		fmt.Println("Seller registered: ", sellerBody.Email)
		ctx.JSON(http.StatusOK, gin.H{"message": "OK"})
	}
}

func SellerLoginPostHandler(sellerRepo postgresql.ISellerRepo) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var sellerBody services.LoginBody
		if bodyErr := ctx.ShouldBindBodyWith(&sellerBody, binding.JSON); bodyErr != nil {
			fmt.Println("body: ", bodyErr)
			ctx.Status(http.StatusInternalServerError)
			return
		}

		token, loginErr := services.SellerLogin(sellerRepo, sellerBody)
		if loginErr != nil {
			fmt.Println("SellerLogin: ", loginErr)
			ctx.Status(http.StatusInternalServerError)
			return
		}

		fmt.Println("Seller logged in: ", sellerBody.Email)
		ctx.JSON(http.StatusOK, gin.H{"message": "OK", "token": token})
	}
}

func AddProduct(authenticator middleware.IUserAuthenticator, productRepo postgresql.IProductRepo, sellerRepo postgresql.ISellerRepo) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var requestBody Product
		auth, authErr := authenticator.UserAuth(ctx)

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

		insertErr := productRepo.InsertProduct(sellerRepo, auth.EMail, requestBody.Title, requestBody.Price, requestBody.Description, requestBody.Stock, requestBody.Stock)
		if insertErr != nil {
			fmt.Println("postgresql (insert): ", insertErr)
			ctx.Status(http.StatusInternalServerError)
			return
		}

		var searchBody = services.SearchProduct{
			Id:          requestBody.Id,
			Title:       requestBody.Title,
			Price:       requestBody.Price,
			Description: requestBody.Description,
			Image:       requestBody.Photo,
			Stock:       requestBody.Stock,
		}
		searchService := new(services.SearchService)
		searchIndexErr := searchService.AddProductToSearchService(searchBody)
		if searchIndexErr != nil {
			fmt.Println("AddProductToSearchService ", searchIndexErr)
			ctx.Status(http.StatusInternalServerError)
			return
		}

		fmt.Println("Seller: ", auth.EMail, " - Added new product:", requestBody.Title)
		ctx.JSON(http.StatusOK, gin.H{"message": "OK", "mail": auth.EMail, "type": auth.Type})
	}
}

func EditProduct(authenticator middleware.IUserAuthenticator, productRepo postgresql.IProductRepo) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var requestBody Product
		auth, authErr := authenticator.UserAuth(ctx)

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

		updateStmt := productRepo.UpdateProduct(requestBody.Title, requestBody.Price, requestBody.Description, requestBody.Photo, requestBody.Stock, requestBody.Id)
		if updateStmt != nil {
			fmt.Println("postgresql (insert): ", updateStmt)
			ctx.Status(http.StatusInternalServerError)
			return
		}

		fmt.Println("Seller: ", auth.EMail, " - Edited a product:", requestBody.Title)
		ctx.JSON(http.StatusOK, gin.H{"message": "OK", "mail": auth.EMail, "type": auth.Type})

	}
}

func GetSellerProducts(authenticator middleware.IUserAuthenticator, postgresqlRepo postgresql.IProductRepo, sellerRepo postgresql.ISellerRepo) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		auth, authErr := authenticator.UserAuth(ctx)
		if authErr != nil {
			fmt.Println("authentication: ", authErr)
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "auth error"})
			return
		}

		if auth.Type != "seller" {
			fmt.Println("authentication: ", "type error")
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "type error"})
			return
		}

		products, pqErr := postgresqlRepo.GetSellerProducts(sellerRepo, auth.EMail)
		if pqErr != nil {
			fmt.Println("postgresql (get): ", pqErr)
			ctx.Status(http.StatusInternalServerError)
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "OK", "products": products})
	}
}

func DeleteProduct(authenticator middleware.IUserAuthenticator, productRepo postgresql.IProductRepo) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		auth, authErr := authenticator.UserAuth(ctx)
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

		delErr := productRepo.DeleteProduct(requestBody.Id)
		if delErr != nil {
			fmt.Println("postgresql (delete): ", delErr)
			ctx.Status(http.StatusInternalServerError)
			return
		}

		fmt.Println("Seller: ", auth.EMail, " - Delete a product:", requestBody.Id)
		ctx.JSON(http.StatusOK, gin.H{"message": "OK"})
	}
}

func GetProductsSellers(authenticator middleware.IUserAuthenticator, productRepo postgresql.IProductRepo, sellerRepo postgresql.ISellerRepo) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		_, authErr := authenticator.UserAuth(ctx)
		if authErr != nil {
			fmt.Println("auth: ", authErr)
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "auth error"})
			return
		}

		var requestBody ProductsSeller
		if bodyErr := ctx.ShouldBindBodyWith(&requestBody, binding.JSON); bodyErr != nil {
			fmt.Println("body: ", bodyErr)
			ctx.Status(http.StatusInternalServerError)
			return
		}

		var productSeller []ProductAndSeller
		for _, j := range requestBody.ProductIDs {
			product, err := productRepo.GetProduct(j)
			if err != nil {
				fmt.Println("postgresql (get): ", err)
				ctx.Status(http.StatusInternalServerError)
				return
			}

			seller, err := sellerRepo.GetSellerNameByID(product.SellerID)
			if err != nil {
				fmt.Println("postgresql (get-seller): ", err)
				ctx.Status(http.StatusInternalServerError)
				return
			}
			fmt.Println(seller)
			var prodSeller ProductAndSeller
			prodSeller.ProductID = product.Id
			prodSeller.Seller = seller.CompanyName

			productSeller = append(productSeller, prodSeller)
		}

		ctx.JSON(http.StatusOK, gin.H{"message": "OK", "productsAndSellers": productSeller})
	}
}
