package adminmodule

import (
	adminmiddleware "backend/app/admin/middleware"
	"backend/app/admin/module/adminrole"
	"backend/app/admin/module/auth"
	"backend/app/admin/module/configuration"
	"backend/app/admin/module/index"
	"backend/app/admin/module/operationlog"
	adminrouter "backend/app/admin/router"
	"backend/core/router"

	"backend/app/admin/module/admin"

	"backend/app/admin/module/permission"

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
	// #inject:module (do not remove this comment, it is used by the code generator)

	// middleware
	fx.Provide(adminmiddleware.NewMiddleware),

	// register module router to global router
	router.Add(adminrouter.NewRouter),
)
