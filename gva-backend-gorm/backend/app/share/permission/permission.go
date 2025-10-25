package permission

import (
	"fmt"
	"strings"

	admincontext "backend/app/admin/context"

	apperror "backend/app/share/error"

	"github.com/labstack/echo/v4"
)

var (
	validGroups = make(map[TPermissionGroup]struct{})
	validscopes = make(map[permissionScope]struct{})
)

func Groups() (groups []TPermissionGroup) {
	for g := range validGroups {
		groups = append(groups, g)
	}
	return groups
}

func Scopes() (scopes []permissionScope) {
	for g := range validscopes {
		scopes = append(scopes, g)
	}
	return scopes
}

func HasGroup(group TPermissionGroup) bool {
	_, has := validGroups[group]
	return has
}

func HasKey(scope permissionScope) bool {
	_, has := validscopes[scope]
	return has
}

func newKey(group TPermissionGroup, action PermissionAction) permissionScope {
	key := permissionScope(fmt.Sprintf("%s%s%s", group, PermissionSeperator, action))
	if _, has := validGroups[group]; !has {
		validGroups[group] = struct{}{}
	}
	if _, has := validscopes[key]; !has {
		validscopes[key] = struct{}{}
	}
	return key
}

type (
	// Admin_Role:View
	permissionScope string

	// Admin | Admin_Role
	TPermissionGroup string

	// View | Add
	PermissionAction string
)

const (
	ActionSuper  PermissionAction = "super"
	ActionView   PermissionAction = "view"
	ActionAdd    PermissionAction = "add"
	ActionEdit   PermissionAction = "edit"
	ActionDelete PermissionAction = "delete"
)

var (
	PermissionSeperator = ":"
)

func (k permissionScope) Value() (group TPermissionGroup, action PermissionAction, err error) {
	if err := k.Valid(); err != nil {
		return group, action, err
	}

	parts := strings.SplitN(string(k), PermissionSeperator, 2)
	return TPermissionGroup(parts[0]), PermissionAction(parts[1]), nil
}

func (k permissionScope) Name() string {
	name := strings.ReplaceAll(string(k), "_", " ")
	name = strings.ReplaceAll(name, PermissionSeperator, " ")
	return name
}

func (k permissionScope) Valid() error {
	if HasKey(k) {
		return nil
	}

	return fmt.Errorf("invalid permission key %s", k)
}

func (p TPermissionGroup) Valid() error {
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

func RequireAnys(permissions ...permissionScope) echo.HandlerFunc {
	requireds := make(map[string]struct{}, len(permissions))

	return func(c echo.Context) error {
		// rctx, rerr := corecontext.Must(c.Request().Context())
		// if rerr != nil {
		// 	return rerr
		// }
		// rctx.LogFields.Scopes = scopes

		adminCtx := admincontext.MustAdminContext(c.Request().Context())
		if adminCtx.IsSuperAdmin() {
			return nil
		}

		for _, adminScope := range adminCtx.PermissionScopes() {
			if _, hasScope := requireds[adminScope]; hasScope {
				return nil
			}
		}

		return apperror.ErrUnauthorized // None of the required permissions were found
	}
}

func RequireAny(permissions ...permissionScope) echo.HandlerFunc {
	return RequireAnys(WithSuper(permissions...)...)
}

// only supper admin can access
func OnlySuperAdmin() echo.HandlerFunc {
	return func(c echo.Context) error {
		adminCtx := admincontext.MustAdminContext(c.Request().Context())
		if adminCtx.IsSuperAdmin() {
			return nil
		}

		return apperror.ErrUnauthorized // None of the required permissions were found
	}
}

func WithSuper(permissions ...permissionScope) []permissionScope {
	if len(permissions) == 0 {
		return []permissionScope{}
	}

	first := permissions[0]
	group, _, err := first.Value()
	if err != nil {
		panic(err)
	}

	superPermission := newKey(group, ActionSuper)
	if err := superPermission.Valid(); err != nil {
		panic(err)
	}

	return append(permissions, superPermission)
}
