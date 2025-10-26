package corecontext

import (
	"context"
	"fmt"
)

func Set(ctx context.Context, key Key, value any) {
	MustRequestContext(ctx).Set(key, value)
}

func SetOrGet[T any](ctx context.Context, key Key, value T) (T, bool) {
	v, loaded := MustRequestContext(ctx).GetOrSet(key, value)
	return v.(T), loaded
}

func Get[T any](ctx context.Context, key Key) (T, error) {
	v, ok := MustRequestContext(ctx).Get(key)
	if !ok {
		var zero T
		return zero, fmt.Errorf("key %v not found in context", key)
	}

	expectedType, ok := v.(T)
	if !ok {
		return expectedType, fmt.Errorf("invalid type, expected %T, got %T", v, expectedType)
	}

	return expectedType, nil
}

func Must[T any](ctx context.Context, key Key) T {
	v, err := Get[T](ctx, key)
	if err != nil {
		panic(err)
	}

	return v
}
