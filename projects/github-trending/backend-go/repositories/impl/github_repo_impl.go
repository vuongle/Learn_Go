package repo_impl

import (
	"context"
	"database/sql"
	"github-trending-api/db"
	app_errors "github-trending-api/errors"
	"github-trending-api/logger"
	"github-trending-api/models"
	"github-trending-api/repositories"
	"time"

	"github.com/lib/pq"
)

type GithubRepoImpl struct {
	sql *db.Sql
}

func NewGithubRepo(sql *db.Sql) repositories.GithubRepo {
	return &GithubRepoImpl{
		sql: sql,
	}
}

func (g GithubRepoImpl) SelectRepoByName(ctx context.Context, name string) (models.Github, error) {
	var repo = models.Github{}
	err := g.sql.Db.GetContext(ctx, &repo,
		`SELECT * FROM repos WHERE name = ?`, name)

	if err != nil {
		if err == sql.ErrNoRows {
			return repo, app_errors.ErrRepoNotFound
		}
		logger.Error(err.Error())
		return repo, err
	}
	return repo, nil
}

func (g GithubRepoImpl) SaveRepo(context context.Context, repo models.Github) (models.Github, error) {
	// name, description, url, color, lang, fork, stars, stars_today, build_by, created_at, updated_at
	statement := `INSERT INTO repos(
					name, description, url, color, lang, fork, stars, 
 			        stars_today, build_by, created_at, updated_at) 
          		  VALUES(
					:name,:description, :url, :color, :lang, :fork, :stars, 
					:stars_today, :build_by, :created_at, :updated_at
				  )`

	repo.CreatedAt = time.Now()
	repo.UpdatedAt = time.Now()

	_, err := g.sql.Db.NamedExecContext(context, statement, repo)
	if err != nil {
		logger.Error(err.Error())
		return repo, app_errors.ErrRepoInsertFail
	}

	return repo, nil
}

func (g GithubRepoImpl) SelectRepos(context context.Context, userId string, limit int) ([]models.Github, error) {
	var repos []models.Github
	err := g.sql.Db.SelectContext(context, &repos,
		`
			SELECT 
				repos.name, repos.description, repos.url, repos.color, repos.lang, 
				repos.fork, repos.stars, repos.stars_today, repos.build_by, repos.updated_at, 
				FALSE as bookmarked
			FROM repos 
			WHERE repos.name IS NOT NULL 
			ORDER BY updated_at ASC LIMIT ?
		`, limit)

	if err != nil {
		logger.Error(err.Error())
		return repos, err
	}

	return repos, nil
}

func (g GithubRepoImpl) UpdateRepo(context context.Context, repo models.Github) (models.Github, error) {
	// name, description, url, color, lang, fork, stars, stars_today, build_by, created_at, updated_at
	sqlStatement := `
		UPDATE repos
		SET 
			stars  = :stars,
			fork = :fork,
			stars_today = :stars_today,
			build_by = :build_by,
			updated_at = :updated_at
		WHERE name = :name
	`
	result, err := g.sql.Db.NamedExecContext(context, sqlStatement, repo)
	if err != nil {
		logger.Error(err.Error())
		return repo, err
	}

	count, err := result.RowsAffected()
	if err != nil {
		logger.Error(err.Error())
		return repo, app_errors.ErrRepoUpdatedFail
	}
	if count == 0 {
		return repo, app_errors.ErrRepoUpdatedFail
	}

	return repo, nil
}

func (g GithubRepoImpl) SelectAllBookmarks(context context.Context, userId string) ([]models.Github, error) {
	repos := []models.Github{}
	err := g.sql.Db.SelectContext(context, &repos,
		`SELECT 
					repos.name, repos.description, repos.url, 
					repos.color, repos.lang, repos.fork, repos.stars, 
					repos.stars_today, repos.build_by, true as bookmarked
				FROM bookmarks 
				INNER JOIN repos
				ON bookmarks.user_id = ? AND repos.name = bookmarks.repo_name`, userId)

	if err != nil {
		if err == sql.ErrNoRows {
			return repos, app_errors.ErrBookmarkNotFound
		}
		logger.Error(err.Error())
		return repos, err
	}
	return repos, nil
}

func (g GithubRepoImpl) Bookmark(context context.Context, bid, nameRepo, userId string) error {
	statement := `INSERT INTO bookmarks(
					bid, user_id, repo_name, created_at, updated_at) 
          		  VALUES(?, ?, ?, ?, ?)`

	now := time.Now()
	_, err := g.sql.Db.ExecContext(
		context, statement, bid, userId,
		nameRepo, now, now)

	if err != nil {
		if err, ok := err.(*pq.Error); ok {
			if err.Code.Name() == "unique_violation" {
				return app_errors.ErrBookmarkConflict
			}
		}
		logger.Error(err.Error())
		return app_errors.ErrBookmarkFail
	}

	return nil
}

func (g GithubRepoImpl) DelBookmark(context context.Context, nameRepo, userId string) error {
	result := g.sql.Db.MustExecContext(
		context,
		"DELETE FROM bookmarks WHERE repo_name = ? AND user_id = ?",
		nameRepo, userId)

	_, err := result.RowsAffected()
	if err != nil {
		logger.Error(err.Error())
		return app_errors.ErrDelBookmarkFail
	}

	return nil
}
