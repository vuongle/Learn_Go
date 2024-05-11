package common

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

type AppError struct {
	StatusCode int    `json:"status_code"`
	RootErr    error  `json:"-"`
	Message    string `json:"message"`
	Log        string `json:"log"`
	Key        string `json:"error_key"`
}

func NewAppErrorResponse(
	statusCode int,
	root error,
	msg, log, key string,
) *AppError {
	return &AppError{
		StatusCode: statusCode,
		RootErr:    root,
		Message:    msg,
		Log:        log,
		Key:        key,
	}
}

func NewBadReqErrorResponse(root error, msg, log, key string) *AppError {
	return &AppError{
		StatusCode: http.StatusBadRequest,
		RootErr:    root,
		Message:    msg,
		Log:        log,
		Key:        key,
	}
}

func NewCustomError(root error, msg, key string) *AppError {
	if root != nil {
		return NewBadReqErrorResponse(root, msg, root.Error(), key)
	}

	return NewBadReqErrorResponse(errors.New(msg), msg, msg, key)
}

func ErrDB(err error) *AppError {
	return NewAppErrorResponse(
		http.StatusInternalServerError,
		err,
		"something went wrong with DB",
		err.Error(),
		"DB_ERROR",
	)
}

func ErrInvalidRequest(err error) *AppError {
	return NewBadReqErrorResponse(
		err,
		"invalid request",
		err.Error(),
		"ErrInvalidRequest",
	)
}

func ErrInternal(err error) *AppError {
	return NewAppErrorResponse(
		http.StatusInternalServerError,
		err,
		"something went wrong in the server",
		err.Error(),
		"ErrInternal",
	)
}

func ErrCannotListEntity(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("Can not list %s", strings.ToLower(entity)),
		fmt.Sprintf("CanNotList%s", entity),
	)
}

func ErrCannotDelEntity(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("Can not delete %s", strings.ToLower(entity)),
		fmt.Sprintf("CanNotDelete%s", entity),
	)
}

func ErrCannotUpdateEntity(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("Can not update %s", strings.ToLower(entity)),
		fmt.Sprintf("CanNotUpdate%s", entity),
	)
}

func ErrCannotGetEntity(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("Can not get %s", strings.ToLower(entity)),
		fmt.Sprintf("CanNotGet%s", entity),
	)
}

func ErrCannotCreateEntity(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("Can not create %s", strings.ToLower(entity)),
		fmt.Sprintf("CanNotCreate%s", entity),
	)
}

func ErrEntityDeleted(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("%s has been deleted", strings.ToLower(entity)),
		fmt.Sprintf("Err%sDeleted", entity),
	)
}

func ErrEntityExisted(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("%s already existed", strings.ToLower(entity)),
		fmt.Sprintf("Err%sExisted", entity),
	)
}

func ErrEntityNotFound(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("%s not found", strings.ToLower(entity)),
		fmt.Sprintf("Err%sNotFound", entity),
	)
}

func ErrNoPermisson(err error) *AppError {
	return NewCustomError(
		err,
		"you have no permisson",
		"ErrNoPermission",
	)
}

var RecordNotFound = errors.New("record not found")

func (e *AppError) RootError() error {
	if err, ok := e.RootErr.(*AppError); ok {
		return err.RootError()
	}

	return e.RootErr
}

// implement the following method to make "AppError" struct be an error struct
func (e *AppError) Error() string {
	return e.RootError().Error()
}
