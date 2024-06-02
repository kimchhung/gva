package constant

import (
	"github.com/kimchhung/gva/extra/internal/echoc"
	"go.uber.org/fx"
)

const (
	TagAdminController = `group:"admin-controllers"`
	TagWebController   = `group:"web-controllers"`

	TagModule = `group:"modules"`
)

// register to container  type echoc.Controller and tag as admin module
func ProvideAdminController(contructor any) fx.Option {
	return fx.Provide(
		fx.Annotate(
			contructor,
			fx.As(new(echoc.Controller)),
			fx.ResultTags(TagAdminController),
		),
	)
}

// register to container as type echoc.Controller and tag as web module
func ProvideWebController(contructor any) fx.Option {
	return fx.Provide(
		fx.Annotate(
			contructor,
			fx.As(new(echoc.Controller)),
			fx.ResultTags(TagWebController),
		),
	)
}