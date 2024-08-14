package gochannel

import (
	"context"

	"github.com/gva/internal/pubsub"
)

var _ pubsub.Publisher = (*Publisher)(nil)

type Publisher struct {
	publishReqeust chan pubsub.Payload
}

// NewPublisher new publisher using redis
func NewPublisher(publishReqeust chan pubsub.Payload) pubsub.Publisher {
	return &Publisher{
		publishReqeust: publishReqeust,
	}
}

func (p *Publisher) Publish(ctx context.Context, channelName string, data any) error {
	payload := pubsub.Payload{
		Topic: channelName,
		Data:  data,
	}

	select {
	case p.publishReqeust <- payload:
	case <-ctx.Done():
		return ctx.Err()
	}

	return nil
}
