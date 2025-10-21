package appctx

import (
	"context"
	"errors"
	"fmt"
	"runtime/debug"
	"time"

	apperror "backend/app/common/error"
	"backend/env"
	"backend/internal/logger"
	"backend/internal/pxid"
	rerror "backend/internal/response/error"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var IsProduction bool

type LogFields struct {
	ID        string
	RemoteIP  string
	Host      string
	Method    string
	Path      string
	Protocol  string
	HttpCode  int
	ErrorCode int
	ErrorMsg  string
	Latency   time.Duration
	Error     error
	Stack     []byte
	// permission scope
	Scopes    []string
	UserAgent string

	MetaData map[string]any
}

func (l LogFields) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	if l.Error != nil {
		enc.AddString("error", l.Error.Error())
	}
	enc.AddString("id", l.ID)
	enc.AddString("remoteIP", l.RemoteIP)
	enc.AddString("host", l.Host)
	enc.AddString("method", l.Method)
	enc.AddString("path", l.Path)
	enc.AddString("protocol", l.Protocol)
	enc.AddInt("httpCode", l.HttpCode)
	enc.AddInt("errorCode", l.ErrorCode)
	enc.AddString("errorMsg", l.ErrorMsg)
	enc.AddDuration("latency", l.Latency)

	if l.MetaData != nil {
		zap.Any("meta", l.MetaData).AddTo(enc)
	}

	return nil
}

type (
	RequestContextKey    struct{}
	RequestContextOption func(*RequestContext)
)

type RequestContext struct {
	context.Context
	startTime time.Time
	LogFields *LogFields
	appEnv    string
}

func (c *RequestContext) IsProd() bool {
	return c.appEnv == env.Prod
}

func (c *RequestContext) IsDev() bool {
	return c.appEnv == env.Dev
}

func (c *RequestContext) IsStaging() bool {
	return c.appEnv == env.Stag
}

func (ctx *RequestContext) PrintLog() {
	switch {
	case len(ctx.LogFields.Stack) != 0:
		logger.G().Error("panic recover", zap.Inline(ctx.LogFields))
	case ctx.LogFields.HttpCode >= 500:
		logger.G().Error("server error", zap.Inline(ctx.LogFields))
	case ctx.LogFields.HttpCode >= 400:
		logger.G().Info("client error", zap.Inline(ctx.LogFields))
	case ctx.LogFields.HttpCode >= 300:
		logger.G().Info("redirect", zap.Inline(ctx.LogFields))
	case ctx.LogFields.HttpCode >= 200:
		logger.G().Info("success", zap.Inline(ctx.LogFields))
	case ctx.LogFields.HttpCode >= 100:
		logger.G().Info("informative", zap.Inline(ctx.LogFields))
	}
}

func defaultLogFields(c echo.Context) *LogFields {
	req := c.Request()
	rid := req.Header.Get(echo.HeaderXRequestID)
	if rid == "" {
		rid = string(pxid.New("req"))
		c.Set(echo.HeaderXRequestID, rid)
	}

	fields := &LogFields{
		ID:        rid,
		RemoteIP:  c.RealIP(),
		Method:    req.Method,
		Host:      req.Host,
		Path:      req.URL.Path,
		Protocol:  req.Proto,
		UserAgent: req.UserAgent(),
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
			ctx.LogFields = defaultLogFields(c)
			ctx.appEnv = cfg.App.Env

			//logger set context
			c.SetRequest(c.Request().WithContext(ctx))
			defer func() {
				if rvr := recover(); rvr != nil {

					var ok bool

					if err, ok = rvr.(*apperror.Error); !ok {
						// unknown internal error stacks
						ctx.LogFields.Stack = debug.Stack()
					}

					ctx.LogFields.Stack = debug.Stack()

					if err, ok = rvr.(error); !ok {
						// internal error stacks
						err = fmt.Errorf("%v", rvr)
					}

				}

				if err != nil {

					// case using panic to handler error instead of return error
					err = ErrorHandler(err, c)
				}

				ctx.LogFields.Latency = time.Since(ctx.startTime)
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
	rctx.LogFields.HttpCode = httpCode
	rctx.LogFields.ErrorCode = errorCode
	return rctx
}

func SetRequestParams(ctx context.Context, params map[string]any) *RequestContext {
	rctx := MustRequestContext(ctx)
	if rctx.LogFields.MetaData == nil {
		rctx.LogFields.MetaData = make(map[string]any)
	}
	rctx.LogFields.MetaData["params"] = params
	return rctx
}

// Default error handler
func ErrorHandler(anyErr error, c echo.Context) error {

	perr, err := rerror.ParseError(c, anyErr)

	ctx := c.Request().Context()

	rctx, rerr := GetRequestContext(ctx)
	if rerr != nil {
		logger.G().Error("GetRequestContext", zap.Error(err))
		return err
	}

	rctx.LogFields.HttpCode = perr.HttpCode
	rctx.LogFields.ErrorCode = perr.ErrorCode
	rctx.LogFields.ErrorMsg = perr.Message

	if perr.ErrorCode == apperror.ErrUnknownError.ErrorCode {
		rctx.LogFields.Error = apperror.NewError(perr, apperror.Join(anyErr))
	} else {
		rctx.LogFields.Error = perr
	}

	if !rctx.IsProd() && !perr.IsPublic() {
		rctx.PrintLog()
	}

	return err
}
