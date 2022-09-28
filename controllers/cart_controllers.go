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
	"go.mongodb.org/mongo-driver/mongo"
)

type CartRequest struct {
	Token string
	Id    string
	Qty   string
}

type Body struct {
	Token string
}

type PurchaseBody struct {
	Token      string
	TotalPrice string
}

type Item struct {
	Id         string
	Title      string
	TotalPrice string
}

type CartInfo struct {
	Id             string
	Items          []Item
	TotalCartPrice string
}

func AddProductToCart(authenticator middleware.IUserAuthenticator, productRepo postgresql.IProductRepo, ur mongodb.IUsersRepo, userCartRepo mongodb.IUserCartRepo) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		auth, authErr := authenticator.UserAuth(ctx)
		if authErr != nil {
			fmt.Println("auth: ", authErr)
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "auth error"})
			return
		}

		var cartRequest CartRequest
		if bodyErr := ctx.ShouldBindBodyWith(&cartRequest, binding.JSON); bodyErr != nil {
			fmt.Println("body: ", bodyErr)
			ctx.Status(http.StatusInternalServerError)
			return
		}

		addErr := userCartRepo.AddProductToCart(productRepo, ur, auth.EMail, cartRequest.Id, cartRequest.Qty)
		if addErr == mongo.ErrNoDocuments {
			if err := userCartRepo.NewCart(ur, auth.EMail); err != nil {
				fmt.Println("mongodb (new-cart): ", err)
				ctx.Status(http.StatusInternalServerError)
				return
			}

			err := userCartRepo.AddProductToCart(productRepo, ur, auth.EMail, cartRequest.Id, cartRequest.Qty)
			if err != nil {
				fmt.Println("mongodb (add-2): ", addErr)
				ctx.Status(http.StatusInternalServerError)
				return
			}

			ctx.JSON(http.StatusOK, gin.H{"message": "OK"})
			return
		}
		if addErr != nil {
			fmt.Println("mongodb (add): ", addErr)
			ctx.Status(http.StatusInternalServerError)
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"message": "OK"})
	}
}

func GetCartInfo(authenticator middleware.IUserAuthenticator, productRepo postgresql.IProductRepo, ur mongodb.IUsersRepo, userCartRepo mongodb.IUserCartRepo) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		auth, authErr := authenticator.UserAuth(ctx)
		if authErr != nil {
			fmt.Println("auth: ", authErr)
			ctx.Status(http.StatusInternalServerError)
			return
		}

		cart, err := userCartRepo.FindAllCartProducts(ur, auth.EMail)
		if err != nil {
			fmt.Println("mongodb (find): ", err)
			ctx.Status(http.StatusInternalServerError)
			return
		}

		var cartInfo CartInfo
		for i := 0; i < len(cart.Products); i++ {
			product, getErr := productRepo.GetProduct(cart.Products[i].Id)
			if getErr != nil {
				fmt.Println("postgresql (get): ", getErr)
				ctx.Status(http.StatusInternalServerError)
				return
			}

			itemQty, convErr := strconv.Atoi(cart.Products[i].Qty)
			if convErr != nil {
				fmt.Println("convErr: ", convErr)
				ctx.Status(http.StatusInternalServerError)
				return
			}
			itemPrice, convErr2 := strconv.Atoi(product.Price)
			if convErr2 != nil {
				fmt.Println("convErr2: ", convErr2)
				ctx.Status(http.StatusInternalServerError)
				return
			}
			totalItemPrice := itemQty * itemPrice
			totalItemPriceStr := strconv.Itoa(totalItemPrice)

			var newProduct = Item{
				Id:         product.Id,
				Title:      product.Title,
				TotalPrice: totalItemPriceStr,
			}

			cartInfo.Items = append(cartInfo.Items, newProduct)
		}

		cartInfo.Id = cart.Id
		cartInfo.TotalCartPrice = cart.TotalPrice

		ctx.JSON(http.StatusOK, gin.H{"message": "OK", "cartInfo": cartInfo})
	}
}

