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
}

func PrivateRoutes(g *gin.RouterGroup) {
	g.POST("/profile", controllers.UserProfile())
	g.POST("/seller-dashboard", controllers.SellerDashboard())
	g.POST("/addProduct", controllers.AddProduct())
}
