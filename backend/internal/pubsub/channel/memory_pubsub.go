package pubsubchannel

import (
	"context"
	"errors"

	"github.com/gva/internal/pubsub"
)

type memoryPubsub struct {
	topics           map[pubsub.Topic]map[*chan pubsub.Data]struct{}
	addSubscriber    chan Subscriber
	removeSubscriber chan Subscriber
	publishPayload   chan pubsub.Payload
	isStarted        bool
}

type Subscriber struct {
	topic           pubsub.Topic
	subId           *chan pubsub.Data
	unsubscribeFunc func() error
}

func (ms Subscriber) Data() <-chan pubsub.Data {
	return *ms.subId
}

func (ms Subscriber) UnSub() error {
	return ms.unsubscribeFunc()
}

func NewMemoryPubsub() pubsub.Pubsub {
	m := &memoryPubsub{
		topics:           make(map[pubsub.Topic]map[*chan pubsub.Data]struct{}),
		addSubscriber:    make(chan Subscriber),
		removeSubscriber: make(chan Subscriber),
		publishPayload:   make(chan pubsub.Payload, 100),
	}
	return m
}

func (b *memoryPubsub) Pub(ctx context.Context, topic pubsub.Topic, data pubsub.Data) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		b.publishPayload <- pubsub.Payload{Topic: topic, Data: data}
	}
	return nil
}

func (b *memoryPubsub) Sub(ctx context.Context, topic pubsub.Topic) (pubsub.SubResult, error) {
	ch := make(chan pubsub.Data, 1)
	subscriber := Subscriber{
		topic: topic,
		subId: &ch,
	}

	subscriber.unsubscribeFunc = func() error {
		b.removeSubscriber <- subscriber
		return nil
	}

	b.addSubscriber <- subscriber
	return subscriber, nil
}

func (m *memoryPubsub) Listen(ctx context.Context) error {
	if m.isStarted {
		return errors.New("can't listen more than once")
	}

	m.isStarted = true
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()

		case sub := <-m.addSubscriber:
			if m.topics[sub.topic] == nil {
				m.topics[sub.topic] = map[*chan pubsub.Data]struct{}{}
			}

			m.topics[sub.topic][sub.subId] = struct{}{}

		case sub := <-m.removeSubscriber:
			subs, hasTopic := m.topics[sub.topic]
			if !hasTopic {
				continue
			}

			_, hasSub := subs[sub.subId]
			if !hasSub {
				continue
			}

			close(*sub.subId)
			delete(subs, sub.subId)

		case pub := <-m.publishPayload:
			subs, hasTopic := m.topics[pub.Topic]
			if !hasTopic {
				continue
			}

			for subId := range subs {
				*subId <- pub.Data
			}
		}
	}
}
