package gqlgen

import (
	"context"
	"fmt"
	"testing"

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

// BenchmarkHighLoadPublish-16    	    2176	    537807 ns/op	     337 B/op	      11 allocs/op
func BenchmarkHighLoadPublish(b *testing.B) {
	rc := RedisC()
	broker := NewRedisBroker(rc, "test-gql")
	pub := NewRedisPublisher(rc, "test-gql")
	numSubscriptions := 1
	numPublishs := b.N

	topic := "test-topic1"

	for range numSubscriptions {
		ch := make(chan string)
		_, err := broker.Subscribe(topic, ch, func(subscriptionID string, payload Payload) (interface{}, error) {
			return string(payload.Data), nil
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
		pub.Publish(topic, []byte("test-value"))
	}
}

// BenchmarkHighLoadSubsciptionMessage-16    	  286711	      4172 ns/op	    1127 B/op	       8 allocs/op
// PASS
// BenchmarkHighLoadSubsciptionMessage-16    	  576044	      2717 ns/op	     924 B/op	       7 allocs/op
func BenchmarkHighLoadSubsciptionMessage(b *testing.B) {
	rc := RedisC()
	broker := NewRedisBroker(rc, "test-gql")
	pub := NewRedisPublisher(rc, "test-gql")

	numSubscriptions := b.N
	numPublishs := 1
	topic := "testTopic"

	b.ResetTimer()
	for range numSubscriptions {
		ch := make(chan string)
		_, err := broker.Subscribe(topic, ch, func(subscriptionID string, payload Payload) (interface{}, error) {
			return string(payload.Data), nil
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

	for i := range numPublishs {
		msg := fmt.Sprintf("pub-%d", i+1)
		pub.Publish(topic, []byte(msg))
	}
}
