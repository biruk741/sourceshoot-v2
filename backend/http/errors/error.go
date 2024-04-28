package errors

import (
	"errors"
	"net/http"
)

type ErrorType string

const (
	NotFoundError      ErrorType = "NOT_FOUND"
	InternalError      ErrorType = "INTERNAL"
	ValidationError    ErrorType = "VALIDATION"
	BadRequestError    ErrorType = "BAD_REQUEST"
	AuthorizationError ErrorType = "AUTHORIZATION"
	// ... add other error types as needed
)

type Error struct {
	Type    ErrorType
	Message string
	Err     error
}

func (c Error) Error() string {
	if c.Err != nil {
		return c.Message + ": " + c.Err.Error()
	}
	return c.Message
}

func (c Error) Unwrap() error {
	return c.Err
}

func New(errorType ErrorType, msg string) error {
	return Error{
		Type:    errorType,
		Message: msg,
	}
}

func Wrap(errorType ErrorType, msg string, err error) error {
	return Error{
		Type:    errorType,
		Message: msg,
		Err:     err,
	}
}

func GetHTTPStatusCode(err error) int {
	var customErr Error
	if ok := As(err, &customErr); ok {
		switch customErr.Type {
		case NotFoundError:
			return http.StatusNotFound
		case ValidationError:
		case BadRequestError:
			return http.StatusBadRequest
		case AuthorizationError:
			return http.StatusUnauthorized
		case InternalError:
		default:
			return http.StatusInternalServerError
		}
	}
	return http.StatusInternalServerError
}

func Is(err error, errorType ErrorType) bool {
	var customErr Error
	if ok := As(err, &customErr); ok {
		return customErr.Type == errorType
	}
	return false
}

// As helps in type assertion of custom error
func As(err error, target interface{}) bool {
	return errors.As(err, target)
}
