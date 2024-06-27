package bootstrap

import "go.uber.org/fx"

var Module = fx.Module("bootstrap",
	fx.Provide(NewBootstrap),
	fx.Invoke(
		func(b *Bootstrap) {
			b.setup()
		},
	),
)
