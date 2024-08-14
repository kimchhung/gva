package mdatetime

import (
	"time"

	"github.com/gva/app/common/service"
	"github.com/gva/internal/ctr"
	"github.com/gva/internal/request"
	"github.com/gva/internal/response"
	"github.com/gva/utils/sse"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
)

var _ interface{ ctr.CTR } = (*DatetimeController)(nil)

type DatetimeController struct {
	index_s  *DatetimeService
	log      zerolog.Logger
	service  *DatetimeService
	pubsub_s *service.PubsubService
}

func (con *DatetimeController) Init() *ctr.Ctr {
	return ctr.New(
		ctr.Group("datetime"),
	)
}

func NewIndexController(
	index_s *DatetimeService,
	log *zerolog.Logger,
	service *DatetimeService,
	pubsub_s *service.PubsubService,
) *DatetimeController {
	return &DatetimeController{
		index_s: index_s,
		log: log.With().
			Str("module", "index").
			Str("provider", "controller").Logger(),
		pubsub_s: pubsub_s,
		service:  service,
	}
}

// @Tags        Time
// @Summary     Current Server Time
// @ID          now
// @Produce     json
// @Success     200 {object} response.Response{data=string} "format time.RFC3339"
// @Router      /datetime/now [get]
func (con *DatetimeController) Now() *ctr.Route {
	return ctr.GET("/now").Do(func() []ctr.H {
		return []ctr.H{
			func(c echo.Context) error {
				now, err := con.index_s.Now(c.Request().Context())
				if err != nil {
					return err
				}

				return request.Response(c,
					response.Data(now.UTC().Format(time.RFC3339)),
				)
			},
		}
	})
}

// @Tags        Time
// @Summary     SSE Current Server Time
// @ID          sse-now
// @Accept 		text/event-stream
// @Success     200 {object} string "format time.RFC3339"
// @Router      /datetime/sse/now [get]
func (con *DatetimeController) SSENow() *ctr.Route {
	return ctr.GET("/sse/now").Do(func() []ctr.H {
		return []ctr.H{
			func(c echo.Context) error {
				w := c.Response()
				w.Header().Set("Content-Type", "text/event-stream")
				w.Header().Set("Cache-Control", "no-cache")
				w.Header().Set("Connection", "keep-alive")

				nowChan, err := con.service.NowChannel(c.Request().Context())
				if err != nil {
					return err
				}

				for now := range nowChan {
					event := sse.Event{
						Data: []byte(now.Format(time.RFC3339)),
					}
					if err := event.MarshalTo(w); err != nil {
						return err
					}

					w.Flush()
				}
				return nil
			},
		}
	})
}

// @Tags        Time
// @Summary     WS Current Server Time
// @ID          ws-now
// @Accept 		text/event-stream
// @Success     200 {object} string "format time.RFC3339"
// @Router      /datetime/now [get]
func (con *DatetimeController) WSNow() *ctr.Route {
	upgrader := websocket.Upgrader{}
	return ctr.GET("/ws/now").Do(func() []ctr.H {
		return []ctr.H{
			func(c echo.Context) error {
				ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
				if err != nil {
					return err
				}

				defer ws.Close()
				nowChan, err := con.service.NowChannel(c.Request().Context())
				if err != nil {
					return nil
				}

				for now := range nowChan {
					err := ws.WriteJSON(now)
					if err != nil {
						con.log.Err(err).Msg("write message")
					}
				}
				return nil
			},
		}
	})
}
