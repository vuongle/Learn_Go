package main

import (
	"blog-go-api/database"
	"blog-go-api/routes"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment
	err := godotenv.Load()
	if err != nil {
		log.Fatal("can not load environment variables")
	}

	// Connect to db
	database.Connect()
	log.Println("Connected to DB")

	// Initialize a new Fiber app
	app := fiber.New()
	routes.SetupRoutes(app)

	// Start the server on port 3000
	port := os.Getenv("APP_PORT")
	log.Printf("App running on port %v", port)
	log.Fatal(app.Listen(":" + port))
}
