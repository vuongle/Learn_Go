package handlers

import (
	"github-trending-api/models"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type GracefulHandler struct {
}

func (r GracefulHandler) HandleShutdown(c echo.Context) error {
	time.Sleep(60 * time.Second)
	return c.JSON(
		http.StatusOK,
		models.Response{
			StatusCode: http.StatusOK,
			Message:    "DONE",
			Data:       "",
		})
}
