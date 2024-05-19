package index

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/kimchhung/gva/extra/internal/rctrl"
	"github.com/kimchhung/gva/extra/internal/request"
	"github.com/kimchhung/gva/extra/internal/response"
)

var _ interface{ rctrl.Controller } = (*IndexController)(nil)

type IndexController struct {
	index_s *IndexService
}

func (con *IndexController) Init(r fiber.Router) fiber.Router {
	return r
}

func NewIndexController(index_s *IndexService) *IndexController {
	return &IndexController{
		index_s: index_s,
	}
}

func (con *IndexController) Now(meta *rctrl.RouteMeta) rctrl.MetaHandler {
	return meta.Get("/now").Do(func(c *fiber.Ctx) error {
		now, err := con.index_s.Now(c.UserContext())
		if err != nil {
			return err
		}

		return request.Response(c, response.Data(now.UTC().Format(time.RFC3339)))
	})
}
