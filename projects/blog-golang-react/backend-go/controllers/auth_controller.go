package controllers

import (
	"blog-go-api/models"
	"blog-go-api/repositories"
	"blog-go-api/utils"
	"net/http"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	var data map[string]interface{}

	if err := c.BodyParser(&data); err != nil {
		c.Status(http.StatusBadRequest)
		return c.JSON(
			fiber.Map{
				"error": "Invalid body",
			},
		)
	}

	// validate body
	if len(data["password"].(string)) < 6 {
		c.Status(http.StatusBadRequest)
		return c.JSON(
			fiber.Map{
				"error": "Password must be greater than 6",
			},
		)
	}

	// Check email existed or not
	if existed := repositories.IsEmailExisted(data["email"].(string)); existed {
		c.Status(http.StatusBadRequest)
		return c.JSON(
			fiber.Map{
				"error": "Email already existed",
			},
		)
	}

	user := models.User{
		FirstName: data["first_name"].(string),
		LastName:  data["last_name"].(string),
		Email:     data["email"].(string),
		Phone:     data["phone"].(string),
	}
	user.HashPassword(data["password"].(string))
	err := repositories.RegisterUser(&user)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return c.JSON(
			fiber.Map{
				"error": err.Error(),
			},
		)
	}

	// response
	c.Status(http.StatusOK)
	return c.JSON(
		fiber.Map{
			"user": user,
		},
	)
}

func Login(c *fiber.Ctx) error {
	var data map[string]string

	// Parse the body
	if err := c.BodyParser(&data); err != nil {
		c.Status(http.StatusBadRequest)
		return c.JSON(
			fiber.Map{
				"error": "Invalid body",
			},
		)
	}

	// Check email existed or not
	user, err := repositories.FindUserByEmail(data["email"])
	if err != nil {
		c.Status(http.StatusNotFound)
		return c.JSON(
			fiber.Map{
				"error": "Email not existed",
			},
		)
	}

	// Check password
	if err := user.CompareWithHashedPassword(data["password"]); err != nil {
		c.Status(http.StatusBadRequest)
		return c.JSON(
			fiber.Map{
				"error": "Password is incorrect",
			},
		)
	}

	// Generate jwt and return to client
	token, err := utils.GenerateJwt(strconv.Itoa(int(user.Id)))
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return nil
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)
	c.Status(http.StatusOK)
	return c.JSON(
		fiber.Map{
			"user": user,
		},
	)
}
