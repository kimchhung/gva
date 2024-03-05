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
	ErrUnknownError    = newError(http.StatusInternalServerError, -1, "unknown error")
	ErrValidationError = newError(http.StatusBadRequest, -2, "validation error")
	ErrUnauthorized    = newError(http.StatusForbidden, -3, "unauthorized")
	ErrLoginError      = newError(http.StatusBadRequest, -4, "login error")
	ErrUnauthenticated = newError(http.StatusUnauthorized, -5, "unauthenticated")
	ErrForbidden       = newError(http.StatusForbidden, -6, "access denied")
	ErrRecordNotFound  = newError(http.StatusNotFound, -7, "record not found")
	ErrInvalidTOTP     = newError(http.StatusForbidden, -8, "invalid otp code")
	ErrNotFound        = newError(http.StatusNotFound, -9, "not found")
	ErrInvalidToken    = newError(http.StatusForbidden, -10, "invalid token")
	ErrTooManyRetries  = newError(http.StatusForbidden, -11, "too many retries")

	// users
	ErrPasswordValidationError  = newError(http.StatusBadRequest, -501, "wrong password")
	ErrUserNotFound             = newError(http.StatusNotFound, -502, "user not found")
	ErrPhoneNumberAlreadyExists = newError(http.StatusConflict, -503, "phone number already exists")
	ErrWrongVerifyCode          = newError(http.StatusBadRequest, -504, "wrong verify code")
	ErrUsernameExists           = newError(http.StatusConflict, -505, "username already exists")
)
