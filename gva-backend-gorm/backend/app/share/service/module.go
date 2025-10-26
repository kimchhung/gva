package service

import (
	"go.uber.org/fx"
)

func NewServiceModule() fx.Option {
	return fx.Module("ServiceModule",
		fx.Provide(
			NewPasswordService,
			NewJwtService,
			NewIPService,
			NewS3Service,
			NewTOTPService,
			NewRedisService,
		),
	)
}
