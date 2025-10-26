package middleware

import (
	admincontext "backend/app/admin/context"
	"backend/app/share/model"
	corecontext "backend/core/context"
	"backend/core/utils/json"
	"backend/core/utils/request"
	"backend/core/utils/response"
	"bytes"
	"context"
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

func AddOperationLogData(ctx context.Context, data model.OperationLogData) model.OperationLogData {
	operationData, _ := corecontext.SetOrGet(ctx, OperationLogDataKey, model.OperationLogData{})
	for k, v := range data {
		operationData[k] = v
	}
	return operationData
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

			switch c.Request().Method {
			case http.MethodPost, http.MethodPut, http.MethodPatch:
				if c.Request().Body != nil {
					bodyBytes, _ := io.ReadAll(c.Request().Body)

					if strings.Contains(c.Request().Header.Get("Content-Type"), "json") {
						AddOperationLogData(ctx, model.OperationLogData{
							"body": json.JSON(bodyBytes).Object(),
						})

					} else {
						AddOperationLogData(ctx, model.OperationLogData{
							"body": json.JSON(bodyBytes).String(),
						})
					}

					c.Request().Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
				}

			default:
				queries := c.Request().URL.Query()
				if len(queries) > 0 {
					AddOperationLogData(ctx, model.OperationLogData{
						"queries": flattenQueryParams(queries),
					})
				}
			}

			hook.OnBeforeResponse(func(c echo.Context, resp *response.Response) {
				ctx := c.Request().Context()
				if _, err := corecontext.Get[struct{}](ctx, skipOperationLog); err == nil {
					return
				}

				adminctx, err := admincontext.GetAdminContext(ctx)
				if err != nil {
					log.Error("get adminctx failed", zap.Error(err), zap.String("path", c.Path()))
					return
				}

				data := AddOperationLogData(ctx, model.OperationLogData{})
				if c.Request().Method != http.MethodGet {
					data["response"] = resp
				}

				op := &model.OperationLog{
					Method:    c.Request().Method,
					Path:      c.Path(),
					IP:        c.RealIP(),
					Data:      data,
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
