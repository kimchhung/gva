package service

import (
	"context"
	"time"

	"github.com/gva/app/database/schema/pxid"
	"github.com/gva/internal/pubsub"
	"github.com/gva/internal/pubsub/gochannel"

	"github.com/gva/internal/utils"
	"github.com/rs/zerolog"
)

type PubsubService struct {
	log zerolog.Logger

	localPubsub pubsub.Pubsub
}

func NewPubsubService(log *zerolog.Logger) *PubsubService {
	s := &PubsubService{
		localPubsub: gochannel.NewPubsub(make(chan pubsub.Payload), func() string {
			return string(pxid.New(""))
		}),
		log: log.With().Str("service", "PubsubService").Logger(),
	}
	return s
}

func (s *PubsubService) Listen() {
	go func() {
		reties := 20
		delay := time.Second

		utils.Retry(func() error {
			if err := s.localPubsub.Receive(context.TODO()); err != nil {
				s.log.Error().Err(err).Msg("Listen")
				return err
			}

			return nil
		}, reties, delay)
	}()
}

func (p *PubsubService) Local() pubsub.Pubsub {
	return p.localPubsub
}
