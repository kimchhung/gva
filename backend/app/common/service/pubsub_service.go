package service

import (
	"context"

	"github.com/gva/internal/bootstrap/database"
	"github.com/gva/internal/pubsub"
	pubsubchannel "github.com/gva/internal/pubsub/channel"
	pubsubredis "github.com/gva/internal/pubsub/redis"
	"github.com/rs/zerolog"
)

type PubsubService struct {
	log   zerolog.Logger
	redis *database.Redis

	localPubsub       pubsub.Pubsub
	subsciptionPubsub pubsub.Pubsub
}

func NewPubsubService(log *zerolog.Logger, redis *database.Redis) *PubsubService {
	s := &PubsubService{
		localPubsub: pubsubchannel.NewMemoryPubsub(),
		subsciptionPubsub: pubsubredis.NewRedisPubSub(
			redis.Client,
			"gql:subsciption",
		),

		redis: redis,
		log:   log.With().Str("service", "PubsubService").Logger(),
	}
	return s
}

func (s *PubsubService) Listen() {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				s.log.Error().Any("recover", r).Msg("Listen")
			}
		}()

		if err := s.localPubsub.Listen(context.TODO()); err != nil {
			s.log.Error().Err(err).Msg("Listen")
		}
	}()
}

func (p *PubsubService) Local() pubsub.Pubsub {
	return p.localPubsub
}

func (p *PubsubService) GqlSub() pubsub.Pubsub {
	return p.localPubsub
}
