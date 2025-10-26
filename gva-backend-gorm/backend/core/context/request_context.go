package corecontext

import (
	"context"
	"errors"
	"sync"
)

type (
	Key string
)

var (
	RequestContextKey Key = "req_context"
	ErrNotFound           = errors.New("request context not found")
)

type RequestContext struct {
	paylaod sync.Map
}

func WithRequestContext(ctx context.Context) context.Context {
	reqctx := &RequestContext{
		paylaod: sync.Map{},
	}

	return context.WithValue(ctx, RequestContextKey, reqctx)
}

func (rc *RequestContext) Set(key Key, value any) {
	rc.paylaod.Store(key, value)
}

func (rc *RequestContext) Get(key Key) (v any, loaded bool) {
	return rc.paylaod.Load(key)
}

func (rc *RequestContext) GetOrSet(key Key, value any) (v any, loaded bool) {
	return rc.paylaod.LoadOrStore(key, value)
}

func GetRequestContext(ctx context.Context) (*RequestContext, error) {
	v, ok := ctx.Value(RequestContextKey).(*RequestContext)
	if ok {
		return v, nil
	}

	return nil, ErrNotFound
}

func MustRequestContext(ctx context.Context) *RequestContext {
	actx, err := GetRequestContext(ctx)
	if err != nil {
		panic(err)
	}

	return actx
}
