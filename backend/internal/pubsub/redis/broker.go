package pubsubredis

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/gva/app/database/schema/xid"
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

type PublishRequest struct {
	Topic   pubsub.Topic   `json:"topic"`
	Payload pubsub.Payload `json:"payload"`
}

type Subscriber struct {
	topic           pubsub.Topic
	subId           xid.ID
	payloadChan     chan pubsub.Payload
	unsubscribeFunc func() error
}

type topicSubscribers map[xid.ID]Subscriber

func (ms Subscriber) Payload() <-chan pubsub.Payload {
	return ms.payloadChan
}

func (ms Subscriber) UnSub() error {
	return ms.unsubscribeFunc()
}

func (b *redisBroker) Sub(ctx context.Context, topic pubsub.Topic) (pubsub.SubResult, error) {
	ch := make(chan pubsub.Payload, 1)
	subscriber := Subscriber{
		topic:       topic,
		subId:       xid.MustNew(fmt.Sprintln(topic)),
		payloadChan: ch,
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
	pubsub := m.client.Subscribe(ctx, m.redisChannel)
	_, err := pubsub.Receive(ctx)
	if err != nil {
		return err
	}

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
				close(sub.payloadChan)
			}
		}
	}()

	for msgFromPublish := range pubsub.Channel(m.opts...) {
		var msgdata PublishRequest
		err := json.Unmarshal([]byte(msgFromPublish.Payload), &msgdata)
		if err != nil {
			panic(fmt.Errorf("invalid json.Unmarshal %v", err))
		}

		subs, ok := m.topics[msgdata.Topic]
		if !ok {
			continue
		}

		for subId := range subs {
			sub := subs[subId]
			sub.payloadChan <- msgdata.Payload
		}
	}

	return pubsub.Close()
}
