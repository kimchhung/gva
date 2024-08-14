package index

import (
	"time"

	"github.com/gva/internal/ctr"
	"github.com/gva/internal/request"
	"github.com/gva/internal/response"

	"github.com/labstack/echo/v4"
)

var _ interface{ ctr.CTR } = (*IndexController)(nil)

type IndexController struct {
	index_s *IndexService
}

func (con *IndexController) Init() *ctr.Ctr {
	return ctr.New()
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
func (con *IndexController) Now() *ctr.Route {
	return ctr.GET("/now").Do(func() []ctr.H {
		return []ctr.H{
			func(c echo.Context) error {
				now, err := con.index_s.Now(c.Request().Context())
				if err != nil {
					return err
				}

				return request.Response(c, response.Data(now.UTC().Format(time.RFC3339)))
			},
		}
	})
}
