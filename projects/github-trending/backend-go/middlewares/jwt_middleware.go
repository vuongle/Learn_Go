package middlewares

import (
	"github-trending-api/models"
	"github-trending-api/security"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func JwtMiddleware() echo.MiddlewareFunc {
	config := echojwt.Config{
		SigningKey: []byte(security.SECRET_KEY),
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(models.JwtCustomClaims)
		},
	}

	return echojwt.WithConfig(config)
}
