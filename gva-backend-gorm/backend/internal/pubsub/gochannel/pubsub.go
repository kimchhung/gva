package gochannel

import "backend/internal/pubsub"

type Pubsub struct {
	pubsub.Publisher
	pubsub.Broker
}

func NewPubsub(publishChan chan pubsub.Payload, generateId func() string) *Pubsub {
	broker := NewBroker(publishChan, generateId)
	publisher := NewPublisher(publishChan)

	return &Pubsub{
		Publisher: publisher,
		Broker:    broker,
	}
}
