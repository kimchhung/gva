package gqlgen

import (
	"context"
	"encoding/json"

	"github.com/redis/go-redis/v9"
)

// Publisher ...
type Publisher interface {
	// Publish publishes payload data to channel
	Publish(channelName string, data []byte) error
}

var _ Publisher = (*redisPublisher)(nil)

type redisPublisher struct {
	client       *redis.Client
	redisChannel string
}

// NewRedisPublisher new publisher using redis
func NewRedisPublisher(client *redis.Client, redisChannel string) Publisher {
	return &redisPublisher{
		client:       client,
		redisChannel: redisChannel,
	}
}

func (p redisPublisher) Publish(channelName string, data []byte) error {
	payload := Payload{
		ChannelName: channelName,
		Data:        data,
	}

	bytes, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	return p.client.Publish(context.TODO(), p.redisChannel, bytes).Err()
}
