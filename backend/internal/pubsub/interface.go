package pubsub

import (
	"context"
	"errors"
)

type Data any
type Topic string

type Payload struct {
	Topic Topic
	Data  Data
}

// SubResult represents the result of subscribing to a topic.
//
// Methods allow the subscriber to retrieve the topic they are subscribed to,
// receive payloads, and unsubscribe from the topic.
type SubResult interface {
	// Payload returns a channel that receives payloads published to this topic.
	Data() <-chan Data

	// UnSub unsubscribes from the topic, stopping further payload reception.
	UnSub() error
}

type Broker interface {
	Sub(ctx context.Context, topic Topic) (SubResult, error)
	Listen(ctx context.Context) error
}

type Publisher interface {
	Pub(ctx context.Context, topic Topic, data Data) error
}

type Pubsub interface {
	Publisher
	Broker
}

var (
	ErrContextCancel = errors.New("context cancel")
	ErrTopicNotFound = errors.New("topic not found")
)
