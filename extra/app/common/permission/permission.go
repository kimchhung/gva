package permission

import (
	"strings"

	appctx "github.com/kimchhung/gva/extra/app/common/context"
	apperror "github.com/kimchhung/gva/extra/app/common/error"
	"github.com/kimchhung/gva/extra/internal/ent"

	"github.com/labstack/echo/v4"
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

func RequireAny(permissions ...PermissionKey) echo.HandlerFunc {
	return func(c echo.Context) error {
		adminCtx := appctx.MustAdminContext(c.Request().Context())
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

func RequireAll(permissions ...PermissionKey) echo.HandlerFunc {
	return func(c echo.Context) error {
		adminCtx := appctx.MustAdminContext(c.Request().Context())

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

// only supper admin can access
func OnlySuperAdmin() echo.HandlerFunc {
	return func(c echo.Context) error {
		adminCtx := appctx.MustAdminContext(c.Request().Context())
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
