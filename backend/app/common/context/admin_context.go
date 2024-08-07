package appctx

import (
	"context"
	"errors"
	"slices"

	apperror "github.com/gva/app/common/error"
	"github.com/gva/internal/ent"
)

const (
	RoleNameSuperAdmin = "SUPER_ADMIN"
)

type (
	AdminContextScope  struct{}
	AdminContextOption func(*AdminContext)
)

type AdminContext struct {
	context.Context
	Admin *ent.Admin

	isSuperAdmin bool
	permissions  []*ent.Permission
}

// a context help handling error
func NewAdminContext(ctx context.Context, opts ...AdminContextOption) *AdminContext {
	adminctx := new(AdminContext)
	adminctx.Modify(opts...)

	adminctx.Context = context.WithValue(ctx, AdminContextScope{}, adminctx)
	return adminctx
}

func (c *AdminContext) Modify(opts ...AdminContextOption) {
	for _, opt := range opts {
		opt(c)
	}
}

func GetAdminContext(ctx context.Context) (*AdminContext, error) {
	v, ok := ctx.(*AdminContext)
	if ok {
		return v, nil
	}

	v, ok = ctx.Value(AdminContextScope{}).(*AdminContext)
	if ok {
		return v, nil
	}

	return nil, errors.New("context does not contain AdminContext")
}

func (ctx *AdminContext) IsSuperAdmin() bool {
	if ctx.isSuperAdmin {
		return true
	}

	return slices.Contains(ctx.RoleNames(), RoleNameSuperAdmin)
}

func (ctx *AdminContext) Roles() []*ent.Role {
	if ctx.Admin == nil {
		return nil
	}

	return ctx.Admin.Edges.Roles
}

func (ctx *AdminContext) RoleNames() []string {
	roles := ctx.Roles()
	if len(roles) == 0 {
		return nil
	}

	names := make([]string, len(roles))
	for i, r := range roles {
		if r.Name == RoleNameSuperAdmin {
			ctx.isSuperAdmin = true
		}

		names[i] = r.Name
	}

	return names
}

func (ctx *AdminContext) Permissions() []*ent.Permission {
	if ctx.Admin == nil || ctx.Admin.Edges.Roles == nil {
		return nil
	}

	if len(ctx.permissions) > 0 {
		return ctx.permissions
	}

	for _, role := range ctx.Admin.Edges.Roles {
		ctx.permissions = append(ctx.permissions, role.Edges.Permissions...)
	}

	return ctx.permissions
}

func (ctx *AdminContext) PermissionScopes() []string {
	permissions := ctx.Permissions()

	keys := make([]string, len(permissions))
	for i, p := range permissions {
		keys[i] = p.Scope
	}

	return keys
}

func WithAdmin(admin *ent.Admin) AdminContextOption {
	return func(ac *AdminContext) {
		ac.Admin = admin
	}
}

func MustAdminContext(ctx context.Context) *AdminContext {
	actx, err := GetAdminContext(ctx)
	if err != nil {
		panic(apperror.ErrUnauthorized)
	}

	return actx
}
