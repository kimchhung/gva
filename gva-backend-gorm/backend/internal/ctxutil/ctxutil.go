package ctxutil

import (
	"context"
	"fmt"
)

// ctxKey is a context key for storing values in context.
type ctxKey string

// use `package.name` as key
func CtxKey[T comparable](v T) ctxKey {
	return ctxKey(fmt.Sprintf("%T", v))
}

func With[T comparable](ctx context.Context, value T) context.Context {
	return context.WithValue(ctx, CtxKey(value), value)
}

func Add(ctx context.Context, structOnly ...any) context.Context {
	for _, v := range structOnly {
		ctx = With(ctx, v)
	}

	return ctx
}

func Value[T comparable](ctx context.Context) (T, error) {
	var zero T
	key := CtxKey(zero)

	v, ok := ctx.Value(CtxKey(zero)).(T)
	if !ok {
		return zero, fmt.Errorf("%s not found in context", key)
	}

	return v, nil
}

func MustValue[T comparable](ctx context.Context) T {
	v, err := Value[T](ctx)
	if err != nil {
		panic(err)
	}

	return v
}
