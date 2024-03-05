package admin

import (
	"github.com/kimchhung/gva/extra/app/module/admin/controller"
	"github.com/kimchhung/gva/extra/app/module/admin/repository"
	"github.com/kimchhung/gva/extra/app/module/admin/service"
	"github.com/kimchhung/gva/extra/internal/rctrl"

	"go.uber.org/fx"
)

// Register bulkly
var NewAdminModule = fx.Module("AdminModule",
	// Register Repository & Service
	fx.Provide(repository.NewAdminRepository),
	fx.Provide(service.NewAdminService),

	// Regiser Controller
	fx.Provide(fx.Annotate(
		controller.NewAdminController,
		fx.As(new(rctrl.Controller)),
		fx.ResultTags(`group:"controllers"`),
	)),
)
