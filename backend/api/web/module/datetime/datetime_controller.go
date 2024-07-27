package mdatetime

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gva/internal/ctr"
	"github.com/gva/internal/pubsub"
	"github.com/gva/internal/request"
	"github.com/gva/internal/response"
	"github.com/gva/utils/sse"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
)

var _ interface{ ctr.CTR } = (*DatetimeController)(nil)

type DatetimeController struct {
	index_s *DatetimeService
	log     zerolog.Logger
	psub    pubsub.Pubsub
}

func (con *DatetimeController) Init() *ctr.Ctr {
	return ctr.New(
		ctr.Group("datetime"),
	)
}

func NewIndexController(index_s *DatetimeService, log *zerolog.Logger, psub pubsub.Pubsub) *DatetimeController {
	return &DatetimeController{
		index_s: index_s,
		log: log.With().
			Str("module", "index").
			Str("provider", "controller").Logger(),
		psub: psub,
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

				subscription, err := con.psub.Sub(context.Background(), "now")
				if err != nil {
					return err
				}

				defer func() {
					if err := subscription.UnSub(); err != nil {
						con.log.Error().Err(err).Msg("")
					}
				}()

				for {
					select {
					case <-c.Request().Context().Done():
						log.Printf("SSE client disconnected, ip: %v", c.RealIP())
						return nil

					case payload := <-subscription.Payload():
						defer fmt.Println("recieve:", payload)

						event := sse.Event{
							Data: []byte(payload.(time.Time).Format(time.RFC3339)),
						}
						if err := event.MarshalTo(w); err != nil {
							return err
						}
						w.Flush()
					}
				}
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
				result, err := con.psub.Sub(context.Background(), "now")
				if err != nil {
					return err
				}
				defer result.UnSub()
				ticker := time.NewTicker(1 * time.Second)
				go func() {
					for {
						<-ticker.C
						err := con.psub.Pub(context.Background(), "now", time.Now().UTC().Format(time.RFC3339))
						if err != nil {
							log.Printf("error publishing: %v", err)
						}
					}
				}()

				log.Printf("WS client connected, ip: %v", c.RealIP())

				for {
					select {
					case <-c.Request().Context().Done():
						log.Printf("WS client disconnected, ip: %v", c.RealIP())
						return nil
					case payload := <-result.Payload():
						err := ws.WriteJSON(payload)
						if err != nil {
							con.log.Err(err).Msg("write message")
						}
						return nil
					}
				}
			},
		}
	})
}
