package middleware

import (
	appctx "backend/app/common/context"
	apperror "backend/app/common/error"
	"backend/app/common/model"
	"backend/internal/bootstrap/database"
	"backend/utils/json"
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
)

type OperationLogContextKey string

var SkipOperationLoggerKey OperationLogContextKey = "skip"

// OperationLogger is a middleware to log operation
func OperationLogger(db *database.Database) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			if _, isSkip := c.Request().Context().Value(SkipOperationLoggerKey).(bool); isSkip {
				return next(c)
			}

			operationData := model.OperationLogData{}
			method := c.Request().Method
			queries := c.Request().URL.Query()
			if len(queries) > 0 {
				operationData["queries"] = flattenQueryParams(queries)
			}

			switch method {
			case http.MethodPost, http.MethodPut, http.MethodPatch:
				var bodyBytes []byte
				if c.Request().Body != nil {
					bodyBytes, err = io.ReadAll(c.Request().Body)
					if err != nil {
						return
					}
					c.Request().Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
				}

				if strings.Contains(c.Request().Header.Get("Content-Type"), "application/json") {
					operationData["body"] = json.JSON(bodyBytes).Object()
				} else {
					operationData["body"] = json.JSON(bodyBytes).String()
				}

			default:
			}

			defer func(c echo.Context, operationData model.OperationLogData) {
				ctx := appctx.MustRequestContext(c.Request().Context())
				adminCtx, err := appctx.GetAdminContext(c.Request().Context())
				if err != nil {
					return
				}

				var scope string
				if len(ctx.LogFields.Scopes) > 0 {
					scope = ctx.LogFields.Scopes[0]
				}

				for k, v := range ctx.LogFields.MetaData {
					operationData[k] = v
				}

				op := model.OperationLog{
					Method:    method,
					Path:      ctx.LogFields.Path,
					Data:      operationData,
					IP:        ctx.LogFields.RemoteIP,
					AdminId:   adminCtx.Admin.ID,
					RoleIds:   adminCtx.Admin.RoleIds,
					Scope:     scope,
					Code:      ctx.LogFields.ErrorCode,
					CreatedAt: time.Now(),
					Latency:   ctx.LogFields.Latency.Milliseconds(),
					Msg:       ctx.LogFields.ErrorMsg,
				}

				if ctx.LogFields.ErrorCode == apperror.ErrUnknownError.ErrorCode {
					op.Error = fmt.Sprintf("error=%v\n stacktrace=%v", ctx.LogFields.Error.Error(), string(ctx.LogFields.Stack))
				}

				db.Create(&op)
			}(c, operationData)

			return next(c)
		}
	}
}

// SkipOperationLogger is a middleware to skip operation logger
func SkipOperationLogger() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx := context.WithValue(c.Request().Context(), SkipOperationLoggerKey, true)
			c.SetRequest(c.Request().WithContext(ctx))
			return next(c)
		}
	}
}

// Helper function to flatten query parameters
func flattenQueryParams(queryParams map[string][]string) map[string]string {
	flattened := make(map[string]string)
	for key, values := range queryParams {
		if len(values) > 0 {
			flattened[key] = values[0]
		} else {
			flattened[key] = ""
		}
	}
	return flattened
}
