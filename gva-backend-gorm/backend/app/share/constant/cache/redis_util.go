package cache

import (
	"backend/core/utils/json"
	"context"

	"github.com/redis/go-redis/v9"
)

// support generic types
func OnSubscribe[T any](ctx context.Context, client redis.UniversalClient, channel string, callback func(data T, err error)) (err error) {
	subscriber := client.Subscribe(ctx, channel)

	defer func() {
		if err = subscriber.Unsubscribe(ctx, channel); err != nil {
			return
		}

		if err := subscriber.Close(); err != nil {
			return
		}
	}()

	ch := subscriber.Channel()

	for {
		select {
		case msg := <-ch:
			var data T
			err := json.JSON(msg.Payload).Out(&data)
			callback(data, err)

		case <-ctx.Done():
			return
		}
	}
}
