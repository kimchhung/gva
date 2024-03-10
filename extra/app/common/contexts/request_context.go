package contexts

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

type (
	RequestContextKey    struct{}
	RequestContextOption func(*RequestContext)
)

type RequestContext struct {
	context.Context

	startTime time.Time
	endTime   time.Time
}

// a context help handling error
func NewRequestContext() fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx := &RequestContext{}
		ctx.Context = context.WithValue(c.UserContext(), RequestContextKey{}, ctx)
		ctx.startTime = time.Now()

		defer func() {
			ctx.endTime = time.Now()
			log.Info().
				Str("path", c.Route().Path).
				Str("duration", fmt.Sprintf("%v", ctx.endTime.Sub(ctx.startTime))).
				Msg("Request")
		}()

		c.SetUserContext(ctx)
		return c.Next()
	}
}

func GetRequestContext(ctx context.Context) (*RequestContext, error) {
	v, ok := ctx.(*RequestContext)
	if ok {
		return v, nil
	}

	v, ok = ctx.Value(RequestContextKey{}).(*RequestContext)
	if ok {
		return v, nil
	}

	return nil, errors.New("context does not contain RequestContext")
}

func MustRequestContext(ctx context.Context) *RequestContext {
	actx, err := GetRequestContext(ctx)
	if err != nil {
		panic(err)
	}

	return actx
}

func StartTime(ctx context.Context) time.Time {
	return MustRequestContext(ctx).startTime
}