func GetCartProducts(authenticator middleware.IUserAuthenticator, ur mongodb.IUsersRepo, userCartRepo mongodb.IUserCartRepo) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		auth, authErr := authenticator.UserAuth(ctx)
		if authErr != nil {
			fmt.Println("auth: ", authErr)
			ctx.Status(http.StatusInternalServerError)
			return
		}

		cart, err := userCartRepo.FindAllCartProducts(ur, auth.EMail)
		if err != nil {
			fmt.Println("mongodb (find): ", err)
			ctx.Status(http.StatusInternalServerError)
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"message": "OK", "products": cart.Products})
	}
}

func RemoveProductFromCart(authenticator middleware.IUserAuthenticator, ur mongodb.IUsersRepo, userCartRepo mongodb.IUserCartRepo) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		auth, authErr := authenticator.UserAuth(ctx)
		if authErr != nil {
			fmt.Println("auth: ", authErr)
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "auth error"})
			return
		}

		var cartReq CartRequest
		if bodyErr := ctx.ShouldBindBodyWith(&cartReq, binding.JSON); bodyErr != nil {
			fmt.Println("body: ", bodyErr)
			ctx.Status(http.StatusInternalServerError)
			return
		}

		removeErr := userCartRepo.RemoveProductFromCart(ur, auth.EMail, cartReq.Id)
		if removeErr != nil {
			fmt.Println("mongodb (remove): ", removeErr)
			ctx.Status(http.StatusInternalServerError)
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"message": "OK"})
	}
}

func ChangeProductQty(authenticator middleware.IUserAuthenticator, ur mongodb.IUsersRepo, userCartRepo mongodb.IUserCartRepo) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		auth, authErr := authenticator.UserAuth(ctx)
		if authErr != nil {
			fmt.Println("auth: ", authErr)
			ctx.Status(http.StatusInternalServerError)
			return
		}

		var cartReq CartRequest
		if bodyErr := ctx.ShouldBindBodyWith(&cartReq, binding.JSON); bodyErr != nil {
			fmt.Println("body: ", authErr)
			ctx.Status(http.StatusInternalServerError)
			return
		}

		err := userCartRepo.ChangeProductQty(ur, auth.EMail, cartReq.Id, cartReq.Qty)
		if err != nil {
			fmt.Println("mongodb (update): ", err)
			ctx.Status(http.StatusInternalServerError)
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"message": "OK"})
	}
}

func AddTotalToCart(authenticator middleware.IUserAuthenticator, ur mongodb.IUsersRepo, userCartRepo mongodb.IUserCartRepo) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var purchaseBody PurchaseBody
		if bodyErr := ctx.ShouldBindBodyWith(&purchaseBody, binding.JSON); bodyErr != nil {
			fmt.Println("body: ", bodyErr)
			ctx.Status(http.StatusInternalServerError)
			return
		}

		auth, authErr := authenticator.UserAuth(ctx)
		if authErr != nil {
			fmt.Println("auth: ", authErr)
			ctx.Status(http.StatusInternalServerError)
			return
		}

		if err := userCartRepo.AddTotalToCart(ur, auth.EMail, purchaseBody.TotalPrice); err != nil {
			fmt.Println("add total to cart: ", err)
			ctx.Status(http.StatusInternalServerError)
			return
		}

		ctx.Status(http.StatusOK)
	}
}

func GetTotalPrice(authenticator middleware.IUserAuthenticator, ur mongodb.IUsersRepo, userCartRepo mongodb.IUserCartRepo) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		auth, authErr := authenticator.UserAuth(ctx)
		if authErr != nil {
			fmt.Println("auth: ", authErr)
			ctx.Status(http.StatusInternalServerError)
			return
		}

		totalPrice, err := userCartRepo.GetTotalPrice(ur, auth.EMail)
		if err != nil {
			fmt.Println("add total to cart: ", err)
			ctx.Status(http.StatusInternalServerError)
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"message": "OK", "totalPrice": totalPrice})
	}
}

func ClearCart(authenticator middleware.IUserAuthenticator, ur mongodb.IUsersRepo, userCartRepo mongodb.IUserCartRepo) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		auth, authErr := authenticator.UserAuth(ctx)
		if authErr != nil {
			fmt.Println("auth: ", authErr)
			ctx.Status(http.StatusInternalServerError)
			return
		}

		clrErr := userCartRepo.ClearCart(ur, auth.EMail)
		if clrErr != nil {
			fmt.Println("mongodb (clear): ", clrErr)
			ctx.Status(http.StatusInternalServerError)
			return
		}

		ctx.Status(http.StatusOK)
	}
}
