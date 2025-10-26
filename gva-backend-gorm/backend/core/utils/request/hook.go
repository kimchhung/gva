package request

import (
	corecontext "backend/core/context"
	"backend/core/utils/response"
	"context"

	"github.com/labstack/echo/v4"
)

const HookContextKey corecontext.Key = "req_context"

type Hook struct {
	onBeforeResponses []func(c echo.Context, resp *response.Response)
}

func (h *Hook) OnBeforeResponse(fn ...func(c echo.Context, resp *response.Response)) {
	h.onBeforeResponses = append(h.onBeforeResponses, fn...)
}

func SetHook(ctx context.Context) {
	corecontext.Set(ctx, HookContextKey, &Hook{})
}

func GetHook(ctx context.Context) (*Hook, error) {
	return corecontext.Get[*Hook](ctx, HookContextKey)
}
