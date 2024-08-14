package appctx

import (
	"context"
	"errors"
	"fmt"
	"runtime/debug"
	"time"

	apperror "github.com/gva/app/common/error"
	"github.com/gva/env"
	rerror "github.com/gva/internal/response/error"

	"github.com/labstack/echo/v4"

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
	appEnv    string
}

func (c *RequestContext) IsProd() bool {
	return c.appEnv == env.Prod
}

func (c *RequestContext) IsDev() bool {
	return c.appEnv == env.Dev
}

func (c *RequestContext) IsStaging() bool {
	return c.appEnv == env.Staging
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

func defaultLogFields(c echo.Context) *logFields {
	req := c.Request()
	rid := req.Header.Get(echo.HeaderXRequestID)
	if rid == "" {
		rid = uuid.New().String()
		c.Set(echo.HeaderXRequestID, rid)
	}

	fields := &logFields{
		ID:       rid,
		RemoteIP: c.RealIP(),
		Method:   req.Method,
		Host:     req.Host,
		Path:     req.URL.Path,
		Protocol: req.Proto,
	}

	return fields
}

// a context help handling error
func Middleware(cfg *env.Config) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {

		return func(c echo.Context) (err error) {
			ctx := &RequestContext{}
			ctx.Context = context.WithValue(c.Request().Context(), RequestContextKey{}, ctx)
			ctx.startTime = time.Now()
			ctx.logFields = defaultLogFields(c)
			ctx.appEnv = cfg.App.Env

			//logger set context
			c.SetRequest(c.Request().WithContext(ctx))

			defer func() {
				if rvr := recover(); rvr != nil {

					var ok bool

					if err, ok = rvr.(*apperror.Error); !ok {
						// unknown internal error stacks
						ctx.logFields.Stack = debug.Stack()
					}

					ctx.logFields.Stack = debug.Stack()

					if err, ok = rvr.(error); !ok {
						// internal error stacks
						err = fmt.Errorf("%v", rvr)
					}

				}

				if err != nil {

					// case using panic to handler error instead of return error
					err = ErrorHandler(err, c)
				}

				ctx.logFields.Latency = time.Since(ctx.startTime)

			}()

			return next(c)
		}
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
func ErrorHandler(anyErr error, c echo.Context) error {
	perr, err := rerror.ParseError(c, anyErr)
	ctx := c.Request().Context()

	rctx, rerr := GetRequestContext(ctx)
	if rerr != nil {
		log.Error().Err(rerr).Msg("GetRequestContext")
		return err
	}

	rctx.logFields.HttpCode = perr.HttpCode
	rctx.logFields.ErrorCode = perr.ErrorCode

	if perr.ErrorCode == apperror.ErrUnknownError.ErrorCode {
		rctx.logFields.Error = apperror.NewError(perr, apperror.Join(anyErr))
	} else {
		rctx.logFields.Error = perr
	}

	if !IsProduction {
		rctx.PrintLog()
	}

	return err
}
