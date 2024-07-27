package resolver

import (
	_ "github.com/gva/api/web/graph/generated"
	mdatetime "github.com/gva/api/web/module/datetime"
	"github.com/gva/internal/bootstrap/database"
	"github.com/gva/internal/pubsub"
)

type Resolver struct {
	db         *database.Database
	datetime_s *mdatetime.DatetimeService
	pubsub     pubsub.Pubsub
}

func NewResolver(
	db *database.Database,
	datetime_s *mdatetime.DatetimeService,
	psub pubsub.Pubsub,
) *Resolver {
	return &Resolver{
		db:         db,
		datetime_s: datetime_s,
		pubsub:     psub,
	}
}
