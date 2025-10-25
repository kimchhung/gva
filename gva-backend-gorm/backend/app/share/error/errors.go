package apperror

import (
	"errors"
	"fmt"
	"net/http"

	"gorm.io/gorm"
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

func newPublicError(httpCode int, errorCode int, message string) *Error {
	return &Error{
		HttpCode:  httpCode,
		ErrorCode: errorCode,
		Message:   message,
		isPublic:  true,
	}
}

func ErrorByCode(code int) (*Error, error) {
	if err, exists := errorByCode[code]; exists {
		return err, nil
	}

	return nil, fmt.Errorf("get by code %d not exist", code)
}

func DBError(err error) error {
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		return ErrRecordNotFound
	case errors.Is(err, gorm.ErrDuplicatedKey):
		return ErrDuplicatedRecord
	}

	return err
}

func DBErrorIs(err error, gormErr error, publicErr *Error) error {
	if errors.Is(err, gormErr) {
		return publicErr
	}

	return err
}

var (
	// map to check uniqness for error code
	errorByCode = map[int]*Error{}

	// general
	ErrUnknownError          = newError(http.StatusInternalServerError, -1, "unknown_error")
	ErrValidationError       = newPublicError(http.StatusBadRequest, -2, "validation_error")
	ErrUnauthorized          = newPublicError(http.StatusForbidden, -3, "unauthorized")
	ErrLoginError            = newPublicError(http.StatusBadRequest, -4, "login_error")
	ErrUnauthenticated       = newPublicError(http.StatusUnauthorized, -5, "unauthenticated")
	ErrForbidden             = newPublicError(http.StatusForbidden, -6, "access_denied")
	ErrRecordNotFound        = newPublicError(http.StatusNotFound, -7, "record_not_found")
	ErrInvalidTOTP           = newPublicError(http.StatusBadRequest, -8, "invalid_otp_code")
	ErrNotFound              = newPublicError(http.StatusNotFound, -9, "not_found")
	ErrInvalidToken          = newPublicError(http.StatusForbidden, -10, "invalid_token")
	ErrBadRequest            = newPublicError(http.StatusBadRequest, -11, "bad_request")
	ErrTokenExpired          = newPublicError(http.StatusForbidden, -12, "token_expired")
	ErrContentRemoved        = newPublicError(http.StatusForbidden, -13, "content_has_been_removed")
	ErrStatusDisable         = newPublicError(http.StatusForbidden, -14, "status_has_been_disabled")
	ErrInvalidCredentials    = newPublicError(http.StatusForbidden, -15, "invalid_credentials")
	ErrTooManyRetries        = newPublicError(http.StatusForbidden, -16, "too_many_retries")
	ErrAdminWhitelistInvalid = newPublicError(http.StatusForbidden, -17, "admin_whitelist_invalid")
	ErrInvalidIP             = newPublicError(http.StatusForbidden, -18, "invalid_ip")
	ErrDuplicatedRecord      = newPublicError(http.StatusConflict, -19, "duplicated_record")
	ErrKeyExists             = newPublicError(http.StatusConflict, -20, "key_already_exists")

	// image upload errors (300 - 399)
	ErrUnsupportedFileFormat = newPublicError(http.StatusBadRequest, -300, "unsupported_file_format")
	ErrWhileUploading        = newPublicError(http.StatusInternalServerError, -301, "error_while_uploading")
	ErrImageTooLarge         = newPublicError(http.StatusBadRequest, -302, "image_too_large")
	ErrImageInvalid          = newPublicError(http.StatusBadRequest, -303, "image_invalid")
	ErrMIMETypeInvalid       = newPublicError(http.StatusBadRequest, -304, "mime_type_invalid")
)
