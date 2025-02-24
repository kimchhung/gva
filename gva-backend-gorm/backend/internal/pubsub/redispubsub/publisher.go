package redispubsub

import (
	"context"
	"encoding/json"

	"backend/internal/pubsub"

	"github.com/redis/go-redis/v9"
)

var _ pubsub.Publisher = (*redisPublisher)(nil)

type redisPublisher struct {
	client       redis.UniversalClient
	redisChannel string
}

// NewPublisher new publisher using redis
func NewPublisher(client redis.UniversalClient, redisChannel string) pubsub.Publisher {
	return &redisPublisher{
		client:       client,
		redisChannel: redisChannel,
	}
}

func (p redisPublisher) Publish(ctx context.Context, channelName string, data any) error {
	payload := pubsub.Payload{
		Topic: channelName,
		Data:  data,
	}

	bytes, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	return p.client.Publish(ctx, p.redisChannel, bytes).Err()
}
