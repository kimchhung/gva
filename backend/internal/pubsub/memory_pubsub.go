package pubsub

import (
	"context"
	"fmt"
	"sync"

	"github.com/pkg/errors"
)

var _ interface {
	Pubsub
} = (*memoryPubsub)(nil)

type MemorySubResult struct {
	topic     Topic
	payloadch <-chan Payload
	unsub     func() error
}

func NewMemorySubResult(
	topic Topic,
	payloadch chan Payload,
	unsub func() error,
) SubResult {
	return &MemorySubResult{
		topic:     topic,
		payloadch: payloadch,
		unsub:     unsub,
	}
}

func (sr *MemorySubResult) Topic() Topic {
	return sr.topic
}

func (sr *MemorySubResult) Payload() <-chan Payload {
	return sr.payloadch
}

func (sr *MemorySubResult) UnSub() error {
	return sr.unsub()
}

type memoryPubsub struct {
	topics map[Topic]map[*chan Payload]struct{}
	mu     sync.Mutex
}

func NewMemoryPubsub() *memoryPubsub {
	return &memoryPubsub{
		topics: make(map[Topic]map[*chan Payload]struct{}),
	}
}

func (b *memoryPubsub) Pub(ctx context.Context, topic Topic, p Payload) error {
	b.mu.Lock()
	defer b.mu.Unlock()

	ch, ok := b.topics[topic]
	if !ok {
		return ErrTopicNotFound
	}

	if len(ch) == 0 {
		fmt.Println("no subsciber")
		return nil
	}

	// Check for context cancellation before broadcasting
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		// Broadcast the payload to all subscribers
		for c := range ch {
			*c <- p
		}
	}

	return nil
}

func (b *memoryPubsub) Sub(ctx context.Context, topic Topic) (SubResult, error) {
	defer fmt.Println("subed")
	b.mu.Lock()
	defer b.mu.Unlock()

	ch := make(chan Payload, 1)
	pch := &ch
	if b.topics[topic] == nil {
		b.topics[topic] = map[*chan Payload]struct{}{}
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

func (b *memoryPubsub) UnSub(topic Topic, pch *chan Payload) error {
	fmt.Println("unsub")
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

func (b *memoryPubsub) Close() error {

	return nil
}
