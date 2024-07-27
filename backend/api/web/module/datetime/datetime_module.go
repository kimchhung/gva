package mdatetime

import (
	"github.com/gva/app/common/controller"
	"github.com/gva/internal/ctr"

	"go.uber.org/fx"
)

// Register bulkly
var DatetimeModule = fx.Module("IndexModule",
	fx.Provide(NewIndexService),

	// Regiser Controller
	fx.Provide(fx.Annotate(NewIndexController,
		fx.As(new(ctr.CTR)),
		fx.ResultTags(controller.TagWebController),
	)),

	fx.Invoke(BackgroundNow),
)
