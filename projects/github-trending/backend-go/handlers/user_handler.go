package handlers

import (
	"fmt"
	"github-trending-api/models"
	"github-trending-api/models/req"
	"github-trending-api/repositories"
	"github-trending-api/security"
	"net/http"

	"github.com/go-playground/validator/v10"
	uuid "github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	UserRepo repositories.UserRepo
}

func (u *UserHandler) HandleSignUp(c echo.Context) error {
	fmt.Sprintln(">>>>>> HandleSignUp")

	// Binding req
	req := req.ReqUserSignUp{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(
			http.StatusBadRequest,
			models.Response{
				StatusCode: http.StatusBadRequest,
				Message:    err.Error(),
				Data:       nil,
			},
		)
	}

	// Validate req
	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		return c.JSON(
			http.StatusBadRequest,
			models.Response{
				StatusCode: http.StatusBadRequest,
				Message:    err.Error(),
				Data:       nil,
			},
		)
	}

	// Handler logic signup
	hash, _ := security.HashPassword(req.Password)
	role := models.MEMBER.String()
	userId, _ := uuid.NewUUID()
	user := models.User{
		UserId:   userId.String(),
		FullName: req.Fullname,
		Email:    req.Email,
		Password: hash,
		Role:     role,
	}

	result, err := u.UserRepo.SignUp(c.Request().Context(), user)
	if err != nil {
		return c.JSON(
			http.StatusConflict,
			models.Response{
				StatusCode: http.StatusConflict,
				Message:    err.Error(),
				Data:       nil,
			},
		)
	}

	result.Password = ""
	return c.JSON(
		http.StatusOK,
		models.Response{
			StatusCode: http.StatusOK,
			Message:    "",
			Data:       result,
		})
}

func (u *UserHandler) HandleSignIn(c echo.Context) error {
	fmt.Sprintln(">>>>>> HandleSignIn")

	// Binding req
	req := req.ReqUserSignIn{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(
			http.StatusBadRequest,
			models.Response{
				StatusCode: http.StatusBadRequest,
				Message:    err.Error(),
				Data:       nil,
			},
		)
	}

	// Validate req
	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		return c.JSON(
			http.StatusBadRequest,
			models.Response{
				StatusCode: http.StatusBadRequest,
				Message:    err.Error(),
				Data:       nil,
			},
		)
	}

	// Logic sign-in
	result, err := u.UserRepo.SignIn(c.Request().Context(), req)
	if err != nil {
		return c.JSON(
			http.StatusUnauthorized,
			models.Response{
				StatusCode: http.StatusUnauthorized,
				Message:    err.Error(),
				Data:       nil,
			},
		)
	}

	ok := security.CheckPasswordHash(req.Password, result.Password)
	if !ok {
		return c.JSON(
			http.StatusUnauthorized,
			models.Response{
				StatusCode: http.StatusUnauthorized,
				Message:    "Password is not correct",
				Data:       nil,
			},
		)
	}

	result.Password = ""
	return c.JSON(
		http.StatusOK,
		models.Response{
			StatusCode: http.StatusOK,
			Message:    "",
			Data:       result,
		})
}
