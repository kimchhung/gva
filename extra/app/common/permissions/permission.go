package permissions

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kimchhung/gva/extra/app/common/contexts"
	app_err "github.com/kimchhung/gva/extra/app/common/error"
)

type PermissionKey string

func RequireAny(permissions ...PermissionKey) fiber.Handler {
	return func(c *fiber.Ctx) error {
		adminCtx, err := contexts.GetAdminContext(c.UserContext())
		if err != nil {
			return err
		}

		if adminCtx.IsSuperAdmin() {
			return nil
		}

		rolePermissionsSet := make(map[string]struct{})
		for _, rolePermission := range adminCtx.PermissionNames() {
			rolePermissionsSet[rolePermission] = struct{}{}
		}

		for _, p := range permissions {
			if _, exists := rolePermissionsSet[string(p)]; exists {
				return nil // Found a matching permission, no need to check further
			}
		}

		return app_err.ErrUnauthorized // None of the required permissions were found
	}
}

func RequireAll(permissions ...PermissionKey) fiber.Handler {
	return func(c *fiber.Ctx) error {
		adminCtx, err := contexts.GetAdminContext(c.UserContext())
		if err != nil {
			return err
		}

		if adminCtx.IsSuperAdmin() {
			return nil
		}

		requireds := make(map[string]struct{})
		for _, p := range permissions {
			requireds[string(p)] = struct{}{}
		}

		for _, adminPermssion := range adminCtx.PermissionNames() {
			delete(requireds, adminPermssion)

			if len(requireds) == 0 {
				return nil
			}
		}

		return app_err.ErrUnauthorized
	}
}
