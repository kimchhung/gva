package index

import (
	"log"
	"time"

	"github.com/gorilla/websocket"
	"github.com/kimchhung/gva/extra/internal/echoc"
	"github.com/kimchhung/gva/extra/internal/request"
	"github.com/kimchhung/gva/extra/internal/response"
	"github.com/kimchhung/gva/extra/utils/sse"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
)

var _ interface{ echoc.Controller } = (*IndexController)(nil)

type IndexController struct {
	index_s *IndexService
	log     zerolog.Logger
}

func (con *IndexController) Init(r *echo.Group) *echo.Group {
	return r
}

func NewIndexController(index_s *IndexService, log *zerolog.Logger) *IndexController {
	return &IndexController{
		index_s: index_s,
		log: log.With().
			Str("module", "index").
			Str("provider", "controller").Logger(),
	}
}

// @Tags        Time
// @Summary     Current Server Time
// @ID          now
// @Produce     json
// @Success     200 {object} response.Response{data=string} "format time.RFC3339"
// @Router      /now [get]
func (con *IndexController) Now(meta *echoc.RouteMeta) echoc.MetaHandler {
	return meta.Get("/now").Do(func(c echo.Context) error {
		now, err := con.index_s.Now(c.Request().Context())
		if err != nil {
			return err
		}

		return request.Response(c, response.Data(now.UTC().Format(time.RFC3339)))
	})
}

// @Tags        Time
// @Summary     SSE Current Server Time
// @ID          sse-now
// @Accept 		text/event-stream
// @Success     200 {object} string "format time.RFC3339"
// @Router      /sse/now [get]
func (con *IndexController) SSENow(meta *echoc.RouteMeta) echoc.MetaHandler {
	return meta.Get("/sse/now").Do(func(c echo.Context) error {
		log.Printf("SSE client connected, ip: %v", c.RealIP())

		w := c.Response()
		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")

		ticker := time.NewTicker(1 * time.Second)
		defer ticker.Stop()
		for {
			select {
			case <-c.Request().Context().Done():
				log.Printf("SSE client disconnected, ip: %v", c.RealIP())
				return nil
			case <-ticker.C:
				event := sse.Event{
					Data: []byte("time: " + time.Now().Format(time.RFC3339)),
				}
				if err := event.MarshalTo(w); err != nil {
					return err
				}
				w.Flush()
			}
		}
	})
}

var upgrader = websocket.Upgrader{}

// @Tags        Time
// @Summary     WS Current Server Time
// @ID          ws-now
// @Accept 		text/event-stream
// @Success     200 {object} string "format time.RFC3339"
// @Router      /ws/now [get]
func (con *IndexController) WSNow(meta *echoc.RouteMeta) echoc.MetaHandler {
	return meta.Get("/ws/now").Do(func(c echo.Context) error {
		ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
		if err != nil {
			return err
		}

		defer ws.Close()

		ticker := time.NewTicker(1 * time.Second)
		defer ticker.Stop()

		log.Printf("WS client connected, ip: %v", c.RealIP())

		for {
			select {
			case <-c.Request().Context().Done():
				log.Printf("WS client disconnected, ip: %v", c.RealIP())
				return nil
			case <-ticker.C:
				err := ws.WriteJSON(time.Now().UTC().Format(time.RFC3339))
				if err != nil {
					con.log.Err(err).Msg("write message")
				}
				return nil
			}
		}
	})
}
