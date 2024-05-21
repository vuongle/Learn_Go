package controllers

import (
	"blog-go-api/models"
	"blog-go-api/repositories"
	"blog-go-api/utils"
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
			"data": fiber.Map{
				"blog": blogpost,
			},
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

func GetPostDetail(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	blog := repositories.GetPostById(&id)
	c.Status(http.StatusOK)
	return c.JSON(
		fiber.Map{
			"message": "",
			"data": fiber.Map{
				"blog": blog,
			},
		},
	)
}

func UpdatePost(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	blog := models.Blog{
		Id: uint(id),
	}
	if err := c.BodyParser(&blog); err != nil {
		c.Status(http.StatusBadRequest)
		return c.JSON(
			fiber.Map{
				"message": "Invalid body",
				"error":   err,
			},
		)
	}

	if err := repositories.UpdatePost(&blog); err != nil {
		c.Status(http.StatusInternalServerError)
		return c.JSON(
			fiber.Map{
				"message": "Can not update the post",
				"error":   err,
			},
		)
	}

	c.Status(http.StatusOK)
	return c.JSON(
		fiber.Map{
			"message": "",
			"data": fiber.Map{
				"blog": blog,
			},
		},
	)
}

func UniquePost(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	id, _ := utils.ParseJwt(cookie)

	blogs := repositories.GetPostsByUserId(&id)
	c.Status(http.StatusOK)
	return c.JSON(
		fiber.Map{
			"message": "",
			"data": fiber.Map{
				"blogs": blogs,
			},
		},
	)
}

func DeletePost(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	blog := models.Blog{
		Id: uint(id),
	}
	if err := repositories.DeletePost(&blog); err != nil {
		c.Status(http.StatusInternalServerError)
		return c.JSON(
			fiber.Map{
				"message": "Can not delete the post",
				"error":   err,
			},
		)
	}

	c.Status(http.StatusOK)
	return c.JSON(
		fiber.Map{
			"message": "",
			"data":    "",
		},
	)

}
