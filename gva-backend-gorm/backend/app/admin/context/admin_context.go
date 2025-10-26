package context

import (
	"context"

	"backend/app/share/model"
	corecontext "backend/core/context"
)

const (
	AdminContextKey = "admin_context"
)

type (
	AdminContextOption func(*AdminContext)
)

type AdminContext struct {
	Admin        *model.Admin
	isSuperAdmin bool
	scopes       []string

	EndpointScopes []string
}

func GetAdminContext(ctx context.Context) (*AdminContext, error) {
	return corecontext.Get[*AdminContext](ctx, AdminContextKey)
}

func SetAdminContext(ctx context.Context, adminctx *AdminContext) {
	corecontext.Set(ctx, AdminContextKey, adminctx)
}

func MustAdminContext(ctx context.Context) *AdminContext {
	return corecontext.Must[*AdminContext](ctx, AdminContextKey)
}

// a context help handling error
func NewAdminContext(opts ...AdminContextOption) *AdminContext {
	adminctx := new(AdminContext)
	adminctx.Modify(opts...)

	return adminctx
}

func (c *AdminContext) Modify(opts ...AdminContextOption) {
	for _, opt := range opts {
		opt(c)
	}
}

func (ctx *AdminContext) PermissionScopes() []string {
	return ctx.scopes
}

func (ctx *AdminContext) IsSuperAdmin() bool {
	return ctx.isSuperAdmin
}

func WithAdmin(admin *model.Admin) AdminContextOption {
	return func(ac *AdminContext) {
		ac.Admin = admin
		ac.scopes = admin.PermissionScope
		ac.isSuperAdmin = admin.IsSuperAdmin
	}
}
