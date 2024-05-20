package controllers

import (
	"blog-go-api/models"
	"blog-go-api/repositories"
	"math"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func CreatePost(c *fiber.Ctx) error {

	// Parse the body
	var blogpost models.Blog
	if err := c.BodyParser(&blogpost); err != nil {
		c.Status(http.StatusBadRequest)
		return c.JSON(
			fiber.Map{
				"message": "Invalid body",
				"error":   err,
			},
		)
	}

	// Create post
	if err := repositories.CreatePost(&blogpost); err != nil {
		c.Status(http.StatusNotAcceptable)
		return c.JSON(
			fiber.Map{
				"message": "Can not create the post",
				"error":   err,
			},
		)
	}

	c.Status(http.StatusOK)
	return c.JSON(
		fiber.Map{
			"message": "Create the post successfully",
			"post":    blogpost,
		},
	)
}

func GetAllPosts(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit := 5
	offset := (page - 1) * limit
	total, blogs := repositories.GetAllPosts(&page, &limit, &offset)
	c.Status(http.StatusOK)
	return c.JSON(
		fiber.Map{
			"message": "",
			"data": fiber.Map{
				"blogs":     blogs,
				"total":     total,
				"last_page": math.Ceil(float64(int(total) / limit)),
			},
		},
	)
}
