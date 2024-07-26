package bootstrap

import (
	"github.com/gva/internal/pubsub"
	"go.uber.org/fx"
)

var Module = fx.Module("bootstrap",
	fx.Provide(NewBootstrap),
	fx.Provide(
		fx.Annotate(
			pubsub.NewMemoryPubsub,
			fx.As(new(pubsub.Pubsub)),
		),
	),
	fx.Invoke(
		func(b *Bootstrap) {
			b.setup()
		},
	),
)
