package index

import (
	"time"

	"github.com/kimchhung/gva/backend-echo/internal/echoc"
	"github.com/kimchhung/gva/backend-echo/internal/request"
	"github.com/kimchhung/gva/backend-echo/internal/response"
	"github.com/labstack/echo/v4"
)

var _ interface{ echoc.Controller } = (*IndexController)(nil)

type IndexController struct {
	index_s *IndexService
}

func (con *IndexController) Init(r *echo.Group) *echo.Group {
	return r
}

func NewIndexController(index_s *IndexService) *IndexController {
	return &IndexController{
		index_s: index_s,
	}
}

// @Tags        Time
// @Summary     Current Server Time
// @ID          now
// @Produce     json
// @Success     200 {object} response.Response{data=string} "format time.RFC3339"
// @Router      /now [get]
// @Security    Bearer
func (con *IndexController) Now(meta *echoc.RouteMeta) echoc.MetaHandler {
	return meta.Get("/now").Do(func(c echo.Context) error {
		now, err := con.index_s.Now(c.Request().Context())
		if err != nil {
			return err
		}

		return request.Response(c, response.Data(now.UTC().Format(time.RFC3339)))
	})
}
