package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kimchhung/gva/extra/internal/rctrl"
	"github.com/kimchhung/gva/extra/internal/request"
	"github.com/kimchhung/gva/extra/internal/response"
)

// don't remove for runtime type checking
var _ interface{ rctrl.Controller } = (*DemoController)(nil)

type DemoController struct {
}

func (con *DemoController) Routes(r fiber.Router) {
	Demo := r.Group("demo")
	rctrl.Register(Demo, con)
}

func NewDemoController() *DemoController {
	return &DemoController{}
}

// @Tags Demo
// @Summary List all Demos
// @Description Welcome
// @ID welcome-all-Demos
// @Accept  json
// @Produce  json
// @Success 200 {object} response.Response{} "Welcome"
// @Demor / [get]
func (con *DemoController) Welcome(meta *rctrl.RouteMeta) rctrl.MetaHandler {
	return meta.Get("/").Name("demo").Do(
		func(c *fiber.Ctx) error {

			return request.Response(c,
				response.Data(map[string]any{}),
			)
		},
	)
}
