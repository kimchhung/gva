package resolver

import (
	"github.com/gva/api/web/module/index"
	"github.com/gva/internal/bootstrap/database"
)

type Resolver struct {
	db      *database.Database
	index_s *index.IndexService
}

func NewResolver(
	db *database.Database,
	index_s *index.IndexService,
) *Resolver {
	return &Resolver{
		db:      db,
		index_s: index_s,
	}
}
