package permission

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	appctx "github.com/kimchhung/gva/extra/app/common/context"
	apperror "github.com/kimchhung/gva/extra/app/common/error"
	"github.com/kimchhung/gva/extra/internal/ent"
)

var (
	groups []string
	keys   []PermissionKey
)

type PermissionKey string

func (p PermissionKey) Group() string {
	parts := strings.SplitN(string(p), ".", 2)
	if len(parts) > 1 {
		return parts[0]
	}
	return ""
}

func Groups() []string {
	list := make([]string, len(groups))
	copy(list, groups)
	return list
}

func Keys() []PermissionKey {
	list := make([]PermissionKey, len(keys))
	copy(list, keys)
	return list
}

func RequireAny(permissions ...PermissionKey) fiber.Handler {
	return func(c *fiber.Ctx) error {
		adminCtx := appctx.MustAdminContext(c.UserContext())
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

		return apperror.ErrUnauthorized // None of the required permissions were found
	}
}

func RequireAll(permissions ...PermissionKey) fiber.Handler {
	return func(c *fiber.Ctx) error {
		adminCtx := appctx.MustAdminContext(c.UserContext())

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

		return apperror.ErrUnauthorized
	}
}

func RequireSuperAdmin() fiber.Handler {
	return func(c *fiber.Ctx) error {
		adminCtx := appctx.MustAdminContext(c.UserContext())
		if adminCtx.IsSuperAdmin() {
			return nil
		}

		return apperror.ErrUnauthorized // None of the required permissions were found
	}
}

func createBulkPermissionDto(conn *ent.Client, group string, keys ...PermissionKey) []*ent.PermissionCreate {
	bulks := make([]*ent.PermissionCreate, len(keys))

	for i, key := range keys {
		bulks[i] = conn.Permission.Create().
			SetGroup(group).
			SetKey(string(key)).
			SetName(string(key)).
			SetOrder(i)
	}

	return bulks
}
