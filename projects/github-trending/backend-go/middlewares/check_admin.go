package middlewares

import (
	"github-trending-api/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func IsAdmin() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			//handle logic of this middleware here
			if 1 == 1 {
				return c.JSON(
					http.StatusInternalServerError,
					models.Response{
						StatusCode: http.StatusInternalServerError,
						Message:    "Interal Server Error",
						Data:       nil,
					},
				)
			}

			return next(c)
		}
	}
}
