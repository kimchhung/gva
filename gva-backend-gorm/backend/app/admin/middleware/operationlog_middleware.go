package middleware

import (
	admincontext "backend/app/admin/context"
	"backend/app/share/model"
	corecontext "backend/core/context"
	"backend/core/utils/json"
	"backend/core/utils/request"
	"backend/core/utils/response"
	"bytes"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"go.uber.org/zap"
)

const (
	skipOperationLog    corecontext.Key = "skip_opration_log"
	OperationLogDataKey corecontext.Key = "operationlog_data"
)

// EnableOperationLog is a middleware to skip operation logger
func SkipOperationLog() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			corecontext.Set(c.Request().Context(), skipOperationLog, struct{}{})
			return next(c)
		}
	}
}

func (m *Middleware) OperationLog() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			ctx := c.Request().Context()

			hook, err := request.GetHook(ctx)
			if err != nil {
				return next(c)
			}

			now := time.Now()
			hook.OnBeforeResponse(func(c echo.Context, resp *response.Response) {
				ctx := c.Request().Context()
				if _, err := corecontext.Get[struct{}](ctx, skipOperationLog); err == nil {
					return
				}

				adminctx, err := admincontext.GetAdminContext(ctx)
				if err != nil {
					log.Error("get adminctx failed", zap.Error(err))
					return
				}

				method := c.Request().Method
				operationData, _ := corecontext.SetOrGet(ctx, OperationLogDataKey, model.OperationLogData{})

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
					queries := c.Request().URL.Query()
					if len(queries) > 0 {
						operationData["queries"] = flattenQueryParams(queries)
					}
				}

				op := &model.OperationLog{
					Method:    method,
					Path:      c.Path(),
					Data:      operationData,
					IP:        c.RealIP(),
					AdminId:   adminctx.Admin.ID,
					RoleIds:   adminctx.Admin.RoleIds,
					Scope:     strings.Join(adminctx.EndpointScopes, ","),
					Code:      resp.Code,
					Msg:       resp.Message,
					CreatedAt: time.Now(),
					Latency:   time.Since(now).Milliseconds(),
				}

				if err := m.db.Create(op).Error; err != nil {
					log.Error("create operationlog failed", zap.Error(err))
				}
			})

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
