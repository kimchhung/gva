package gochannel

import (
	"context"
	"testing"

	"github.com/gva/app/database/schema/pxid"
	"github.com/gva/internal/logging"
	"github.com/gva/internal/pubsub"
)

// BenchmarkHighLoadPublish-16    	   10000	    100183 ns/op	      40 B/op	       2 allocs/op
// BenchmarkHighLoadPublish-16    	   10000	    100297 ns/op	      40 B/op	       2 allocs/op
func BenchmarkHighLoadPublish(b *testing.B) {
	idGen := func() string {
		return pxid.New("").String()
	}
	ctx := context.Background()
	pubReq := make(chan pubsub.Payload)
	broker := NewBroker(pubReq, idGen)
	go broker.Receive(ctx)
	pub := NewPublisher(pubReq)
	numSubscriptions := 1
	numPublishs := b.N

	topic := "test-topic1"

	for range numSubscriptions {
		ch := make(chan string, 1)
		_, err := broker.Subscribe(topic, ch, func(_ string, payload pubsub.Payload) (any, error) {
			return string(payload.Data.([]byte)), nil
		})
		if err != nil {
			b.Errorf("error on sub %v", err)
		}
		go func() {
			msg := <-ch
			logging.Log(msg)
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
		pub.Publish(ctx, topic, []byte("test-value"))
	}
}

// BenchmarkHighLoadSubsciptionMessage-16    	  326073	      4634 ns/op	     890 B/op	       6 allocs/op
func BenchmarkHighLoadSubsciptionMessage(b *testing.B) {
	idGen := func() string {
		return pxid.New("").String()
	}
	pubReq := make(chan pubsub.Payload)
	broker := NewBroker(pubReq, idGen)
	go broker.Receive(context.TODO())
	pub := NewPublisher(pubReq)
	ctx := context.Background()

	numSubscriptions := b.N
	numPublishs := 1
	topic := "testTopic"

	b.ResetTimer()
	for range numSubscriptions {
		ch := make(chan string, 1)
		_, err := broker.Subscribe(topic, ch, func(subscriptionID string, payload pubsub.Payload) (interface{}, error) {
			return string(payload.Data.([]byte)), nil
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
		pub.Publish(ctx, topic, []byte("test-value"))
	}
}
