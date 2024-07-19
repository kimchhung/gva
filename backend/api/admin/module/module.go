package module

import (
	"github.com/gva/api/admin/module/admin"
	"github.com/gva/api/admin/module/auth"
	"github.com/gva/api/admin/module/index"
	"github.com/gva/api/admin/module/menu"
	"github.com/gva/api/admin/module/permission"

	"github.com/gva/app/common/controller"
	"github.com/gva/internal/echoc"

	"github.com/gva/api/admin/module/todo"
	"go.uber.org/fx"
	// #inject:moduleImport (do not remove this comment, it is used by the code generator)
)

var NewAdminModules = fx.Module("admin-module",
	admin.AdminModule,
	auth.AuthModule,
	menu.MenuModule,
	permission.PermissionModule,
	index.IndexModule,

	todo.TodoModule,
	// #inject:module (do not remove this comment, it is used by the code generator)
	// Add Router
	fx.Provide(
		fx.Annotate(NewRouter,
			// convert type *Router => echoc.ModuleRouter
			fx.As(new(echoc.ModuleRouter)),

			// take group params from container => []echoc.Controller -> NewRouter
			fx.ParamTags(controller.TagAdminController),

			// register to container as member of module group
			fx.ResultTags(controller.TagModule),
		),
	),
)
