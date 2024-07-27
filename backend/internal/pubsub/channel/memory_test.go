package pubsubchannel

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/gva/internal/pubsub"
)

func TestHighLoadPubsub(t *testing.T) {
	ctx := context.Background()
	psub := NewMemoryPubsub()

	// Simulate high load by creating  large number of subscriptions
	numSubscriptions := 1_000_000
	subChans := make([]pubsub.SubResult, numSubscriptions)

	for i := 0; i < numSubscriptions; i++ {
		r, err := psub.Sub(ctx, "testTopic")
		subChans[i] = r
		if err != nil {
			t.Fatalf("Failed to subscribe: %v", err)
		}

		go func(sc pubsub.SubResult) {
			msg := <-sc.Payload()
			if msg != "Test" {
				t.Errorf("Received unexpected message: %v", msg)
			}
		}(r)
	}

	psub.Pub(ctx, "testTopic", "Test")
}

func TestHighLoadPublishAndSub(t *testing.T) {
	ctx := context.Background()
	psub := NewMemoryPubsub()

	numSubscriptions := 168
	numPublishs := 178

	count := 0
	mu := sync.Mutex{}
	increaseCount := func() {
		mu.Lock()
		defer mu.Unlock()
		count += 1
		fmt.Println("count ?", count)
	}

	subs := make([]pubsub.SubResult, numSubscriptions)
	for i := range make([]int, numSubscriptions) {
		r, err := psub.Sub(ctx, "testTopic")
		if err != nil {
			t.Fatalf("Failed to subscribe: %v", err)
		}
		subs[i] = r
		go func(sc pubsub.SubResult, increaseCount func()) {
			for msg := range sc.Payload() {
				increaseCount()
				if msg != "Test" {
					t.Errorf("Received unexpected message: %v", msg)
				}
			}
		}(r, increaseCount)
	}

	for range numPublishs {
		psub.Pub(ctx, "testTopic", "Test")
	}

	time.Sleep(time.Second)
	for _, sub := range subs {
		sub.UnSub()
	}

	if count != numSubscriptions*numPublishs {
		t.Errorf("Expected %d messages, got %d", numSubscriptions*numPublishs, count)
	}
}
