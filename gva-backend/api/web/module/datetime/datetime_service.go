package mdatetime

import (
	"context"
	"time"

	"github.com/gva/app/common/service"
	"github.com/gva/internal/bootstrap/database"
	"github.com/gva/internal/pubsub"
)

type DatetimeService struct {
	db       *database.Database
	pubsub_s *service.PubsubService
}

// NewAuthService initializes a new AuthService with a JwtService and a UserStore.
func NewIndexService(db *database.Database, pubsub_s *service.PubsubService) *DatetimeService {
	return &DatetimeService{
		db:       db,
		pubsub_s: pubsub_s,
	}
}

func (s *DatetimeService) Now(ctx context.Context) (*time.Time, error) {
	rows, err := s.db.QueryContext(ctx, "SELECT NOW()")
	if err != nil {
		panic(err)
	}

	var now time.Time
	for rows.Next() {
		if err := rows.Scan(&now); err != nil {
			panic(err)
		}
	}

	return &now, nil
}

func (s *DatetimeService) NowChannel(ctx context.Context) (<-chan *time.Time, error) {
	ch := make(chan *time.Time)

	sub, err := s.pubsub_s.Local().Subscribe("now", ch,
		func(_ string, payload pubsub.Payload) (interface{}, error) {
			return payload.Data.(*time.Time), nil
		},
	)
	if err != nil {
		return nil, err
	}
	go pubsub.CloseSubscription(ctx, sub, ch)
	return ch, nil
}
