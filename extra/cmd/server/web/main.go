package main

import (
	fxzerolog "github.com/efectn/fx-zerolog"
	"github.com/kimchhung/gva/extra/app/common"
	"go.uber.org/fx"

	"github.com/kimchhung/gva/extra/app/module/web"

	"github.com/kimchhung/gva/extra/internal/bootstrap"

	// #inject:moduleImport (do not remove this comment, it is used by the code generator)
	_ "github.com/kimchhung/gva/extra/docs"
	_ "github.com/kimchhung/gva/extra/internal/ent/runtime"
)

// @title Web API
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
		common.NewCommonModule,

		// * web modules * //
		web.NewWebModules,

		// #inject:module (do not remove this comment, it is used by the code generator)

		// Start Application
		fx.Invoke(bootstrap.Start),

		// Define logger
		fx.WithLogger(fxzerolog.InitPtr()),
	).Run()
}
