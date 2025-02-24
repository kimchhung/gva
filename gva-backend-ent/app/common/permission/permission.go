package permission

import (
	"fmt"
	"strings"

	appctx "github.com/gva/app/common/context"
	apperror "github.com/gva/app/common/error"

	"github.com/labstack/echo/v4"
)

var (
	validGroups = make(map[PermissionGroup]struct{})
	validscopes = make(map[permissionScope]struct{})
)

func Groups() (groups []PermissionGroup) {
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

func HasGroup(group PermissionGroup) bool {
	_, has := validGroups[group]
	return has
}

func HasKey(scope permissionScope) bool {
	_, has := validscopes[scope]
	return has
}

func newKey(group PermissionGroup, action PermissionAction) permissionScope {
	key := permissionScope(fmt.Sprintf("%s%s%s", group, PermissionSeperator, action))

	validGroups[group] = struct{}{}
	validscopes[key] = struct{}{}
	return key
}

type (
	// Admin_Role:View
	permissionScope string

	// Admin | Admin_Role
	PermissionGroup string

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

func (k permissionScope) Value() (group PermissionGroup, action PermissionAction, err error) {
	if err := k.Valid(); err != nil {
		return group, action, err
	}

	parts := strings.SplitN(string(k), PermissionSeperator, 2)
	return PermissionGroup(parts[0]), PermissionAction(parts[1]), nil
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

func RequireAny(permissions ...permissionScope) echo.HandlerFunc {
	requireds := make(map[string]struct{}, len(permissions))
	for _, p := range permissions {
		requireds[string(p)] = struct{}{}
	}

	return func(c echo.Context) error {
		adminCtx := appctx.MustAdminContext(c.Request().Context())
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
