package common

import (
	"github.com/kimchhung/gva/app/common/services"
	"go.uber.org/fx"
)

var NewCommonModule = fx.Module("CommonModule",
	fx.Provide(services.NewPasswordService),
	fx.Provide(services.NewJwtService),
)
