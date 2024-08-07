package system

import (
	"github.com/gva/app/common/controller"
	"github.com/gva/internal/ctr"
	"github.com/gva/internal/request"
	"github.com/gva/internal/response"
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

type SystemController struct {
}

func NewMenuController() *SystemController {
	return &SystemController{}
}

func (con *SystemController) Init(r *echo.Group) *echo.Group {
	return r.Group("/system")
}

func (con *SystemController) Test() *ctr.Route {
	return ctr.GET("/").Do(func() []ctr.H {
		return []ctr.H{
			func(c echo.Context) error {
				return request.Response(c,
					response.Data("test"),
				)
			},
		}
	})
}

// Register bulkly
var SystemModule = fx.Module("MenuModule",
	// Register Service
	fx.Provide(NewMenuController),

	// Regiser Controller
	controller.ProvideAdminController(NewMenuController),
)
