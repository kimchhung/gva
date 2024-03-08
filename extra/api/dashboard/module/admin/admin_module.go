package admin

import (
	"github.com/kimchhung/gva/extra/api/dashboard/module/admin/controller"
	"github.com/kimchhung/gva/extra/api/dashboard/module/admin/repository"
	"github.com/kimchhung/gva/extra/api/dashboard/module/admin/service"
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
		fx.ResultTags(`group:"dashboard-controller"`),
	)),
)
