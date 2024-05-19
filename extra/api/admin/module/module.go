package module

import (
	admin "github.com/kimchhung/gva/extra/api/admin/module/admin"
	auth "github.com/kimchhung/gva/extra/api/admin/module/auth"
	index "github.com/kimchhung/gva/extra/api/admin/module/index"
	permission "github.com/kimchhung/gva/extra/api/admin/module/permission"
	route "github.com/kimchhung/gva/extra/api/admin/module/route"
	"github.com/kimchhung/gva/extra/app/constant"
	"github.com/kimchhung/gva/extra/internal/rctrl"
	"go.uber.org/fx"
)

var APIAdminModules = fx.Module("admin-module",
	admin.AdminModule,
	auth.AuthModule,
	route.RouteModule,
	permission.PermissionModule,
	index.IndexModule,

	// #inject:module (do not remove this comment, it is used by the code generator)
	// Add Router
	fx.Provide(
		fx.Annotate(NewRouter,
			// convert type *Router => rctrl.ModuleRouter
			fx.As(new(rctrl.ModuleRouter)),

			// take group params from container => []rctrl.Controller -> NewRouter
			fx.ParamTags(constant.TagAdminController),

			// register to container as member of module group
			fx.ResultTags(constant.TagModule),
		),
	),
)
