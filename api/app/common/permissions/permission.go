package permissions

import (
	"slices"

	"github.com/gofiber/fiber/v2"
	"github.com/kimchhung/gva/app/common/contexts"
	app_err "github.com/kimchhung/gva/app/common/error"
)

type PermissionKey string

func RequireAny(permissions ...PermissionKey) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		adminCtx := contexts.MustAdminContext(c.UserContext())
		if adminCtx.IsSuperAdmin() {
			return c.Next()
		}

		required := make(map[string]struct{})
		for _, p := range permissions {
			required[string(p)] = struct{}{}
		}

		for _, rolePermission := range adminCtx.PermissionNames() {
			if _, ok := required[rolePermission]; ok {
				return c.Next()
			}
		}

		return app_err.ErrUnauthorized
	}
}

func RequireAll(permissions ...PermissionKey) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		adminCtx := contexts.MustAdminContext(c.UserContext())
		if adminCtx.IsSuperAdmin() {
			return c.Next()
		}

		required := make(map[string]struct{})
		for _, p := range permissions {
			required[string(p)] = struct{}{}
		}

		// Check if all required permissions are present.
		for key := range required {
			if !slices.Contains(adminCtx.PermissionNames(), key) {
				return app_err.ErrUnauthorized
			}
		}

		return c.Next()
	}
}
