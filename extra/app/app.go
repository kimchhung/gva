package app

import (
	fxzerolog "github.com/efectn/fx-zerolog"
	"github.com/kimchhung/gva/extra/app/common"
	"github.com/kimchhung/gva/extra/config"
	"github.com/kimchhung/gva/extra/internal/bootstrap"
	"go.uber.org/fx"
)

func New(cfg *config.Config, opts ...fx.Option) *fx.App {

	return fx.New(
		// Provide config
		fx.Supply(cfg),

		/* Common Module */
		common.NewCommonModule,

		// Start Application
		fx.Invoke(bootstrap.Start),

		// Define logger
		fx.WithLogger(fxzerolog.InitPtr()),

		/* add web or admin modules */
		fx.Module("server", opts...),
	)
}
