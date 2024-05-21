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

	app.Post("apis/posts/create", controllers.CreatePost)
	app.Put("apis/posts/:id", controllers.UpdatePost)
	app.Delete("apis/posts/:id", controllers.DeletePost)

	app.Get("apis/posts", controllers.GetAllPosts)
	app.Get("apis/posts/:id", controllers.GetPostDetail)
	app.Get("apis/posts-by-user", controllers.UniquePost)

	app.Post("apis/upload-post-image", controllers.Upload)
	app.Static("apis/storage/", "./storage")
}
