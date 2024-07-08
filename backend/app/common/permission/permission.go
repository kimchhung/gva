package permission

import (
	"fmt"
	"strings"

	appctx "github.com/gva/app/common/context"
	apperror "github.com/gva/app/common/error"
	"github.com/gva/internal/ent"

	"github.com/labstack/echo/v4"
)

var (
	validGroups = make(map[PermissionGroup]struct{})
	validKeys   = make(map[PermissionKey]struct{})
)

func Groups() (groups []PermissionGroup) {
	for g := range validGroups {
		groups = append(groups, g)
	}
	return groups
}

func Keys() (keys []PermissionKey) {
	for g := range validKeys {
		keys = append(keys, g)
	}
	return keys
}

func HasGroup(group PermissionGroup) bool {
	_, has := validGroups[group]
	return has
}

func HasKey(key PermissionKey) bool {
	_, has := validKeys[key]
	return has
}

func newKey(group PermissionGroup, action PermissionAction) PermissionKey {
	key := PermissionKey(fmt.Sprintf("%s.%s", group, action))

	validGroups[group] = struct{}{}
	validKeys[key] = struct{}{}
	return key
}

type (
	// Admin.View
	PermissionKey string

	// Admin
	PermissionGroup string

	// View
	PermissionAction string
)

const (
	ActionSuper  PermissionAction = "Super"
	ActionView   PermissionAction = "View"
	ActionAdd    PermissionAction = "Add"
	ActionEdit   PermissionAction = "Edit"
	ActionDelete PermissionAction = "Delete"
)

var (
	PermissionSeperator = "."
)

func (k PermissionKey) Value() (group PermissionGroup, action PermissionAction, err error) {
	if err := k.Valid(); err != nil {
		return group, action, err
	}

	parts := strings.SplitN(string(k), PermissionSeperator, 2)
	return PermissionGroup(parts[0]), PermissionAction(parts[1]), nil
}

func (k PermissionKey) Valid() error {
	if HasKey(k) {
		return nil
	}

	return fmt.Errorf("invalid permission key %s", k)
}

func (p PermissionGroup) Valid() error {
	if HasGroup(p) {
		return nil
	}

	return fmt.Errorf("invalid permission key %s", p)
}

func (p PermissionAction) Valid() error {
	switch p {
	case ActionSuper, ActionView, ActionAdd, ActionEdit, ActionDelete:
		return nil
	}

	return fmt.Errorf("invalid permission key %s", p)
}

func RequireAny(permissions ...PermissionKey) echo.HandlerFunc {
	return func(c echo.Context) error {
		adminCtx := appctx.MustAdminContext(c.Request().Context())
		if adminCtx.IsSuperAdmin() {
			return nil
		}

		rolePermissionsSet := make(map[string]struct{})
		for _, rolePermission := range adminCtx.PermissionKeys() {
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

		for _, adminPermssion := range adminCtx.PermissionKeys() {
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

func createBulkPermissionDto(conn *ent.Client, keys ...PermissionKey) []*ent.PermissionCreate {
	bulks := make([]*ent.PermissionCreate, len(keys))

	for i, key := range keys {
		group, _, err := key.Value()
		if err != nil {
			panic(err)
		}

		bulks[i] = conn.Permission.Create().
			SetGroup(string(group)).
			SetKey(string(key)).
			SetName(string(key)).
			SetOrder(i)
	}

	return bulks
}
