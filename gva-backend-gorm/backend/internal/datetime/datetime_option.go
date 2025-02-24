package datetime

import (
	"context"
	"time"

	"github.com/nleeper/goment"
)

// get now from context or new now
func WithContext(ctx context.Context) Option {
	return func(config *DateConfig) {
		config.ctx = &ctx
	}
}

func WithLocation(loc *time.Location) Option {
	return func(config *DateConfig) {
		config.location = loc
	}
}

func WithCurrentTime(currentTime *goment.Goment) Option {
	return func(config *DateConfig) {
		config.currentTime = currentTime
	}
}
