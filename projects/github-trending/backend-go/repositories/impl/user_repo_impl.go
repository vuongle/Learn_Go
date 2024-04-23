package repo_impl

import (
	"context"
	"database/sql"
	"fmt"
	"github-trending-api/db"
	app_errors "github-trending-api/errors"
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
	fmt.Sprintln(">>>>>> NewUserRepo")
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	statement := `
		insert into users(user_id,full_name,email,password,role,created_at,updated_at)
		values(:user_id,:full_name,:email,:password,:role,:created_at,:updated_at)
	`
	_, err := u.sql.Db.NamedExecContext(ctx, statement, user)
	if err != nil {
		fmt.Sprintln(err.Error())

		return user, app_errors.UserSignUpFail
	}

	return user, nil
}

func (u *UserRepoImpl) SignIn(ctx context.Context, loginReq req.ReqUserSignIn) (models.User, error) {
	var user = models.User{}

	if err := u.sql.Db.GetContext(ctx, &user, "select * from users where email = ?", loginReq.Email); err != nil {
		fmt.Sprintln(err.Error())
		if err == sql.ErrNoRows {
			return user, app_errors.UserNotFound
		}

		return user, err
	}

	return user, nil
}
