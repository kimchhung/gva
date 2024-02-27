package contexts

import (
	"context"
	"errors"
	"slices"

	"github.com/kimchhung/gva/internal/ent"
)

const (
	RoleNameSuperAdmin = "SUPER_ADMIN"
)

type (
	AdminContextOption func(*AdminContext)
	AdminContext       struct {
		context.Context

		Admin *ent.Admin

		isSuperAdmin bool
		permissions  []*ent.Permission
	}
)

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

func (ctx *AdminContext) PermissionNames() []string {
	permissions := ctx.Permissions()

	names := make([]string, len(permissions))
	for i, p := range permissions {
		names[i] = p.Name
	}

	return names
}

func WithAdmin(admin *ent.Admin) AdminContextOption {
	return func(ac *AdminContext) {
		ac.Admin = admin
	}
}

func NewAdminContext(parentCtx context.Context, opts ...AdminContextOption) *AdminContext {
	ctx := &AdminContext{
		Context: parentCtx,
	}

	for _, opt := range opts {
		opt(ctx)
	}

	return ctx
}

func GetAdminContext(ctx context.Context) (*AdminContext, error) {
	actx, ok := ctx.(*AdminContext)
	if !ok {
		return nil, errors.New("context is not adminContext")
	}

	return actx, nil
}

func MustAdminContext(ctx context.Context) *AdminContext {
	actx, err := GetAdminContext(ctx)
	if err != nil {
		panic(err)
	}

	return actx
}
