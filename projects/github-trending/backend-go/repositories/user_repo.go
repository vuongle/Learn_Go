package repositories

import (
	"context"
	"github-trending-api/models"
	"github-trending-api/models/req"
)

type UserRepo interface {
	SignUp(ctx context.Context, user models.User) (models.User, error)
	SignIn(ctx context.Context, loginReq req.ReqUserSignIn) (models.User, error)
	SelectUserById(ctx context.Context, userId string) (models.User, error)
	UpdateUser(ctx context.Context, user models.User) (models.User, error)
}
