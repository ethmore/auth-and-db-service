package routes

import (
	"e-comm/authService/controllers"
	"e-comm/authService/repositories/postgresql"

	"github.com/gin-gonic/gin"
)

func PublicRoutes(g *gin.RouterGroup) {
	postgresql := new(postgresql.Postgresql)
	g.POST("/sellerRegister", controllers.SellerRegisterPostHandler(postgresql))
	g.POST("/userRegister", controllers.UserRegisterPostHandler())
	g.POST("/sellerLogin", controllers.SellerLoginPostHandler(postgresql))
	g.POST("/userLogin", controllers.UserLoginPostHandler())
	g.POST("/getAllProducts", controllers.GetAllProducts(postgresql))
	g.POST("/getProduct", controllers.GetProduct(postgresql))
}

func PrivateRoutes(g *gin.RouterGroup) {
	postgresql := new(postgresql.Postgresql)
	g.POST("/profile", controllers.UserProfile())
	g.POST("/getUserInfo", controllers.GetUserInfo())
	g.POST("/changeUserPassword", controllers.ChangeUserPassword())

	g.POST("/addProduct", controllers.AddProduct(postgresql))
	g.POST("/getProducts", controllers.GetSellerProducts(postgresql))
	g.POST("/deleteProduct", controllers.DeleteProduct(postgresql))
	g.POST("/editProduct", controllers.EditProduct(postgresql))

	g.POST("/newUserAddress", controllers.NewUserAddress())
	g.POST("/getUserAddressById", controllers.GetUserAddressById())
	g.POST("/getUserAddresses", controllers.GetUserAddresses())

	g.POST("/addProductToCart", controllers.AddProductToCart(postgresql))
	g.POST("/getCartProducts", controllers.GetCartProducts())
	g.POST("/getCartInfo", controllers.GetCartInfo(postgresql))
	g.POST("/removeProductFromCart", controllers.RemoveProductFromCart())
	g.POST("/changeProductQty", controllers.ChangeProductQty())
	// g.POST("/increaseProductQty", controllers.IncreaseProductQty())
	// g.POST("/decreaseProductQty", controllers.DecreaseProductQty())
	g.POST("/addTotalToCart", controllers.AddTotalToCart())
	g.POST("/getTotalPrice", controllers.GetTotalPrice())

	g.POST("/clearCart", controllers.ClearCart())

	g.POST("/createPayment", controllers.CreatePayment(postgresql))
	g.POST("/updatePaymentStatus", controllers.UpdatePaymentStatus(postgresql))

	g.POST("/getProductsSellers", controllers.GetProductsSellers(postgresql))

	g.POST("/insertOrder", controllers.InsertOrder(postgresql))
	g.POST("/getAllOrders", controllers.GetAllOrders(postgresql))
}
