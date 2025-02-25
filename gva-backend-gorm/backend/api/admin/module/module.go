package module

import (
	"backend/api/admin/module/adminrole"
	"backend/api/admin/module/auth"
	"backend/api/admin/module/configuration"
	"backend/api/admin/module/index"
	"backend/api/admin/module/operationlog"
	"backend/app/common/controller"
	"backend/internal/ctr"

	"backend/api/admin/module/admin"

	"backend/api/admin/module/permission"

	"backend/api/admin/module/blog"
	"go.uber.org/fx"
	// #inject:moduleImport (do not remove this comment, it is used by the code generator)
)

var NewAdminModules = fx.Module("admin-module",
	auth.AuthModule,
	index.IndexModule,

	admin.AdminModule,
	adminrole.AdminRoleModule,
	permission.PermissionModule,
	configuration.ConfigurationModule,
	operationlog.OperationLogModule,

	blog.BlogModule,
	// #inject:module (do not remove this comment, it is used by the code generator)
	// Add Router
	fx.Provide(
		fx.Annotate(NewRouter,
			// convert type *Router => ctr.ModuleRouter
			fx.As(new(ctr.ModuleRouter)),

			// take group params from container => []ctr.CTR -> NewRouter
			fx.ParamTags(controller.TagAdminController),

			// register to container as member of module group
			fx.ResultTags(controller.TagModule),
		),
	),
)
