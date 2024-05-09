package contexts

import (
	"context"
	"errors"
	"fmt"
	"runtime/debug"
	"time"

	"github.com/gofiber/fiber/v2"
	app_err "github.com/kimchhung/gva/extra/app/common/error"
	rerror "github.com/kimchhung/gva/extra/internal/response/error"

	"github.com/google/uuid"
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
	HttpCode  int
	ErrorCode int
	Latency   time.Duration
	Error     error
	Stack     []byte
}

func (lf *logFields) MarshalZerologObject(e *zerolog.Event) {
	e.Err(lf.Error).
		Str("id", lf.ID).
		Str("remoteIp", lf.RemoteIP).
		Str("host", lf.Host).
		Str("method", lf.Method).
		Str("path", lf.Path).
		Str("protocol", lf.Protocol).
		Int("httpCode", lf.HttpCode).
		Int("errorCode", lf.ErrorCode).
		Str("latency", fmt.Sprintf("%v", lf.Latency)).
		Str("tag", "request")

	if lf.Stack != nil {
		fmt.Println(string(lf.Stack))
	}
}

type (
	RequestContextKey    struct{}
	RequestContextOption func(*RequestContext)
)

type RequestContext struct {
	context.Context
	startTime time.Time
	logFields *logFields
}

func (ctx *RequestContext) PrintLog() {
	switch {
	case len(ctx.logFields.Stack) != 0:
		log.Error().EmbedObject(ctx.logFields).Msg("panic recover")
	case ctx.logFields.HttpCode >= 500:
		log.Error().EmbedObject(ctx.logFields).Msg("server error")
	case ctx.logFields.HttpCode >= 400:
		log.Warn().EmbedObject(ctx.logFields).Msg("client error")
	case ctx.logFields.HttpCode >= 300:
		log.Warn().EmbedObject(ctx.logFields).Msg("redirect")
	case ctx.logFields.HttpCode >= 200:
		log.Info().EmbedObject(ctx.logFields).Msg("success")
	case ctx.logFields.HttpCode >= 100:
		log.Info().EmbedObject(ctx.logFields).Msg("informative")
	default:
		log.Warn().EmbedObject(ctx.logFields).Msg("unknown status")
	}
}

func defaultLogFields(c *fiber.Ctx) *logFields {
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
func NewRequestContext() func(c *fiber.Ctx) (err error) {
	return func(c *fiber.Ctx) error {
		ctx := &RequestContext{}
		ctx.Context = context.WithValue(c.Context(), RequestContextKey{}, ctx)
		ctx.startTime = time.Now()
		ctx.logFields = defaultLogFields(c)
		c.SetUserContext(ctx)

		var err error
		defer func() {
			if rvr := recover(); rvr != nil {
				fmt.Printf("errr rrr %v", rvr)
				var ok bool

				if err, ok = rvr.(*app_err.Error); !ok {
					// unknown internal error stacks
					ctx.logFields.Stack = debug.Stack()
				}

				if err, ok = rvr.(error); !ok {
					// internal error stacks
					err = fmt.Errorf("%v", rvr)
				}
			}

			ctx.logFields.Latency = time.Since(ctx.startTime)
			if err != nil {

				// case using panic to handler error instead of return error
				err = ErrorHandler(c, err)
			}
		}()

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

func SetRequestStatus(ctx context.Context, errorCode int, httpCode int) *RequestContext {
	rctx := MustRequestContext(ctx)
	rctx.logFields.HttpCode = httpCode
	rctx.logFields.ErrorCode = errorCode
	return rctx
}

// Default error handler
func ErrorHandler(c *fiber.Ctx, err error) error {
	perr, err := rerror.ParseError(c, err)

	rctx, rerr := GetRequestContext(c.UserContext())
	if rerr != nil {
		log.Error().Err(rerr).Msg("GetRequestContext")
		return err
	}

	rctx.logFields.HttpCode = perr.HttpCode
	rctx.logFields.ErrorCode = perr.ErrorCode
	rctx.logFields.Error = perr

	if !IsProduction {
		rctx.PrintLog()
	}

	return err
}
