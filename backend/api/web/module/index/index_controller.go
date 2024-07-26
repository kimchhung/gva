package index

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

var _ interface{ ctr.CTR } = (*IndexController)(nil)

type IndexController struct {
	index_s *IndexService
	log     zerolog.Logger
	psub    pubsub.Pubsub
}

func (con *IndexController) Init() *ctr.Ctr {
	return ctr.New()
}

func NewIndexController(index_s *IndexService, log *zerolog.Logger, psub pubsub.Pubsub) *IndexController {
	return &IndexController{
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
// @Router      /now [get]
func (con *IndexController) Now() *ctr.Route {
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
// @Router      /sse/now [get]
func (con *IndexController) SSENow() *ctr.Route {
	ticker := time.NewTicker(3 * time.Second)
	topic := "time_updates"

	go func() {
		for {
			<-ticker.C
			ts := "time: " + time.Now().Format(time.RFC3339)
			if err := con.psub.Pub(context.Background(), topic, ts); err != nil {
				con.log.Error().Err(err).Msg("")
			}
			fmt.Println("publish:", ts)
		}
	}()

	return ctr.GET("/sse/now").Do(func() []ctr.H {
		return []ctr.H{
			func(c echo.Context) error {
				w := c.Response()
				w.Header().Set("Content-Type", "text/event-stream")
				w.Header().Set("Cache-Control", "no-cache")
				w.Header().Set("Connection", "keep-alive")

				subscription, err := con.psub.Sub(context.Background(), topic)
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
							Data: []byte(payload.(string)),
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
// @Router      /ws/now [get]
func (con *IndexController) WSNow() *ctr.Route {
	upgrader := websocket.Upgrader{}
	return ctr.GET("/ws/now").Do(func() []ctr.H {
		return []ctr.H{
			func(c echo.Context) error {
				ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
				if err != nil {
					return err
				}

				defer ws.Close()
				result, err := con.psub.Sub(context.Background(), "time_updates")
				if err != nil {
					return err
				}
				defer result.UnSub()
				ticker := time.NewTicker(1 * time.Second)
				go func() {
					for {
						<-ticker.C
						err := con.psub.Pub(context.Background(), "time_updates", time.Now().UTC().Format(time.RFC3339))
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
