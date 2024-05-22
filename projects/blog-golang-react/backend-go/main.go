package main

import (
	"blog-go-api/database"
	"blog-go-api/routes"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
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
	app.Use(cors.New(cors.Config{
		AllowHeaders:     "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin",
		AllowOrigins:     "http://localhost:3000,http://127.0.0.1:3000",
		AllowCredentials: true,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))
	routes.SetupRoutes(app)

	// Start the server on port 3000
	port := os.Getenv("APP_PORT")
	log.Printf("App running on port %v", port)
	log.Fatal(app.Listen(":" + port))
}
