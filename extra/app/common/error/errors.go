package app_err

/*
	standard http status code

Informational responses (100 – 199)
Successful responses (200 – 299)
Redirection messages (300 – 399)
Client error responses (400 – 499)
Server error responses (500 – 599)
*/
var (
	// general
	ErrUnknownError    = &Error{HttpCode: 500, ErrorCode: -1, Message: "unknown error"}
	ErrValidationError = &Error{HttpCode: 400, ErrorCode: -2, Message: "validation error"}
	ErrUnauthorized    = &Error{HttpCode: 403, ErrorCode: -3, Message: "unauthorized"}
	ErrLoginError      = &Error{HttpCode: 400, ErrorCode: -4, Message: "login error"}
	ErrUnauthenticated = &Error{HttpCode: 401, ErrorCode: -5, Message: "unauthenticated"}
	ErrForbidden       = &Error{HttpCode: 403, ErrorCode: -6, Message: "access denied"}
	ErrRecordNotFound  = &Error{HttpCode: 404, ErrorCode: -7, Message: "record not found"}
	ErrInvalidTOTP     = &Error{HttpCode: 403, ErrorCode: -8, Message: "invalid otp code"}
	ErrNotFound        = &Error{HttpCode: 404, ErrorCode: -9, Message: "not found"}
	ErrInvalidToken    = &Error{HttpCode: 403, ErrorCode: -10, Message: "invalid token"}
	ErrTooManyRetries  = &Error{HttpCode: 403, ErrorCode: -11, Message: "too many retries"}

	// users
	ErrPasswordValidationError  = &Error{HttpCode: 400, ErrorCode: -501, Message: "wrong password"}
	ErrUserNotFound             = &Error{HttpCode: 404, ErrorCode: -502, Message: "user not found"}
	ErrPhoneNumberAlreadyExists = &Error{HttpCode: 400, ErrorCode: -503, Message: "phone number already exists"}
	ErrWrongVerifyCode          = &Error{HttpCode: 400, ErrorCode: -504, Message: "wrong verify code"}
	ErrUsernameExists           = &Error{HttpCode: 400, ErrorCode: -507, Message: "username already exists"}
)
