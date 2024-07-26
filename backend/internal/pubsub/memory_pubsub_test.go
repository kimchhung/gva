package pubsub

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMultipleSubscriptions(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	pubsubInstance := NewMemoryPubsub()

	// Subscribe twice to the same topic
	subscription1, err := pubsubInstance.Sub(ctx, "test")
	assert.NoError(t, err)
	subscription2, err := pubsubInstance.Sub(ctx, "test")
	assert.NoError(t, err)

	// Wait for both subscriptions to receive the message
	received1 := false
	received2 := false
	go func() {
		for msg := range subscription1.Payload() {
			assert.Equal(t, "message", msg)
			received1 = true
			break // Break after receiving once
		}
	}()

	go func() {
		for msg := range subscription2.Payload() {
			assert.Equal(t, "message", msg)
			received2 = true
			break // Break after receiving once
		}
	}()

	// Publish a message
	pubsubInstance.Pub(context.Background(), "test", "message")

	// Give some time for the messages to be processed
	time.Sleep(time.Second)

	assert.True(t, received1)
	assert.True(t, received2)
}

func TestMessageDelivery(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	pubsub := NewMemoryPubsub()
	sub, _ := pubsub.Sub(context.Background(), "testTopic")

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		for msg := range sub.Payload() {
			fmt.Printf("Received: %v\n", msg)
		}
	}()

	err := pubsub.Pub(ctx, "testTopic", "Hello, World!")
	if err != nil {
		t.Fatalf("Failed to publish message: %v", err)
	}
	sub.UnSub()

	wg.Wait()
}

func TestHighLoadPubsub(t *testing.T) {
	ctx := context.Background()
	pubsub := NewMemoryPubsub()

	// Simulate high load by creating  large number of subscriptions
	numSubscriptions := 1_000_000
	subChans := make([]SubResult, numSubscriptions)

	for i := 0; i < numSubscriptions; i++ {
		r, err := pubsub.Sub(ctx, "testTopic")
		subChans[i] = r
		if err != nil {
			t.Fatalf("Failed to subscribe: %v", err)
		}

		go func(sc SubResult) {
			msg := <-sc.Payload()
			if msg != "Test" {
				t.Errorf("Received unexpected message: %v", msg)
			}
		}(r)
	}

	pubsub.Pub(ctx, "testTopic", "Test")
}
