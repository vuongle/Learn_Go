package routes

import (
	"blog-go-api/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Post("/apis/auth/register", controllers.Register)
	app.Post("/apis/auth/login", controllers.Login)

	// Register authentication middleware to protect routes
	//app.Use(middlewares.IsAuthenticated)

	app.Post("apis/post/create", controllers.CreatePost)
}
