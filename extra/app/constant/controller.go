package constant

import (
	"github.com/kimchhung/gva/extra/internal/rctrl"
	"go.uber.org/fx"
)

const (
	TagAdminController = `group:"admin-controllers"`
	TagWebController   = `group:"web-controllers"`

	TagModule = `group:"modules"`
)

// register to container  type rctrl.Controller and tag as admin module
func ProvideAdminController(contructor any) fx.Option {
	return fx.Provide(
		fx.Annotate(
			contructor,
			fx.As(new(rctrl.Controller)),
			fx.ResultTags(TagAdminController),
		),
	)
}

// register to container as type rctrl.Controller and tag as web module
func ProvideWebController(contructor any) fx.Option {
	return fx.Provide(
		fx.Annotate(
			contructor,
			fx.As(new(rctrl.Controller)),
			fx.ResultTags(TagWebController),
		),
	)
}
