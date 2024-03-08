package role

import (
	"github.com/kimchhung/gva/extra/app/module/dashboard/role/controller"
	"github.com/kimchhung/gva/extra/app/module/dashboard/role/repository"
	"github.com/kimchhung/gva/extra/app/module/dashboard/role/service"
	"github.com/kimchhung/gva/extra/internal/rctrl"

	"go.uber.org/fx"
)

// Register bulkly
var NewRoleModule = fx.Module("RoleModule",
	// Register Repository & Service
	fx.Provide(repository.NewRoleRepository),
	fx.Provide(service.NewRoleService),

	// Regiser Controller
	fx.Provide(fx.Annotate(controller.NewRoleController,
		fx.As(new(rctrl.Controller)),
		fx.ResultTags(`group:"dashboard-controller"`),
	)),
)
