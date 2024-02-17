package main

import (
	fxzerolog "github.com/efectn/fx-zerolog"
	"github.com/kimchhung/gva/app/middleware"
	"github.com/kimchhung/gva/app/module"
	"go.uber.org/fx"

	"github.com/kimchhung/gva/config"
	"github.com/kimchhung/gva/internal/bootstrap"
	"github.com/kimchhung/gva/internal/bootstrap/database"

	"github.com/kimchhung/gva/app/module/admin"

	"github.com/kimchhung/gva/app/module/permission"
	"github.com/kimchhung/gva/app/module/role"
	// #inject:moduleImport (do not remove this comment, it is used by the code generator)
)

// @title Fiber Example API
// @version 1.0
// @description This is a sample swagger for Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
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

		admin.NewAdminModule,
		permission.NewPermissionModule,
		role.NewRoleModule,
		// #inject:module (do not remove this comment, it is used by the code generator)

		// Start Application
		fx.Invoke(bootstrap.Start),

		// Define logger
		fx.WithLogger(fxzerolog.InitPtr()),
	).Run()
}
