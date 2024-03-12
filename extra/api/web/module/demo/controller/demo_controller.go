package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kimchhung/gva/extra/internal/rctrl"
	"github.com/kimchhung/gva/extra/internal/request"
	"github.com/kimchhung/gva/extra/internal/response"
)

type DemoController struct {
	rctrl.Controller
}

func (con *DemoController) Init(r fiber.Router) fiber.Router {
	return r.Group("demo")
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
// @Success   200 {object} response.Response{data=any}
// @Router /demo/ [get]
func (con *DemoController) Welcome(meta *rctrl.RouteMeta) rctrl.MetaHandler {
	return meta.Get("/").Name("demo").Do(
		func(c *fiber.Ctx) error {
			d := response.Response{}

			return request.Response(c,
				response.Data(d),
			)
		},
	)
}
