package bootstrap

import "go.uber.org/fx"

var Module = fx.Module("bootstrap",
	fx.Provide(New),
	fx.Invoke(
		func(b *Bootstrap) {
			b.setup()
		},
	),
)
