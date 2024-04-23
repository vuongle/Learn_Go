package app_errors

import "errors"

var (
	UserConflict   = errors.New("User already existed")
	UserSignUpFail = errors.New("Can not sign up.")
	UserNotFound   = errors.New("User not found")
)
