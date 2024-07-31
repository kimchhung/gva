package pubsubredis

import (
	"github.com/gva/internal/pubsub"
	"github.com/redis/go-redis/v9"
)

type RedisSubSub struct {
	*redisBroker
	*redisPublisher
}

func NewRedisPubSub(client *redis.Client, channels []string, opts ...redis.ChannelOption) pubsub.Pubsub {
	broker := NewBroker(client, channels, opts...)
	publisher := NewPublisher(client, channels)

	return &RedisSubSub{
		broker,
		publisher,
	}
}
