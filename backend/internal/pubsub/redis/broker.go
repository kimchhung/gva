package pubsubredis

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/gva/app/database/schema/pxid"
	"github.com/gva/internal/pubsub"
	"github.com/redis/go-redis/v9"
)

var _ pubsub.Broker = (*redisBroker)(nil)

type redisBroker struct {
	client       redis.UniversalClient
	redisChannel []string
	isStarted    bool

	topics           map[pubsub.Topic]topicSubscribers // topicName => topicSubscribers
	addSubscriber    chan Subscriber
	removeSubscriber chan Subscriber

	opts []redis.ChannelOption
}

func NewBroker(client redis.UniversalClient, channels []string, opts ...redis.ChannelOption) *redisBroker {
	return &redisBroker{
		client:           client,
		redisChannel:     channels,
		topics:           map[pubsub.Topic]topicSubscribers{},
		addSubscriber:    make(chan Subscriber),
		removeSubscriber: make(chan Subscriber),
		opts:             opts,
	}
}

type Subscriber struct {
	topic           pubsub.Topic
	subId           pxid.ID
	payload         chan pubsub.Data
	unsubscribeFunc func() error
}

type topicSubscribers map[pxid.ID]Subscriber

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
		subId:   pxid.New("sub"),
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
	rpsub := m.client.Subscribe(ctx, m.redisChannel...)
	_, err := rpsub.Receive(ctx)
	if err != nil {
		return err
	}

	go func() {
		defer close(m.addSubscriber)
		defer close(m.removeSubscriber)

		for {
			select {
			case <-ctx.Done():
				return

			case sub := <-m.addSubscriber:
				_, ok := m.topics[sub.topic]
				if !ok {
					m.topics = map[pubsub.Topic]topicSubscribers{
						sub.topic: {sub.subId: sub},
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

	for msg := range rpsub.Channel(m.opts...) {
		m.processMessage(msg)
	}

	return rpsub.Close()
}

func (m *redisBroker) processMessage(msg *redis.Message) {
	var msgdata struct {
		Topic pubsub.Topic
		Data  []byte
	}

	err := json.Unmarshal([]byte(msg.Payload), &msgdata)
	if err != nil {
		panic(fmt.Errorf("invalid json.Unmarshal %v", err))
	}

	subs, ok := m.topics[msgdata.Topic]
	if !ok {
		return
	}

	for subId := range subs {
		sub, ok := subs[subId]
		if !ok {
			continue
		}

		sub.payload <- string(msgdata.Data)
	}
}
