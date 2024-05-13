package handlers

import (
	"github-trending-api/constants"
	"github-trending-api/helper"
	"github-trending-api/models"
	"github-trending-api/repositories"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type RepoHandler struct {
	GithubRepo repositories.GithubRepo
}

func (r RepoHandler) RepoTrending(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*models.JwtCustomClaims)

	// Get data from cache if it is
	value, err := helper.GetCache(c.Request().Context(), constants.REDIS_TRENDING_REPO_KEY)
	// If there is any error from cache -> get data from db -> return this data
	// else, return data from cache
	if err != nil {
		repos, _ := r.GithubRepo.SelectRepos(c.Request().Context(), claims.UserId, 25)
		for i, repo := range repos {
			repos[i].Contributors = strings.Split(repo.BuildBy, ",")
		}

		return c.JSON(http.StatusOK, models.Response{
			StatusCode: http.StatusOK,
			Message:    "Data from DB",
			Data:       repos,
		})
	} else {
		repos, _ := r.GithubRepo.SelectRepos(c.Request().Context(), claims.UserId, 25)
		for i, repo := range repos {
			repos[i].Contributors = strings.Split(repo.BuildBy, ",")
		}

		return c.JSON(http.StatusOK, models.Response{
			StatusCode: http.StatusOK,
			Message:    "Data from CACHE",
			Data:       value,
		})
	}

}
