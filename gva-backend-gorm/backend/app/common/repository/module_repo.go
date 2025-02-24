package repository

import (
	"go.uber.org/fx"
)

var dependencies = []fx.Option{}

func NewRepositoyModule() fx.Option {
	return fx.Module("RepositoryModule", dependencies...)
}
