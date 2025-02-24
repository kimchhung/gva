package appctx

import (
	"context"
	"errors"

	apperror "backend/app/common/error"
	"backend/app/common/model"
)

const (
	RoleIdSuperAdmin     = 1
	RoleNameIDSuperAdmin = "SUPER_ADMIN"
)

type (
	AdminContextScope  struct{}
	AdminContextOption func(*AdminContext)
)

type AdminContext struct {
	context.Context
	Admin *model.Admin

	isSuperAdmin bool
	scopes       []string
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

func MustAdminContext(ctx context.Context) *AdminContext {
	actx, err := GetAdminContext(ctx)
	if err != nil {
		panic(apperror.ErrUnauthorized)
	}

	return actx
}
