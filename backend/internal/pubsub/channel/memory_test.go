package pubsubchannel

import (
	"context"
	"sync"
	"testing"
)

func TestHighLoadPubsub(t *testing.T) {
	ctx := context.Background()
	psub := NewMemoryPubsub()
	go psub.Listen(ctx)

	// Simulate high load by creating  large number of subscriptions
	numSubscriptions := 1_000_000
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

func TestHighLoadPublishAndSub(t *testing.T) {
	ctx := context.Background()
	psub := NewMemoryPubsub()
	go psub.Listen(ctx)

	numSubscriptions := 1
	numPublishs := 1_000_000

	wg := sync.WaitGroup{}

	for range numSubscriptions {
		r, err := psub.Sub(ctx, "testTopic")
		if err != nil {
			t.Fatalf("Failed to subscribe: %v", err)
		}

		go func() {
			for msg := range r.Payload() {
				wg.Done()
				if msg != "Test" {
					t.Errorf("Received unexpected message: %v", msg)
				}
			}
		}()
	}

	for range numPublishs {
		wg.Add(1)
		psub.Pub(ctx, "testTopic", "Test")
	}
	wg.Wait()
}
