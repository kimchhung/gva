package index

import (
	"github.com/gva/app/constant"

	"go.uber.org/fx"
)

// Register bulkly
var IndexModule = fx.Module("IndexModule",
	fx.Provide(NewIndexService),

	// Regiser Controller
	constant.ProvideAdminController(NewIndexController),
)
