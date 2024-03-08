package app

import (
	fxzerolog "github.com/efectn/fx-zerolog"
	"github.com/kimchhung/gva/extra/app/common"
	"github.com/kimchhung/gva/extra/app/module"
	"github.com/kimchhung/gva/extra/config"
	"github.com/kimchhung/gva/extra/internal/bootstrap"
	"go.uber.org/fx"
)

func New(cfg *config.Config) *fx.App {
	newConfig := fx.Provide(
		func() *config.Config {
			return cfg
		},
	)

	return fx.New(
		// Provide config
		newConfig,

		/* Common Module */
		common.NewCommonModule,

		/* Web, Dashboard |> Module <| */
		module.NewModules(cfg),
		// #inject:module (do not remove this comment, it is used by the code generator)

		// Start Application
		fx.Invoke(bootstrap.Start),

		// Define logger
		fx.WithLogger(fxzerolog.InitPtr()),
	)
}
