package repo_impl

import (
	"context"
	"database/sql"
	"github-trending-api/db"
	app_errors "github-trending-api/errors"
	"github-trending-api/logger"
	"github-trending-api/models"
	"github-trending-api/models/req"
	"github-trending-api/repositories"
	"time"
)

type UserRepoImpl struct {
	sql *db.Sql
}

func NewUserRepo(sql *db.Sql) repositories.UserRepo {
	return &UserRepoImpl{sql: sql}
}

func (u *UserRepoImpl) SignUp(ctx context.Context, user models.User) (models.User, error) {
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	statement := `
		insert into users(user_id,full_name,email,password,role,created_at,updated_at)
		values(:user_id,:full_name,:email,:password,:role,:created_at,:updated_at)
	`
	_, err := u.sql.Db.NamedExecContext(ctx, statement, user)
	if err != nil {
		logger.Error(err.Error())

		return user, app_errors.ErrUserSignUpFail
	}

	return user, nil
}

func (u *UserRepoImpl) SignIn(ctx context.Context, loginReq req.ReqUserSignIn) (models.User, error) {
	var user = models.User{}

	if err := u.sql.Db.GetContext(ctx, &user, "select * from users where email = ?", loginReq.Email); err != nil {
		logger.Error(err.Error())
		if err == sql.ErrNoRows {
			return user, app_errors.ErrUserNotFound
		}

		return user, err
	}

	return user, nil
}

func (u *UserRepoImpl) SelectUserById(
	ctx context.Context,
	userId string,
) (models.User, error) {
	var user models.User
	err := u.sql.Db.GetContext(ctx, &user,
		"select * from users where user_id = ?", userId)
	if err != nil {
		logger.Error(err.Error())
		if err == sql.ErrNoRows {
			return user, app_errors.ErrUserNotFound
		}

		return user, err
	}

	return user, nil
}

func (u *UserRepoImpl) UpdateUser(ctx context.Context, user models.User) (models.User, error) {
	statement := `
	update users
	set 
	full_name = (case when length(:full_name) = 0 then full_name else :full_name end),
	email = (case when length(:email) = 0 then email else :email end),
	updated_at = coalesce(:updated_at, updated_at)
	where user_id = :user_id
	`
	user.UpdatedAt = time.Now()

	result, err := u.sql.Db.NamedExecContext(ctx, statement, user)
	if err != nil {
		logger.Error(err.Error())
		return user, err
	}

	count, _ := result.RowsAffected()
	if count == 0 {
		return user, app_errors.ErrUserUpdatedFail
	}

	return user, nil
}
