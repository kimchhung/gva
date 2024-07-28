package mdatetime

import (
	"context"
	"time"

	"github.com/gva/app/common/service"
	"github.com/gva/internal/bootstrap/database"
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

	if rows.Next() {
		if err := rows.Scan(&now); err != nil {
			panic(err)
		}
	}

	return &now, nil
}

func (s *DatetimeService) NowChannel(ctx context.Context) (<-chan *time.Time, error) {
	ch := make(chan *time.Time)
	sub, err := s.pubsub_s.Local().Sub(ctx, "now")
	if err != nil {
		return nil, err
	}

	go func() {
		for {
			select {
			case <-ctx.Done():
				sub.UnSub()
				return
			case payload := <-sub.Payload():
				if time, ok := payload.(time.Time); ok {
					ch <- &time
				}
			}
		}
	}()

	return ch, nil
}
