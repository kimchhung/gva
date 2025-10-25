package router

import (
	"backend/internal/ctr"
	"fmt"

	"go.uber.org/fx"
)

type Group string

func (tag Group) Add(c any) fx.Option {
	return fx.Provide(
		fx.Annotate(c,
			fx.As(new(ctr.CTR)),
			fx.ResultTags(fmt.Sprintf(`group:"%s"`, string(tag))),
		),
	)
}
