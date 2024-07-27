package pubsubchannel

import (
	"context"
	"log"

	"github.com/gva/internal/pubsub"
)

type memoryPubsub struct {
	topics           map[pubsub.Topic]map[*chan pubsub.Payload]struct{}
	addSubscriber    chan Subscriber
	removeSubscriber chan Subscriber
	publishChan      chan PublishRequest
}

type PublishRequest struct {
	topic   pubsub.Topic
	payload pubsub.Payload
}

type Subscriber struct {
	topic           pubsub.Topic
	subId           *chan pubsub.Payload
	unsubscribeFunc func() error
}

func (ms Subscriber) Payload() <-chan pubsub.Payload {
	return *ms.subId
}

func (ms Subscriber) UnSub() error {
	return ms.unsubscribeFunc()
}

func NewMemoryPubsub() pubsub.Pubsub {
	m := &memoryPubsub{
		topics:           make(map[pubsub.Topic]map[*chan pubsub.Payload]struct{}),
		addSubscriber:    make(chan Subscriber),
		removeSubscriber: make(chan Subscriber),
		publishChan:      make(chan PublishRequest),
	}

	go func(m *memoryPubsub) {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("Recovered from panic: %v\n", r)
			}
		}()

		for {
			select {
			case sub := <-m.addSubscriber:
				if m.topics[sub.topic] == nil {
					m.topics[sub.topic] = map[*chan pubsub.Payload]struct{}{}
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

				delete(subs, sub.subId)
				if len(subs) == 0 {
					delete(m.topics, sub.topic)
				}

			case pub := <-m.publishChan:
				subs, hasTopic := m.topics[pub.topic]
				if !hasTopic {
					continue
				}

				for subId := range subs {
					sub := *subId
					sub <- pub.payload
				}
			}
		}
	}(m)

	return m
}

func (b *memoryPubsub) Pub(ctx context.Context, topic pubsub.Topic, p pubsub.Payload) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		b.publishChan <- PublishRequest{topic: topic, payload: p}
	}
	return nil
}

func (b *memoryPubsub) Sub(ctx context.Context, topic pubsub.Topic) (pubsub.SubResult, error) {
	ch := make(chan pubsub.Payload, 1000)
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
