package stream

import (
	"context"

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

func (p redisPublisher) Pub(ctx context.Context, topic pubsub.Topic, data pubsub.Data) (err error) {
	return p.client.XAdd(ctx, &redis.XAddArgs{
		Stream: "room",
		ID:     "1",
		MaxLen: 1, // max elements to store
		Values: map[string]interface{}{
			"topic": topic,
			"data":  data,
		},
	}).Err()
}
