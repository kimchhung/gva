package index

import (
	"github.com/kimchhung/gva/extra/app/constant"
	"github.com/kimchhung/gva/extra/internal/echoc"
	"go.uber.org/fx"
)

// Register bulkly
var IndexModule = fx.Module("IndexModule",
	fx.Provide(NewIndexService),

	// Regiser Controller
	fx.Provide(fx.Annotate(NewIndexController,
		fx.As(new(echoc.Controller)),
		fx.ResultTags(constant.TagWebController),
	)),
)
