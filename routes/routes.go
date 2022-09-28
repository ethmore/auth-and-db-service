package routes

import (
	"auth-and-db-service/controllers"
	"auth-and-db-service/middleware"
	"auth-and-db-service/repositories/mongodb"
	"auth-and-db-service/repositories/postgresql"

	"github.com/gin-gonic/gin"
)

func PublicRoutes(g *gin.RouterGroup) {
	sellerRepo := &postgresql.SellerRepo{}
	productRepo := &postgresql.ProductRepo{}

	usersRepo := &mongodb.UsersRepo{}

	g.GET("/test", controllers.Test())
	g.POST("/sellerRegister", controllers.SellerRegisterPostHandler(sellerRepo))
	g.POST("/userRegister", controllers.UserRegisterPostHandler(usersRepo))
	g.POST("/sellerLogin", controllers.SellerLoginPostHandler(sellerRepo))
	g.POST("/userLogin", controllers.UserLoginPostHandler(usersRepo))
	g.POST("/getAllProducts", controllers.GetAllProducts(productRepo))
	g.POST("/getProduct", controllers.GetProduct(productRepo))
}

func PrivateRoutes(g *gin.RouterGroup) {
	authenticator := &middleware.UserAuthenticator{}

	sellerRepo := &postgresql.SellerRepo{}
	productRepo := &postgresql.ProductRepo{}
	paymentRepo := &postgresql.PaymentRepo{}
	orderRepo := &postgresql.OrderRepo{}

	usersRepo := &mongodb.UsersRepo{}
	usersAddressRepo := &mongodb.UserAddressesRepo{}
	usersCartRepo := &mongodb.UserCartRepo{}

	g.POST("/profile", controllers.UserProfile(authenticator))
	g.POST("/getUserInfo", controllers.GetUserInfo(authenticator, usersRepo))
	g.POST("/changeUserPassword", controllers.ChangeUserPassword(authenticator, usersRepo))

	g.POST("/addProduct", controllers.AddProduct(authenticator, productRepo, sellerRepo))
	g.POST("/getProducts", controllers.GetSellerProducts(authenticator, productRepo, sellerRepo))
	g.POST("/deleteProduct", controllers.DeleteProduct(authenticator, productRepo))
	g.POST("/editProduct", controllers.EditProduct(authenticator, productRepo))

	g.POST("/newUserAddress", controllers.NewUserAddress(authenticator, usersRepo, usersAddressRepo))
	g.POST("/getUserAddressById", controllers.GetUserAddressById(authenticator, usersAddressRepo))
	g.POST("/getUserAddresses", controllers.GetUserAddresses(authenticator, usersRepo, usersAddressRepo))

	g.POST("/addProductToCart", controllers.AddProductToCart(authenticator, productRepo, usersRepo, usersCartRepo))
	g.POST("/getCartProducts", controllers.GetCartProducts(authenticator, usersRepo, usersCartRepo))
	g.POST("/getCartInfo", controllers.GetCartInfo(authenticator, productRepo, usersRepo, usersCartRepo))
	g.POST("/removeProductFromCart", controllers.RemoveProductFromCart(authenticator, usersRepo, usersCartRepo))
	g.POST("/changeProductQty", controllers.ChangeProductQty(authenticator, usersRepo, usersCartRepo))
	g.POST("/addTotalToCart", controllers.AddTotalToCart(authenticator, usersRepo, usersCartRepo))
	g.POST("/getTotalPrice", controllers.GetTotalPrice(authenticator, usersRepo, usersCartRepo))

	g.POST("/clearCart", controllers.ClearCart(authenticator, usersRepo, usersCartRepo))

	g.POST("/createPayment", controllers.CreatePayment(paymentRepo))
	g.POST("/updatePaymentStatus", controllers.UpdatePaymentStatus(paymentRepo))

	g.POST("/getProductsSellers", controllers.GetProductsSellers(authenticator, productRepo, sellerRepo))

	g.POST("/insertOrder", controllers.InsertOrder(authenticator, orderRepo, usersRepo))
	g.POST("/getAllOrders", controllers.GetAllOrders(authenticator, orderRepo, usersRepo))
}
