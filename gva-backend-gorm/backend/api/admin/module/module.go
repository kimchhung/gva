package module

import (
	"backend/api/admin/module/adminrole"
	"backend/api/admin/module/auth"
	"backend/api/admin/module/configuration"
	"backend/api/admin/module/index"
	"backend/api/admin/module/operationlog"
	"backend/app/common/controller"

	"backend/api/admin/module/admin"

	"backend/api/admin/module/permission"

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
	// Add Router
	controller.Admin.AddRouter(NewRouter),
)
