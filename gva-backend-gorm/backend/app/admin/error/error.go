package adminerror

import (
	coreerror "backend/core/error"
	"net/http"
)

var (

	// users
	ErrEmailExists = coreerror.NewPublicError(http.StatusConflict, -600, "email_already_exists")

	// recaptcha
	ErrRecaptchaInvalid = coreerror.NewPublicError(http.StatusBadRequest, -703, "recaptcha_invalid")

	// send mail
	ErrSendMailFailed = coreerror.NewPublicError(http.StatusInternalServerError, -800, "send_mail_failed")

	// admins
	ErrUsernameExists        = coreerror.NewPublicError(http.StatusConflict, -500, "username_already_exists")
	ErrAdminWhitelistInvalid = coreerror.NewPublicError(http.StatusConflict, -501, "admin_whitelist_invalid")

	// roles (400 - 499)
	ErrAdminRoleIsInUse = coreerror.NewPublicError(http.StatusConflict, -400, "role_is_in_use")
)
