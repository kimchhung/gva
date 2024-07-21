package department

import (
	"github.com/gva/app/common/controller"
	"go.uber.org/fx"
)

// Register bulkly
var DepartmentModule = fx.Module("DepartmentModule",
	// Register Service
	fx.Provide(NewDepartmentService),

	// Regiser Controller
	controller.ProvideAdminController(NewDepartmentController),
)
