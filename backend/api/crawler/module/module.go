package module

import (
	"github.com/gva/api/crawler/module/comic"
	"github.com/gva/api/crawler/module/index"
	"github.com/gva/app/constant"
	"github.com/gva/internal/echoc"

	"go.uber.org/fx"
)

var NewCrawlerModules = fx.Module("crawler-module",
	index.IndexModule,
	comic.ComicModuleModule,

	fx.Provide(
		fx.Annotate(NewRouter,
			// convert type *Router => echoc.ModuleRouter
			fx.As(new(echoc.ModuleRouter)),

			// take group params from container => []echoc.Controller -> NewRouter
			fx.ParamTags(constant.CrawlerAdminController),

			// register to container as member of module group
			fx.ResultTags(constant.TagModule),
		),
	),
)
