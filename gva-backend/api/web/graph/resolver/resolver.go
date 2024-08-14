package resolver

import (
	_ "github.com/gva/api/web/graph/generated"
	mdatetime "github.com/gva/api/web/module/datetime"
	"github.com/gva/internal/bootstrap/database"
)

type Resolver struct {
	db         *database.Database
	datetime_s *mdatetime.DatetimeService
}

func NewResolver(
	db *database.Database,
	datetime_s *mdatetime.DatetimeService,
) *Resolver {
	return &Resolver{
		db:         db,
		datetime_s: datetime_s,
	}
}
