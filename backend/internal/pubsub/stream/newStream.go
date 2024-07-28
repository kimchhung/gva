package stream

import (
	"github.com/gva/internal/pubsub"
	"github.com/redis/go-redis/v9"
)

type RedisSubSub struct {
	redisBroker
	redisPublisher
}

func NewRedisPubSub(client *redis.Client, channel string, opts ...redis.ChannelOption) pubsub.Pubsub {
	broker := redisBroker{
		client:           client,
		redisChannel:     channel,
		topics:           make(map[pubsub.Topic]topicSubscribers),
		addSubscriber:    make(chan Subscriber),
		removeSubscriber: make(chan Subscriber),
		opts:             opts,
	}

	publisher := redisPublisher{
		client:       client,
		redisChannel: channel,
	}

	return &RedisSubSub{
		broker,
		publisher,
	}
}
