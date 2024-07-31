package pubsubredis

import (
	"context"
	"fmt"
	"sync"
	"testing"

	"github.com/gva/internal/pubsub"
	"github.com/redis/go-redis/v9"
)

func NewPubSub(channelNumber int) pubsub.Pubsub {
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

	channels := make([]string, channelNumber)
	for i := range channelNumber {
		channels[i] = fmt.Sprintf("channel:-%v", i)
	}
	psub := NewRedisPubSub(redisC, channels)
	go psub.Listen(ctx)
	return psub
}

func TestNonDuplicateMsg(t *testing.T) {
	ctx := context.Background()
	psub := NewPubSub(1)

	numSubscriptions := 10
	numPublishs := 3 // Use b.N to automatically adjust the number of iterations based on time

	wg := sync.WaitGroup{}
	wg.Add(numPublishs * numSubscriptions)
	messageCounts := make([]int, numSubscriptions)

	for i := 0; i < numSubscriptions; i++ {
		subscriptionChan, err := psub.Sub(ctx, "testTopic")
		if err != nil {
			t.Fatal(err)
		}

		go func(i int) {
			for msg := range subscriptionChan.Data() {
				fmt.Println("<-receive-", msg)
				if msg == "" {
					t.Error("Received unexpected message:", msg)
				}
				wg.Done()
				messageCounts[i]++
			}
		}(i)

	}

	for i := range numPublishs {
		msg := fmt.Sprintf("pub-%d", i+1)
		fmt.Println("\n send->", msg)
		psub.Pub(ctx, "testTopic", msg)
	}

	wg.Wait()
	actualTotal := 0
	for _, count := range messageCounts {
		actualTotal += count
	}
	if expectedCount := numPublishs * numSubscriptions; actualTotal != expectedCount {
		t.Errorf("invalid count expect %v, actual %v", expectedCount, actualTotal)
	}
}

/*
cpu: AMD Ryzen 7 6800U with Radeon Graphics
BenchmarkHighLoadPublish-16    	    1220	    984102 ns/op	    2338 B/op	      71 allocs/op
*/
func BenchmarkHighLoadPublish(b *testing.B) {
	ctx := context.Background()
	psub := NewPubSub(1)

	numSubscriptions := 1
	numPublishs := b.N

	for range numSubscriptions {
		subscriptionChan, err := psub.Sub(ctx, "testTopic")
		if err != nil {
			b.Fatal(err)
		}

		go func() {
			for msg := range subscriptionChan.Data() {
				if msg == "" {
					b.Error("Received unexpected message:", msg)
				}
			}
		}()
	}

	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	defer b.StopTimer()

	for i := range numPublishs {
		msg := fmt.Sprintf("pub-%d", i+1)
		psub.Pub(ctx, "testTopic", msg)
	}
}

/*
cpu: AMD Ryzen 7 6800U with Radeon Graphics
BenchmarkHighLoadSubsciptionMessage-16    	  135996	     10642 ns/op	    2938 B/op	       9 allocs/op
*/
func BenchmarkHighLoadSubsciptionMessage(b *testing.B) {
	ctx := context.Background()
	psub := NewPubSub(1)

	numSubscriptions := b.N
	numPublishs := 1

	b.ResetTimer()
	for range numSubscriptions {
		subscriptionChan, err := psub.Sub(ctx, "testTopic")
		if err != nil {
			b.Fatal(err)
		}

		go func() {
			msg := <-subscriptionChan.Data()
			if msg == "" {
				b.Error("Received unexpected message:", msg)
			}
		}()
	}

	for i := range numPublishs {
		msg := fmt.Sprintf("pub-%d", i+1)
		psub.Pub(ctx, "testTopic", msg)
	}
}

// func TestHighLoadManyServer(b *testing.B) {
// 	b.ReportAllocs()
// 	b.ResetTimer()
// 	ctx := context.TODO()
// 	publisher := NewPubSub()

// 	numSubscriptions := b.N
// 	numPublishs := 1
// 	serverNode := 1

// 	wg := sync.WaitGroup{}
// 	wg.Add(numSubscriptions * numPublishs * serverNode)

// 	for i := range serverNode {
// 		serverSub := NewPubSub()

// 		topic := fmt.Sprintf("testTopic%v", i)
// 		for range numSubscriptions {

// 			r, err := serverSub.Sub(ctx, pubsub.Topic(topic))
// 			if err != nil {
// 				b.Fatalf("Failed to subscribe: %v", err)
// 			}

// 			go func() {
// 				for msg := range r.Data() {
// 					if msg == nil {
// 						b.Errorf("Received unexpected message: %v", msg)
// 					}
// 					wg.Done()
// 				}
// 			}()
// 		}
// 	}

// 	for i := range serverNode {
// 		topic := fmt.Sprintf("testTopic%v", i)
// 		for range numPublishs {
// 			go publisher.Pub(ctx, pubsub.Topic(topic), []byte("test"))
// 		}
// 	}

// 	wg.Wait()
// }
