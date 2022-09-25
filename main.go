package main

import (
	"auth-and-db-service/dotEnv"
	"auth-and-db-service/repositories/mongodb"
	"auth-and-db-service/repositories/postgresql"
	"auth-and-db-service/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	go postgresql.Connect()
	go mongodb.Connect()

	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{dotEnv.GoDotEnvVariable("BFF_URL")}
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
