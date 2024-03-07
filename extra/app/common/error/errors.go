package app_err

import (
	"fmt"
	"net/http"
)

/*
	standard http status code

Informational responses (100 – 199)
Successful responses (200 – 299)
Redirection messages (300 – 399)
Client error responses (400 – 499)
Server error responses (500 – 599)
*/
func newError(httpCode int, errorCode int, message string) *Error {
	if _, exists := errorByCode[errorCode]; exists {
		// If the error code already exists, return an error.
		panic(fmt.Sprintf("[ErrorByCode]: error code %d is not unique", errorCode))
	}

	err := &Error{
		HttpCode:  httpCode,
		ErrorCode: errorCode,
		Message:   message,
	}

	errorByCode[errorCode] = err
	return err
}

func ErrorByCode(code int) (*Error, error) {
	if err, exists := errorByCode[code]; exists {
		return err, nil
	}

	return nil, fmt.Errorf("get by code %d not exist", code)
}

var (
	// map to check uniqness for error code
	errorByCode = map[int]*Error{}

	// general
	ErrUnknownError    = newError(http.StatusInternalServerError, -1, "unknown_error")
	ErrValidationError = newError(http.StatusBadRequest, -2, "validation_error")
	ErrUnauthorized    = newError(http.StatusForbidden, -3, "unauthorized")
	ErrBadRequest      = newError(http.StatusBadRequest, -4, "bad_request")
	ErrUnauthenticated = newError(http.StatusUnauthorized, -5, "unauthenticated")
	ErrForbidden       = newError(http.StatusForbidden, -6, "access_denied")
	ErrInvalidTOTP     = newError(http.StatusBadRequest, -8, "invalid otp code")
	ErrNotFound        = newError(http.StatusNotFound, -9, "not_found")
	ErrInvalidToken    = newError(http.StatusForbidden, -10, "invalid_token")
	ErrTooManyRetries  = newError(http.StatusForbidden, -11, "too_many_retries")

	// users
	ErrPasswordValidationError = newError(http.StatusBadRequest, -501, "wrong_password")
	ErrUserNotFound            = newError(http.StatusNotFound, -502, "user_not_found")
	ErrUsernameExists          = newError(http.StatusConflict, -505, "username_already_exists")
)
