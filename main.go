package main

import (
	// "net/http"

	"e-comm/authService/repositories/mongodb"
	"e-comm/authService/repositories/postgresql"
	"e-comm/authService/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// var client http.Client

func main() {
	go postgresql.Connect()
	go mongodb.Connect()

	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3001/seller-register", "http://localhost:3001"}
	router.Use(cors.New(config))

	public := router.Group("/")
	routes.PublicRoutes(public)

	private := router.Group("/")
	// private.Use(middleware.AuthRequired)
	routes.PrivateRoutes(private)

	if err := router.Run(":3002"); err != nil {
		panic(err)
	}
}
