package index

import (
	"github.com/gva/app/common/controller"
	"github.com/gva/internal/echoc"

	"go.uber.org/fx"
)

// Register bulkly
var IndexModule = fx.Module("IndexModule",
	fx.Provide(NewIndexService),

	// Regiser Controller
	fx.Provide(fx.Annotate(NewIndexController,
		fx.As(new(echoc.Controller)),
		fx.ResultTags(controller.TagWebController),
	)),
)
