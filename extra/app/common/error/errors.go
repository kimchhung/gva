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
func New(httpCode int, errorCode int, message string) *Error {
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
	// map to check uniq error code
	errorByCode = map[int]*Error{}

	// general
	ErrUnknownError    = New(http.StatusInternalServerError, -1, "unknown error")
	ErrValidationError = New(http.StatusBadRequest, -2, "validation error")
	ErrUnauthorized    = New(http.StatusForbidden, -3, "unauthorized")
	ErrLoginError      = New(http.StatusBadRequest, -4, "login error")
	ErrUnauthenticated = New(http.StatusUnauthorized, -5, "unauthenticated")
	ErrForbidden       = New(http.StatusForbidden, -6, "access denied")
	ErrRecordNotFound  = New(http.StatusNotFound, -7, "record not found")
	ErrInvalidTOTP     = New(http.StatusForbidden, -8, "invalid otp code")
	ErrNotFound        = New(http.StatusNotFound, -9, "not found")
	ErrInvalidToken    = New(http.StatusForbidden, -10, "invalid token")
	ErrTooManyRetries  = New(http.StatusForbidden, -11, "too many retries")

	// users
	ErrPasswordValidationError  = New(http.StatusBadRequest, -501, "wrong password")
	ErrUserNotFound             = New(http.StatusNotFound, -502, "user not found")
	ErrPhoneNumberAlreadyExists = New(http.StatusConflict, -503, "phone number already exists")
	ErrWrongVerifyCode          = New(http.StatusBadRequest, -504, "wrong verify code")
	ErrUsernameExists           = New(http.StatusConflict, -507, "username already exists")
)
