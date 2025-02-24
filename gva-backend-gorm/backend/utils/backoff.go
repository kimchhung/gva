package utils

import (
	"math"
	"time"
)

type exponentialBackOff struct {
	min        time.Duration
	max        time.Duration
	multiplier float64
	attempt    uint
	maxAttempt uint
}

func NewExponentialBackOff(min, max time.Duration, multiplier float64, maxAttempt uint) *exponentialBackOff {
	return &exponentialBackOff{
		min:        min,
		max:        max,
		multiplier: multiplier,
		attempt:    0,
		maxAttempt: maxAttempt,
	}
}

func (b *exponentialBackOff) Reset() {
	b.attempt = 0
}

func (b *exponentialBackOff) Wait() {
	sleep := b.min
	if b.attempt > 0 {
		sleep = time.Duration(float64(b.min) * math.Pow(b.multiplier, float64(b.attempt)))
		if sleep > b.max {
			sleep = b.max
		}
	}

	time.Sleep(sleep)
	b.attempt++

	if b.attempt >= b.maxAttempt {
		b.attempt = b.maxAttempt
	}
}

func (b *exponentialBackOff) ShouldRetry() bool {
	return b.attempt < b.maxAttempt
}

func (b *exponentialBackOff) Attempt() uint {
	return b.attempt
}

func (b *exponentialBackOff) MaxAttempt() uint {
	return b.maxAttempt
}
