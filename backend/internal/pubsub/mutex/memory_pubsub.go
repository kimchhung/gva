package pubsubmutex

import (
	"context"
	"fmt"
	"sync"

	"github.com/gva/internal/pubsub"
	"github.com/pkg/errors"
)

var _ interface {
	pubsub.Pubsub
} = (*memoryPubsub)(nil)

type MemorySubResult struct {
	topic     pubsub.Topic
	payloadch chan pubsub.Payload
	unsub     func() error
}

func NewMemorySubResult(
	topic pubsub.Topic,
	payloadch chan pubsub.Payload,
	unsub func() error,
) pubsub.SubResult {
	return &MemorySubResult{
		topic:     topic,
		payloadch: payloadch,
		unsub:     unsub,
	}
}

func (sr *MemorySubResult) Topic() pubsub.Topic {
	return sr.topic
}

func (sr *MemorySubResult) Payload() <-chan pubsub.Payload {
	return sr.payloadch
}

func (sr *MemorySubResult) UnSub() error {
	return sr.unsub()
}

type memoryPubsub struct {
	topics map[pubsub.Topic]map[*chan pubsub.Payload]struct{}
	mu     sync.Mutex
}

func NewMemoryPubsub() *memoryPubsub {
	return &memoryPubsub{
		topics: make(map[pubsub.Topic]map[*chan pubsub.Payload]struct{}),
	}
}

func (b *memoryPubsub) Pub(ctx context.Context, topic pubsub.Topic, p pubsub.Payload) error {
	b.mu.Lock()
	defer b.mu.Unlock()

	ch, ok := b.topics[topic]
	if !ok {
		return pubsub.ErrTopicNotFound
	}

	if len(ch) == 0 {
		fmt.Println("no subsciber")
		return nil
	}

	for c := range ch {
		*c <- p
	}

	return nil
}

func (b *memoryPubsub) Sub(ctx context.Context, topic pubsub.Topic) (pubsub.SubResult, error) {
	b.mu.Lock()
	defer b.mu.Unlock()

	ch := make(chan pubsub.Payload)
	pch := &ch
	if b.topics[topic] == nil {
		b.topics[topic] = map[*chan pubsub.Payload]struct{}{}
	}

	b.topics[topic][pch] = struct{}{}
	sr := NewMemorySubResult(
		topic,
		ch,
		func() error {
			return b.UnSub(topic, pch)
		},
	)
	return sr, nil
}

func (b *memoryPubsub) UnSub(topic pubsub.Topic, pch *chan pubsub.Payload) error {
	b.mu.Lock()
	defer b.mu.Unlock()

	chs, ok := b.topics[topic]
	if !ok || len(chs) == 0 {
		return errors.Errorf("cannot find topic %s", topic)
	}

	if _, ok := chs[pch]; !ok {
		return errors.Errorf("cannot find channel in topic %s", topic)
	}

	close(*pch)
	delete(chs, pch)
	return nil
}
