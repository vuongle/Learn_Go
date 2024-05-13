package handlers

import (
	"fmt"
	app_errors "github-trending-api/errors"
	"github-trending-api/models"
	"github-trending-api/models/req"
	"github-trending-api/repositories"
	"github-trending-api/security"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	uuid "github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	UserRepo repositories.UserRepo
}

func (u *UserHandler) HandleSignUp(c echo.Context) error {

	// Binding req
	req := req.ReqUserSignUp{}
	if err := c.Bind(&req); err != nil {
		fmt.Println(err.Error())
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
	if err := c.Validate(req); err != nil {
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

	// Generate token and delete pass
	token, _ := security.GenerateToken(result)
	result.Token = token

	return c.JSON(
		http.StatusOK,
		models.Response{
			StatusCode: http.StatusOK,
			Message:    "",
			Data:       result,
		})
}

func (u *UserHandler) HandleSignIn(c echo.Context) error {

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
	if err := c.Validate(req); err != nil {
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

	// Generate token and delete pass
	token, _ := security.GenerateToken(result)
	result.Token = token

	return c.JSON(
		http.StatusOK,
		models.Response{
			StatusCode: http.StatusOK,
			Message:    "",
			Data:       result,
		})
}

func (u *UserHandler) GetProfile(c echo.Context) error {
	tokenData := c.Get("user").(*jwt.Token)
	claims := tokenData.Claims.(*models.JwtCustomClaims)

	user, err := u.UserRepo.SelectUserById(c.Request().Context(), claims.UserId)
	if err != nil {
		if err == app_errors.ErrUserNotFound {
			c.JSON(
				http.StatusNotFound,
				models.Response{
					StatusCode: http.StatusNotFound,
					Message:    err.Error(),
					Data:       nil,
				})
		}

		// other error
		c.JSON(
			http.StatusInternalServerError,
			models.Response{
				StatusCode: http.StatusInternalServerError,
				Message:    err.Error(),
				Data:       nil,
			})
	}

	return c.JSON(
		http.StatusOK,
		models.Response{
			StatusCode: http.StatusOK,
			Message:    "",
			Data:       user,
		})
}

func (u *UserHandler) UpdateProfile(c echo.Context) error {

	// Bind request body to struct
	req := req.ReqUserUpdate{}
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
	if err := c.Validate(req); err != nil {
		return c.JSON(
			http.StatusBadRequest,
			models.Response{
				StatusCode: http.StatusBadRequest,
				Message:    err.Error(),
				Data:       nil,
			},
		)
	}

	// Get user id from token
	tokenData := c.Get("user").(*jwt.Token)
	claims := tokenData.Claims.(*models.JwtCustomClaims)
	user := models.User{
		UserId:   claims.UserId,
		FullName: req.Fullname,
		Email:    req.Email,
	}

	result, err := u.UserRepo.UpdateUser(c.Request().Context(), user)
	if err != nil {
		return c.JSON(
			http.StatusUnprocessableEntity,
			models.Response{
				StatusCode: http.StatusUnprocessableEntity,
				Message:    err.Error(),
				Data:       nil,
			},
		)
	}

	return c.JSON(
		http.StatusOK,
		models.Response{
			StatusCode: http.StatusOK,
			Message:    "",
			Data:       result,
		})
}
