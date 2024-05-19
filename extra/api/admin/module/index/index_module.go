package index

import (
	"github.com/kimchhung/gva/extra/app/constant"
	"go.uber.org/fx"
)

// Register bulkly
var IndexModule = fx.Module("IndexModule",
	fx.Provide(NewIndexService),

	// Regiser Controller
	constant.ProvideAdminController(NewIndexController),
)
