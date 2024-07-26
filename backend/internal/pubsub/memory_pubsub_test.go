package pubsub

import (
	"context"
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
