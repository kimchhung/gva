package pubsubredis

import (
	"context"
	"encoding/json"

	"github.com/gva/internal/pubsub"
	"github.com/redis/go-redis/v9"
)

var _ pubsub.Publisher = (*redisPublisher)(nil)

type redisPublisher struct {
	client       *redis.Client
	redisChannel string
}

// NewRedisPublisher new publisher using redis
func NewRedisPublisher(client *redis.Client, redisChannel string) pubsub.Publisher {
	return &redisPublisher{
		client:       client,
		redisChannel: redisChannel,
	}
}

func (p redisPublisher) Pub(ctx context.Context, topic pubsub.Topic, m pubsub.Payload) error {
	body := PublishRequest{
		Topic:   topic,
		Payload: m,
	}

	bytes, err := json.Marshal(body)
	if err != nil {
		return err
	}

	return p.client.Publish(ctx, p.redisChannel, bytes).Err()
}
