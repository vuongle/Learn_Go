package app_errors

import "errors"

var (
	ErrUserConflict    = errors.New("user already existed")
	ErrUserSignUpFail  = errors.New("can not sign up")
	ErrUserNotFound    = errors.New("user not found")
	ErrUserUpdatedFail = errors.New("user was updated fail")

	ErrRepoNotFound    = errors.New("repo not found")
	ErrRepoConflict    = errors.New("repo already existed")
	ErrRepoInsertFail  = errors.New("repo was inserted fail")
	ErrRepoUpdatedFail = errors.New("repo was updated fail")

	ErrBookmarkNotFound = errors.New("bookmark not found")
	ErrBookmarkConflict = errors.New("bookmark already existed")
	ErrBookmarkFail     = errors.New("bookmark was updated fail")
	ErrDelBookmarkFail  = errors.New("bookmark was deleted fail")
)
