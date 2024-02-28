package main

import (
	fxzerolog "github.com/efectn/fx-zerolog"
	"github.com/kimchhung/gva/app/common"
	"github.com/kimchhung/gva/app/middleware"
	"github.com/kimchhung/gva/app/module"
	"go.uber.org/fx"

	"github.com/kimchhung/gva/app/module/admin"
	"github.com/kimchhung/gva/app/module/auth"
	"github.com/kimchhung/gva/app/module/permission"
	"github.com/kimchhung/gva/app/module/role"

	"github.com/kimchhung/gva/config"
	"github.com/kimchhung/gva/internal/bootstrap"
	"github.com/kimchhung/gva/internal/bootstrap/database"

	// #inject:moduleImport (do not remove this comment, it is used by the code generator)
	_ "github.com/kimchhung/gva/internal/ent/runtime"
)

// @title GVA API
// @version 1.0
// @description GO VUE ADMIN Boilerplate
// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
func main() {
	fx.New(
		// Provide patterns
		fx.Provide(config.NewConfig),
		fx.Provide(bootstrap.NewLogger),
		fx.Provide(bootstrap.NewFiber),
		fx.Provide(database.NewDatabase),
		fx.Provide(middleware.NewMiddleware),
		fx.Provide(
			fx.Annotate(module.NewRouter,
				fx.ParamTags(`group:"routers"`),
			),
		),

		common.NewCommonModule,
		admin.NewAdminModule,
		permission.NewPermissionModule,
		role.NewRoleModule,
		auth.NewAuthModule,
		// #inject:module (do not remove this comment, it is used by the code generator)

		// Start Application
		fx.Invoke(bootstrap.Start),

		// Define logger
		fx.WithLogger(fxzerolog.InitPtr()),
	).Run()
}
