package middlewares

import (
	"github.com/labstack/echo/v4"
)

// This is an example of a custom middleware.
// A middleware is actual an middleware function, the middleware function has type "func(next HandlerFunc) HandlerFunc"
// Therefore, this middleware function must return "func(next echo.HandlerFunc) echo.HandlerFunc".
// Inside the HanlderFunc, because it has type "func(c Context) error". Therefore, it must return "func(c Context) error"
// Finally, there are 2 returns inside the MiddlewareFunc.
func IsAdmin() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			//handle logic of this middleware here
			// if 1 == 1 {
			// 	return c.JSON(
			// 		http.StatusInternalServerError,
			// 		models.Response{
			// 			StatusCode: http.StatusInternalServerError,
			// 			Message:    "Interal Server Error",
			// 			Data:       nil,
			// 		},
			// 	)
			// }

			return next(c)
		}
	}
}
