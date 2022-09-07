package routes

import (
	"e-comm/authService/controllers"

	"github.com/gin-gonic/gin"
)

func PublicRoutes(g *gin.RouterGroup) {
	g.POST("/sellerRegister", controllers.SellerRegisterPostHandler())
	g.POST("/userRegister", controllers.UserRegisterPostHandler())
	g.POST("/sellerLogin", controllers.SellerLoginPostHandler())
	g.POST("/userLogin", controllers.UserLoginPostHandler())
	g.POST("/getAllProducts", controllers.GetAllProducts())
	g.POST("/getProduct", controllers.GetProduct())
}

func PrivateRoutes(g *gin.RouterGroup) {
	g.POST("/profile", controllers.UserProfile())
	g.POST("/getUserInfo", controllers.GetUserInfo())
	g.POST("/changeUserPassword", controllers.ChangeUserPassword())

	g.POST("/addProduct", controllers.AddProduct())
	g.POST("/getProducts", controllers.GetSellerProducts())
	g.POST("/deleteProduct", controllers.DeleteProduct())
	g.POST("/editProduct", controllers.EditProduct())

	g.POST("/newUserAddress", controllers.NewUserAddress())
	g.POST("/getUserAddressById", controllers.GetUserAddressById())
	g.POST("/getUserAddresses", controllers.GetUserAddresses())

	g.POST("/addProductToCart", controllers.AddProductToCart())
	g.POST("/getCartProducts", controllers.GetCartProducts())
	g.POST("/getCartInfo", controllers.GetCartInfo())
	g.POST("/removeProductFromCart", controllers.RemoveProductFromCart())
	g.POST("/changeProductQty", controllers.ChangeProductQty())
	g.POST("/increaseProductQty", controllers.IncreaseProductQty())
	g.POST("/decreaseProductQty", controllers.DecreaseProductQty())
	g.POST("/addTotalToCart", controllers.AddTotalToCart())
	g.POST("/getTotalPrice", controllers.GetTotalPrice())

	g.POST("/clearCart", controllers.ClearCart())

	g.POST("/createPayment", controllers.CreatePayment())
	g.POST("/updatePaymentStatus", controllers.UpdatePaymentStatus())

	g.POST("/getProductsSellers", controllers.GetProductsSellers())

	g.POST("/insertOrder", controllers.InsertOrder())
	g.POST("/getAllOrders", controllers.GetAllOrders())
}
