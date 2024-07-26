package pubsub

import (
	"context"
	"errors"
)

type Payload any
type Topic any

// SubResult represents the result of subscribing to a topic.
//
// Methods allow the subscriber to retrieve the topic they are subscribed to,
// receive payloads, and unsubscribe from the topic.
type SubResult interface {
	// Payload returns a channel that receives payloads published to this topic.
	Payload() <-chan Payload

	// UnSub unsubscribes from the topic, stopping further payload reception.
	UnSub() error
}

type Pubsub interface {
	Sub(ctx context.Context, topic Topic) (SubResult, error)
	Pub(ctx context.Context, topic Topic, m Payload) error
}

var (
	ErrContextCancel = errors.New("context cancel")
	ErrTopicNotFound = errors.New("topic not found")
)
