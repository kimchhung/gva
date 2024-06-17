package module

import (
	"github.com/kimchhung/gva/backend/api/crawler/module/comic"
	"github.com/kimchhung/gva/backend/api/crawler/module/index"
	"github.com/kimchhung/gva/backend/app/constant"
	"github.com/kimchhung/gva/backend/internal/echoc"
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
