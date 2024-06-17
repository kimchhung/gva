package module

import (
	admin "github.com/kimchhung/gva/backend/api/admin/module/admin"
	auth "github.com/kimchhung/gva/backend/api/admin/module/auth"
	index "github.com/kimchhung/gva/backend/api/admin/module/index"
	permission "github.com/kimchhung/gva/backend/api/admin/module/permission"
	route "github.com/kimchhung/gva/backend/api/admin/module/route"
	"github.com/kimchhung/gva/backend/app/constant"
	"github.com/kimchhung/gva/backend/internal/echoc"
	"go.uber.org/fx"
)

var NewAdminModules = fx.Module("admin-module",
	admin.AdminModule,
	auth.AuthModule,
	route.RouteModule,
	permission.PermissionModule,
	index.IndexModule,

	// #inject:module (do not remove this comment, it is used by the code generator)
	// Add Router
	fx.Provide(
		fx.Annotate(NewRouter,
			// convert type *Router => echoc.ModuleRouter
			fx.As(new(echoc.ModuleRouter)),

			// take group params from container => []echoc.Controller -> NewRouter
			fx.ParamTags(constant.TagAdminController),

			// register to container as member of module group
			fx.ResultTags(constant.TagModule),
		),
	),
)
