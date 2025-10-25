package coreerror

import "net/http"

var (
	ErrUnknownError       = newError(http.StatusInternalServerError, -1, "unknown_error")
	ErrValidationError    = NewPublicError(http.StatusBadRequest, -2, "validation_error")
	ErrUnauthorized       = NewPublicError(http.StatusForbidden, -3, "unauthorized")
	ErrLoginError         = NewPublicError(http.StatusBadRequest, -4, "login_error")
	ErrUnauthenticated    = NewPublicError(http.StatusUnauthorized, -5, "unauthenticated")
	ErrForbidden          = NewPublicError(http.StatusForbidden, -6, "access_denied")
	ErrRecordNotFound     = NewPublicError(http.StatusNotFound, -7, "record_not_found")
	ErrInvalidTOTP        = NewPublicError(http.StatusBadRequest, -8, "invalid_otp_code")
	ErrNotFound           = NewPublicError(http.StatusNotFound, -9, "not_found")
	ErrInvalidToken       = NewPublicError(http.StatusForbidden, -10, "invalid_token")
	ErrBadRequest         = NewPublicError(http.StatusBadRequest, -11, "bad_request")
	ErrTokenExpired       = NewPublicError(http.StatusForbidden, -12, "token_expired")
	ErrContentRemoved     = NewPublicError(http.StatusForbidden, -13, "content_has_been_removed")
	ErrStatusDisable      = NewPublicError(http.StatusForbidden, -14, "status_has_been_disabled")
	ErrInvalidCredentials = NewPublicError(http.StatusForbidden, -15, "invalid_credentials")
	ErrTooManyRetries     = NewPublicError(http.StatusForbidden, -16, "too_many_retries")
	ErrInvalidIP          = NewPublicError(http.StatusForbidden, -18, "invalid_ip")
	ErrDuplicatedRecord   = NewPublicError(http.StatusConflict, -19, "duplicated_record")
	ErrKeyExists          = NewPublicError(http.StatusConflict, -20, "key_already_exists")
)
