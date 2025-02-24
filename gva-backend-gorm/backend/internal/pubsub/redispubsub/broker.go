package redispubsub

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"sync"

	"backend/internal/pubsub"

	"github.com/redis/go-redis/v9"
)

var _ pubsub.Broker = (*redisBroker)(nil)

type redisBroker struct {
	client        redis.UniversalClient
	redisChannel  string
	subscriptions sync.Map // Topic => TopicSubscriptions
	generateId    func() string
}

// NewBroker new broker using redis
func NewBroker(client redis.UniversalClient, redisChannel string, generateId func() string) pubsub.Broker {
	return &redisBroker{
		client:       client,
		redisChannel: redisChannel,
		generateId:   generateId,
	}
}

func (b *redisBroker) Receive(ctx context.Context) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("Receive panic: %v", r)
		}
	}()

	rpubsub := b.client.Subscribe(ctx, b.redisChannel)
	_, err = rpubsub.Receive(ctx)
	if err != nil {
		return err
	}

	for msg := range rpubsub.Channel() {
		var payload pubsub.Payload
		err := json.Unmarshal([]byte(msg.Payload), &payload)
		if err != nil {
			return err
		}

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

	return rpubsub.Close()
}

func (b *redisBroker) Subscribe(channelName string, channel interface{}, handler pubsub.PayloadHandler) (pubsub.Unsubscriber, error) {
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
