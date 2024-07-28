package stream

import (
	"context"
	"sync"
	"testing"

	"github.com/gva/internal/logger"
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
	numSubscriptions := 2

	wg := sync.WaitGroup{}
	wg.Add(numSubscriptions)
	for range numSubscriptions {
		r, err := psub.Sub(ctx, "testTopic")
		if err != nil {
			t.Fatalf("Failed to subscribe: %v", err)
		}

		go func() {
			for msg := range r.Data() {
				wg.Done()
				logger.Log(msg)
				if msg != "Test" {
					t.Errorf("Received unexpected message: %v", msg)
				}
			}
		}()
	}

	psub.Pub(ctx, "testTopic", []byte("Test"))
	wg.Wait()
}

var Book = map[string]any{
	"ID":    "123456789",
	"Books": []string{"123", "asda1"},
}

// func TestHighLoadPublishAndSub(t *testing.T) {
// 	ctx := context.TODO()

// 	publisher := NewPubSub()
// 	go publisher.Listen(ctx)

// 	numSubscriptions := 100
// 	numPublishs := 100
// 	serverNode := 2

// 	wg := sync.WaitGroup{}
// 	wg.Add(numSubscriptions * numPublishs * serverNode)

// 	for i := range serverNode {
// 		serverSub := NewPubSub()
// 		go publisher.Listen(ctx)
// 		topic := fmt.Sprintf("testTopic%v", i)
// 		for range numSubscriptions {
// 			r, err := serverSub.Sub(ctx, pubsub.Topic(topic))
// 			if err != nil {
// 				t.Fatalf("Failed to subscribe: %v", err)
// 			}

// 			go func() {
// 				for msg := range r.Data() {
// 					if msg == nil {
// 						t.Errorf("Received unexpected message: %v", msg)
// 					}
// 					wg.Done()
// 				}
// 			}()
// 		}
// 	}

// 	for i := range serverNode {
// 		topic := fmt.Sprintf("testTopic%v", i)
// 		for range numPublishs {
// 			publisher.Pub(ctx, pubsub.Topic(topic), []byte("test"))
// 		}
// 	}

// 	wg.Wait()
// }
