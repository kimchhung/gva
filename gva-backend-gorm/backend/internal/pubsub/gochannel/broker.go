package gochannel

import (
	"context"
	"fmt"
	"log"
	"reflect"
	"sync"

	"backend/internal/pubsub"
)

var _ pubsub.Broker = (*Broker)(nil)

type Broker struct {
	publishRequest chan pubsub.Payload
	subscriptions  sync.Map // Topic => TopicSubscriptions
	generateId     func() string
}

// NewBroker new broker using redis
func NewBroker(req chan pubsub.Payload, generateId func() string) pubsub.Broker {
	return &Broker{
		publishRequest: req,
		generateId:     generateId,
	}
}

func (b *Broker) Receive(ctx context.Context) error {
	for {
		select {

		case <-ctx.Done():
			return ctx.Err()
		default:
			payload := <-b.publishRequest

			chValue, exist := b.subscriptions.Load(payload.Topic)
			if !exist {
				continue
			}
			ch, ok := chValue.(*channelSubscriptions)
			if !ok {
				continue
			}

			ch.RLock()
			for id, subscription := range ch.subscriptions {
				message, err := subscription.handler(id, payload)
				if err != nil {
					log.Println(err)
					continue
				}

				subscription.channel.Send(reflect.ValueOf(message))
			}
			ch.RUnlock()
		}
	}
}

func (b *Broker) Subscribe(channelName string, channel interface{}, handler pubsub.PayloadHandler) (pubsub.Unsubscriber, error) {
	value := reflect.ValueOf(channel)
	if kind := value.Kind(); kind != reflect.Chan {
		return nil, fmt.Errorf("channel must be channel, but %v", kind)
	}

	chValue, ok := b.subscriptions.Load(channelName)
	if !ok {
		newChannel := &channelSubscriptions{subscriptions: make(map[string]*subscription)}
		chValue, _ = b.subscriptions.LoadOrStore(channelName, newChannel)
	}

	ch, _ := chValue.(*channelSubscriptions)
	id := b.generateId()
	return ch.append(channelName, value, id, handler), nil
}

type subscription struct {
	topic   string
	channel reflect.Value
	handler pubsub.PayloadHandler
}

type channelSubscriptions struct {
	sync.RWMutex
	subscriptions map[string]*subscription // subscription id => subscription
}

func (c *channelSubscriptions) append(topic string, channel reflect.Value, id string, handler pubsub.PayloadHandler) pubsub.Unsubscriber {
	c.Lock()
	defer c.Unlock()
	c.subscriptions[id] = &subscription{
		topic:   topic,
		channel: channel,
		handler: handler,
	}

	return pubsub.CloserFunc(func() error {
		c.remove(id)
		return nil
	})
}

func (c *channelSubscriptions) remove(id string) {
	c.Lock()
	defer c.Unlock()
	delete(c.subscriptions, id)
}
