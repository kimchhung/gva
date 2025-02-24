package service

import (
	"go.uber.org/fx"
)

func NewServiceModule() fx.Option {
	return fx.Module("ServiceModule",
		fx.Provide(NewPubsubService),
		fx.Provide(NewPasswordService),
		fx.Provide(NewJwtService),
		fx.Provide(NewIPService),
		fx.Provide(NewS3Service),
		fx.Provide(NewTOTPService),
		fx.Provide(NewRedisService),
	)
}
