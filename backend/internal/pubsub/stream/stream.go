package stream

import (
	"context"
	"errors"

	"github.com/gva/app/database/schema/xid"
	"github.com/gva/internal/logger"
	"github.com/gva/internal/pubsub"
	"github.com/redis/go-redis/v9"
)

var _ pubsub.Broker = (*redisBroker)(nil)

type redisBroker struct {
	client       *redis.Client
	redisChannel string
	isStarted    bool

	topics           map[pubsub.Topic]topicSubscribers // topicName => topicSubscribers
	addSubscriber    chan Subscriber
	removeSubscriber chan Subscriber

	opts []redis.ChannelOption
}

type Subscriber struct {
	topic           pubsub.Topic
	subId           xid.ID
	payload         chan pubsub.Data
	unsubscribeFunc func() error
}

type topicSubscribers map[xid.ID]Subscriber

func (ms Subscriber) Data() <-chan pubsub.Data {
	return ms.payload
}

func (ms Subscriber) UnSub() error {
	return ms.unsubscribeFunc()
}

func (b *redisBroker) Sub(ctx context.Context, topic pubsub.Topic) (pubsub.SubResult, error) {
	ch := make(chan pubsub.Data, 1)
	subscriber := Subscriber{
		topic:   topic,
		subId:   xid.MustNew("sub"),
		payload: ch,
	}

	subscriber.unsubscribeFunc = func() error {
		b.removeSubscriber <- subscriber
		return nil
	}
	b.addSubscriber <- subscriber
	return subscriber, nil
}

func (m *redisBroker) Listen(ctx context.Context) error {
	if m.isStarted {
		return errors.New("can't listen more than once")
	}

	m.isStarted = true

	go func() {
		for {
			select {
			case <-ctx.Done():
				return

			case sub := <-m.addSubscriber:
				_, ok := m.topics[sub.topic]
				if !ok {
					m.topics[sub.topic] = topicSubscribers{
						sub.subId: sub,
					}
					continue
				}

				m.topics[sub.topic][sub.subId] = sub

			case sub := <-m.removeSubscriber:
				topicSub, ok := m.topics[sub.topic]
				if !ok {
					continue
				}
				delete(topicSub, sub.subId)
				close(sub.payload)
			}
		}
	}()

	for {
		logger.Log("stream start")
		streams, err := m.client.XRead(ctx, &redis.XReadArgs{
			Streams: []string{"room", "1"},
			Count:   1,
			Block:   0,
		}).Result()
		if !errors.Is(err, nil) {
			panic(err)
		}

		msg := streams[0].Messages[0]
		logger.Log("msg", msg)
		payload := pubsub.Payload{
			Topic: pubsub.Topic(msg.Values["topic"].(string)),
			Data:  msg.Values["data"],
		}

		for _, sub := range m.topics[payload.Topic] {
			sub.payload <- payload.Data
		}
	}
}
