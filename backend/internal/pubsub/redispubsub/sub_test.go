package redispubsub

import (
	"context"
	"testing"

	"github.com/gva/app/database/schema/pxid"
	"github.com/gva/internal/pubsub"
	"github.com/redis/go-redis/v9"
)

func RedisC() *redis.Client {
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

	return redisC
}

// BenchmarkHighLoadPublish-16    	    1950	    540619 ns/op	     345 B/op	      12 allocs/op
// BenchmarkHighLoadPublish-16    	    1206	    931845 ns/op	     334 B/op	      10 allocs/op
func BenchmarkHighLoadPublish(b *testing.B) {
	rc := RedisC()
	idGen := func() string {
		return pxid.New("").String()
	}
	broker := NewBroker(rc, "test-gql", idGen)
	go broker.Receive(context.Background())
	pub := NewPublisher(rc, "test-gql")
	numSubscriptions := 1
	numPublishs := b.N

	topic := "test-topic1"
	ctx := context.Background()

	for range numSubscriptions {
		ch := make(chan string, 1)
		_, err := broker.Subscribe(topic, ch, func(_ string, payload pubsub.Payload) (any, error) {
			return string(payload.Data.(string)), nil
		})
		if err != nil {
			b.Errorf("error on sub %v", err)
		}
		go func() {
			msg := <-ch
			if msg != "test-value" {
				b.Errorf("wrong value")
			}
		}()
	}

	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	defer b.StopTimer()

	for range numPublishs {
		pub.Publish(ctx, topic, "test-value")
	}
}

// BenchmarkHighLoadSubsciptionMessage-16    	  438751	      3003 ns/op	     949 B/op	       7 allocs/op
// BenchmarkHighLoadSubsciptionMessage-16    	  361729	      4334 ns/op	     895 B/op	       7 allocs/op
//
//	298346	      4688 ns/op	     919 B/op	       8 allocs/op
func BenchmarkHighLoadSubsciptionMessage(b *testing.B) {
	rc := RedisC()
	idGen := func() string {
		return pxid.New("").String()
	}
	broker := NewBroker(rc, "test-gql", idGen)
	go broker.Receive(context.Background())
	pub := NewPublisher(rc, "test-gql")
	ctx := context.Background()

	numSubscriptions := b.N
	numPublishs := 1
	topic := "testTopic"

	b.ResetTimer()
	for range numSubscriptions {
		ch := make(chan string, 1)
		_, err := broker.Subscribe(topic, ch, func(subscriptionID string, payload pubsub.Payload) (interface{}, error) {
			return string(payload.Data.(string)), nil
		})
		if err != nil {
			b.Errorf("error on sub %v", err)
		}
		go func() {
			msg := <-ch
			if msg != "test-value" {
				b.Errorf("wrong value")
			}
		}()
	}

	for range numPublishs {
		msg := "test-value"
		pub.Publish(ctx, topic, msg)
	}
}
