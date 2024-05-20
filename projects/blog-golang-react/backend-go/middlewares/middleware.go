package middlewares

import (
	"blog-go-api/utils"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func IsAuthenticated(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	if _, err := utils.ParseJwt(cookie); err != nil {
		c.Status(http.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"error": "unauthenticated",
		})
	}

	return c.Next()
}
