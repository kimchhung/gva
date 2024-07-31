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
	redisChannel []string
}

// NewRedisPublisher new publisher using redis
func NewPublisher(client *redis.Client, shuffleChannels []string) *redisPublisher {
	return &redisPublisher{
		client:       client,
		redisChannel: shuffleChannels,
	}
}

func (p redisPublisher) Pub(ctx context.Context, topic pubsub.Topic, data pubsub.Data) (err error) {
	payload := pubsub.Payload{
		Topic: topic,
	}

	dataByte, ok := data.([]byte)
	if ok {
		payload.Data = dataByte
	} else {
		payload.Data, err = json.Marshal(data)
		if err != nil {
			return err
		}
	}

	bytes, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	return p.client.Publish(ctx, p.redisChannel[0], bytes).Err()
}
