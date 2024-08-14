package bootstrap

import (
	"github.com/gva/app/common/service"
	"github.com/gva/internal/bootstrap/database"
	"go.uber.org/fx"
)

var Module = fx.Module("bootstrap",
	fx.Provide(NewLogger),
	fx.Provide(NewEcho),
	fx.Provide(database.NewDatabase),
	fx.Provide(database.NewRedis),
	fx.Provide(NewBootstrap),
	fx.Invoke(
		func(b *Bootstrap) {
			b.setup()
		},
		func(pubsub_s *service.PubsubService) {
			pubsub_s.Listen()
		},
	),
)
