package controllers

import (
	"math/rand"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

var letters = []rune("abcdefghijklmnopqrsuvwxyz")

func randomLetter(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}

func Upload(c *fiber.Ctx) error {
	form, err := c.MultipartForm()
	if err != nil {
		return err
	}

	files := form.File["image"]
	fileName := ""

	for _, file := range files {
		fileName = randomLetter(5) + "_" + file.Filename
		if err := c.SaveFile(file, "./storage/"+fileName); err != nil {
			return nil
		}
	}

	c.Status(http.StatusOK)
	return c.JSON(
		fiber.Map{
			"url": "http://localhost:3000/apis/storage/" + fileName,
		},
	)
}
