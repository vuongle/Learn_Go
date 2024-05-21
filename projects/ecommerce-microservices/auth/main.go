package main

import (
	"ecommerce-microservices/auth/app"
	"ecommerce-microservices/auth/database"
	"ecommerce-microservices/auth/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	app := app.NewApplication(
		database.ProductData(database.Client, "Products"),
		database.UserData(database.Client, "Users"),
	)

	ginInstance := gin.New()
	routes.AddRoutes(ginInstance)
	ginInstance.Use(gin.Logger())
	//ginInstance.Use(middleware.Authentication())

	log.Fatal(ginInstance.Run(":" + port))
}
