package adminerror

import (
	apperror "backend/app/share/error"
	"net/http"
)

var (

	// users
	ErrEmailExists = apperror.NewPublicError(http.StatusConflict, -600, "email_already_exists")

	// recaptcha
	ErrRecaptchaInvalid = apperror.NewPublicError(http.StatusBadRequest, -703, "recaptcha_invalid")

	// send mail
	ErrSendMailFailed = apperror.NewPublicError(http.StatusInternalServerError, -800, "send_mail_failed")

	// admins
	ErrUsernameExists        = apperror.NewPublicError(http.StatusConflict, -500, "username_already_exists")
	ErrAdminWhitelistInvalid = apperror.NewPublicError(http.StatusConflict, -501, "admin_whitelist_invalid")

	// roles (400 - 499)
	ErrAdminRoleIsInUse = apperror.NewPublicError(http.StatusConflict, -400, "role_is_in_use")
)
