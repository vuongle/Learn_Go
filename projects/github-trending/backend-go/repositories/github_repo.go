package repositories

import (
	"context"
	"github-trending-api/models"
)

type GithubRepo interface {
	SaveRepo(ctx context.Context, user models.Github) (models.Github, error)
	SelectRepos(ctx context.Context, userId string, limit int) ([]models.Github, error)
	SelectRepoByName(ctx context.Context, name string) (models.Github, error)
	UpdateRepo(ctx context.Context, user models.Github) (models.Github, error)
}
