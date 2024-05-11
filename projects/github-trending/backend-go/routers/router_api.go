package routers

import (
	"github-trending-api/handlers"
	"github-trending-api/middlewares"

	"github.com/labstack/echo/v4"
)

type API struct {
	Echo        *echo.Echo
	UserHandler handlers.UserHandler
}

func (api *API) SetupRouter() {
	api.Echo.POST("/user/sign-in", api.UserHandler.HandleSignIn, middlewares.IsAdmin())
	api.Echo.POST("/user/sign-up", api.UserHandler.HandleSignUp)
}
