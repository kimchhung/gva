package contexts

import (
	"context"
	"errors"
	"fmt"
	"runtime/debug"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	rerror "github.com/kimchhung/gva/extra/internal/response/error"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var IsProduction bool

type logFields struct {
	ID        string
	RemoteIP  string
	Host      string
	Method    string
	Path      string
	Protocol  string
	httpCode  int
	ErrorCode int
	Latency   time.Duration
	Error     error
	Stack     []byte
}

func (lf *logFields) MarshalZerologObject(e *zerolog.Event) {
	e.
		Str("id", lf.ID).
		Str("remoteIp", lf.RemoteIP).
		Str("host", lf.Host).
		Str("method", lf.Method).
		Str("path", lf.Path).
		Str("protocol", lf.Protocol).
		Int("httpCode", lf.httpCode).
		Int("errorCode", lf.ErrorCode).
		Int64("latencyMs", lf.Latency.Milliseconds()).
		Str("tag", "request")

	if lf.Error != nil {
		e.Err(lf.Error)
	}

	if lf.Stack != nil {
		e.Bytes("stack", lf.Stack)
	}
}

type (
	RequestContextKey    struct{}
	RequestContextOption func(*RequestContext)
)

type RequestContext struct {
	context.Context

	startTime time.Time
	endTime   time.Time

	LogFields *logFields
}

func defaultFields(c *fiber.Ctx) *logFields {
	rid := c.Get(fiber.HeaderXRequestID)
	if rid == "" {
		rid = uuid.New().String()
		c.Set(fiber.HeaderXRequestID, rid)
	}

	fields := &logFields{
		ID:       rid,
		RemoteIP: c.IP(),
		Method:   c.Method(),
		Host:     c.Hostname(),
		Path:     c.Path(),
		Protocol: c.Protocol(),
	}

	return fields
}

// a context help handling error
func NewRequestContext() fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx := &RequestContext{}
		ctx.Context = context.WithValue(c.UserContext(), RequestContextKey{}, ctx)
		ctx.startTime = time.Now()
		ctx.LogFields = defaultFields(c)

		c.SetUserContext(ctx)
		var err error

		defer func() {
			rvr := recover()

			if rvr != nil {
				perr, ok := rvr.(error)
				if !ok {
					err = fmt.Errorf("%v", rvr)
					ctx.LogFields.Stack = debug.Stack()
					ctx.LogFields.Error = err
				} else {
					err = perr
				}
			}

			a, _ := rerror.ParseError(c, err)

			ctx.endTime = time.Now()
			ctx.LogFields.ErrorCode = a.HttpCode
			ctx.LogFields.httpCode = c.Response().StatusCode()
			ctx.LogFields.Latency = ctx.endTime.Sub(ctx.startTime)

			if !IsProduction {
				ctx.Print()
			}
		}()

		err = c.Next()
		return err
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

func (ctx *RequestContext) Print() {
	switch {
	case len(ctx.LogFields.Stack) != 0:
		log.Error().EmbedObject(ctx.LogFields).Msg("panic recover")
	case ctx.LogFields.httpCode >= 500:
		log.Error().EmbedObject(ctx.LogFields).Msg("server error")
	case ctx.LogFields.httpCode >= 400:
		log.Error().EmbedObject(ctx.LogFields).Msg("client error")
	case ctx.LogFields.httpCode >= 300:
		log.Warn().EmbedObject(ctx.LogFields).Msg("redirect")
	case ctx.LogFields.httpCode >= 200:
		log.Info().EmbedObject(ctx.LogFields).Msg("success")
	case ctx.LogFields.httpCode >= 100:
		log.Info().EmbedObject(ctx.LogFields).Msg("informative")
	default:
		log.Warn().EmbedObject(ctx.LogFields).Msg("unknown status")
	}
}
