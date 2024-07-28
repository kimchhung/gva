package pubsubredis

import (
	"context"
	"fmt"
	"sync"
	"testing"

	"github.com/gva/internal/pubsub"
	"github.com/redis/go-redis/v9"
)

func NewPubSub() pubsub.Pubsub {
	ctx := context.Background()
	redisC := redis.NewClient(
		&redis.Options{
			Addr:     "localhost:6379",
			Password: "123456",
		},
	)
	if err := redisC.Ping(ctx).Err(); err != nil {
		panic(err)
	}
	psub := NewRedisPubSub(redisC, "gql:subs")
	go psub.Listen(ctx)
	return psub
}

func TestHighLoadPubsub(t *testing.T) {
	ctx := context.Background()
	psub := NewPubSub()

	// Simulate high load by creating  large number of subscriptions
	numSubscriptions := 1

	for range numSubscriptions {
		r, err := psub.Sub(ctx, "testTopic")
		if err != nil {
			t.Fatalf("Failed to subscribe: %v", err)
		}

		go func() {
			for msg := range r.Payload() {
				if msg != "Test" {
					t.Errorf("Received unexpected message: %v", msg)
				}
			}
		}()
	}

	psub.Pub(ctx, "testTopic", "Test")
}

var Book = map[string]any{
	"ID":    "123456789",
	"Books": []string{"123", "asda1"},
}

func TestHighLoadPublishAndSub(t *testing.T) {
	ctx := context.TODO()

	publisher := NewPubSub()

	numSubscriptions := 5
	numPublishs := 1
	serverNode := 5

	wg := sync.WaitGroup{}
	wg.Add(numSubscriptions * numPublishs * serverNode)

	for i := range serverNode {
		serverSub := NewPubSub()
		topic := fmt.Sprintf("testTopic%v", i)
		for range numSubscriptions {

			r, err := serverSub.Sub(ctx, pubsub.Topic(topic))
			if err != nil {
				t.Fatalf("Failed to subscribe: %v", err)
			}

			go func() {
				for msg := range r.Payload() {
					if msg == nil {
						t.Errorf("Received unexpected message: %v", msg)
					}
					wg.Done()
				}
			}()
		}
	}

	for i := range serverNode {
		topic := fmt.Sprintf("testTopic%v", i)
		for range numPublishs {
			publisher.Pub(ctx, pubsub.Topic(topic), Book)
		}
	}

	wg.Wait()
}
