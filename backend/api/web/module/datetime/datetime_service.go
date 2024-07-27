package mdatetime

import (
	"context"
	"time"

	"github.com/gva/internal/bootstrap/database"
	"github.com/gva/internal/pubsub"
)

type DatetimeService struct {
	db   *database.Database
	psub pubsub.Pubsub
}

// NewAuthService initializes a new AuthService with a JwtService and a UserStore.
func NewIndexService(db *database.Database, psub pubsub.Pubsub) *DatetimeService {
	return &DatetimeService{
		db:   db,
		psub: psub,
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
	sub, err := s.psub.Sub(ctx, "now")
	if err != nil {
		return nil, err
	}

	go func() {
		defer sub.UnSub()

		for {
			select {
			case <-ctx.Done():
				return
			default:
				for payload := range sub.Payload() {
					if time, ok := payload.(time.Time); ok {
						ch <- &time
					}
				}
			}
		}
	}()

	return ch, nil
}
