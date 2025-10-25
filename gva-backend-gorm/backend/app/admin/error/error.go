package adminerror

import (
	coreerror "backend/core/error"
	"net/http"
)

var (
	// admins
	ErrUsernameExists = coreerror.NewPublicError(http.StatusConflict, -500, "username_already_exists")

	// roles (400 - 499)
	ErrAdminRoleIsInUse = coreerror.NewPublicError(http.StatusConflict, -400, "role_is_in_use")
)
