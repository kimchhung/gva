package service

import (
	"context"
	"time"

	"backend/internal/bootstrap/database"
	"backend/internal/pubsub"
	"backend/internal/pubsub/redispubsub"
	"backend/internal/pxid"
	"backend/utils"

	"go.uber.org/zap"
)

type PubsubService struct {
	log *zap.Logger

	redisPubsub pubsub.Pubsub
	redis       *database.Redis
}

func NewPubsubService(log *zap.Logger, redis *database.Redis) *PubsubService {
	s := &PubsubService{
		redis: redis,
		redisPubsub: redispubsub.NewPubsub(redis, "global", func() string {
			return string(pxid.New("global"))
		}),
		log: log.Named("pubsub"),
	}

	s.listen(context.Background())
	return s
}

func (s *PubsubService) listen(ctx context.Context) {
	s.log.Info("listen")

	go func() {
		time.Sleep(time.Second * 5)
		backOff := utils.NewExponentialBackOff(time.Second, time.Minute, 1.5, 20)
		for backOff.ShouldRetry() {
			var err error
			if err = s.redisPubsub.Receive(ctx); err != nil {
				s.log.Error("s.redis.Receive",
					zap.Error(err),
					zap.Uint("attempt", backOff.Attempt()),
					zap.Uint("maxAttemp", backOff.MaxAttempt()),
				)
				backOff.Wait()
				continue
			}

			backOff.Reset()
		}

		backOff.Reset()
	}()
}
