package pubsub

import (
	"context"
	"io"
	"log"
	"reflect"
)

type Pubsub interface {
	Broker
	Publisher
}

type Broker interface {
	// Subscribe subscribes messages by topic. It will handle payload, then send handled message to subscription channel.
	Subscribe(topic string, channel interface{}, payloadHandler PayloadHandler) (Unsubscriber, error)
	// Receive blocks to receive data from Publisher
	Receive(ctx context.Context) error
}

// Publisher ...
type Publisher interface {
	// Publish publishes payload data to channel
	Publish(ctx context.Context, topic string, data any) error
}

type Payload struct {
	Topic string
	Data  any
}

// PayloadHandler represents function that returns message to subscription channel from payload
type PayloadHandler func(sid string, payload Payload) (any, error)

// Unsubscriber is the interface the wrap the Close method to unsubscribe
type Unsubscriber io.Closer

// CloserFunc is an adapter to allow use of ordinary functions as io.Closer
type CloserFunc func() error

// Close calls c
func (c CloserFunc) Close() error {
	return c()
}

func CloseSubscription(ctx context.Context, u Unsubscriber, channel interface{}) {
	<-ctx.Done()
	u.Close()

	cValue := reflect.ValueOf(channel)
	if cValue.Kind() == reflect.Chan {
		cValue.Close()
	} else {
		log.Printf("type of channel is not channel but %s", cValue.Kind())
	}
}
