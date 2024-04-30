package controller

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/kimchhung/gva/extra/internal/bootstrap/database"
	"github.com/kimchhung/gva/extra/internal/rctrl"
	"github.com/kimchhung/gva/extra/internal/request"
	"github.com/kimchhung/gva/extra/internal/response"
)

type DemoController struct {
	rctrl.Controller
	db *database.Database
}

func (con *DemoController) Init(r fiber.Router) fiber.Router {
	return r
}

func NewDemoController(db *database.Database) *DemoController {
	return &DemoController{}
}

// @Tags        Demo
// @Summary     List all Demos
// @Description Welcome
// @ID          welcome-all-Demos
// @Accept      json
// @Produce     json
// @Success     200 {object} response.Response{data=any}
// @Router      / [get]
func (con *DemoController) Welcome(meta *rctrl.RouteMeta) rctrl.MetaHandler {

	return meta.Get("/").Do(
		func(c *fiber.Ctx) error {

			now := ""
			rows, err := con.db.Sql().QueryContext(c.UserContext(), "select now()")
			if err != nil {
				panic(err)
			}

			rows.Scan(&now)
			return request.Response(c,
				response.Data(
					fmt.Sprintf("Welcome to my web api %v", now),
				),
			)
		},
	)
}
