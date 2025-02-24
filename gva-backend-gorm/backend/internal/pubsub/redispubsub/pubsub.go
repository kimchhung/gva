package redispubsub

import (
	"backend/internal/pubsub"

	"github.com/redis/go-redis/v9"
)

type Pubsub struct {
	pubsub.Publisher
	pubsub.Broker
}

func NewPubsub(client redis.UniversalClient, channel string, generateId func() string) *Pubsub {
	broker := NewBroker(client, channel, generateId)
	publisher := NewPublisher(client, channel)

	return &Pubsub{
		Publisher: publisher,
		Broker:    broker,
	}
}
