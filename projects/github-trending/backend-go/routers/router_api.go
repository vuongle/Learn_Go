package routers

import (
	"github-trending-api/handlers"
	"github-trending-api/middlewares"

	"github.com/labstack/echo/v4"
)

type API struct {
	Echo            *echo.Echo
	UserHandler     handlers.UserHandler
	RepoHandler     handlers.RepoHandler
	GracefulHandler handlers.GracefulHandler
}

func (api *API) SetupRouter() {

	// For authentication apis
	api.Echo.POST("/user/sign-in", api.UserHandler.HandleSignIn, middlewares.IsAdmin())
	api.Echo.POST("/user/sign-up", api.UserHandler.HandleSignUp)

	// For user apis
	user := api.Echo.Group("/user", middlewares.JwtMiddleware())
	user.GET("/profile", api.UserHandler.GetProfile)
	user.PUT("/edit-profile", api.UserHandler.UpdateProfile)

	github := api.Echo.Group("/github", middlewares.JwtMiddleware())
	github.GET("/trending", api.RepoHandler.RepoTrending)

	// For testing graceful shutdown
	api.Echo.GET("/gracefull", api.GracefulHandler.HandleShutdown)
}
