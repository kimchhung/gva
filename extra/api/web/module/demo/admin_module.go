package demo

import (
	"github.com/kimchhung/gva/extra/api/web/module/demo/controller"
	"github.com/kimchhung/gva/extra/internal/rctrl"
	"go.uber.org/fx"
)

// Register bulkly
var NewDemoModule = fx.Module("DemoModule",

	// Regiser Controller
	fx.Provide(fx.Annotate(
		controller.NewDemoController,
		fx.As(new(rctrl.Controller)),
		fx.ResultTags(`group:"web-controller"`),
	)),
)
