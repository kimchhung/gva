package mdatetime

import (
	"context"
	"fmt"
	"time"

	"github.com/gva/internal/bootstrap"
	"github.com/gva/internal/pubsub"
	"github.com/rs/zerolog"
)

func BackgroundNow(log *zerolog.Logger, boot *bootstrap.Bootstrap, psub pubsub.Pubsub) {
	go func() {
		<-boot.Done()

		ticker := time.NewTicker(5 * time.Second)
		topic := "now"

		for {
			<-ticker.C
			ts := time.Now()
			if err := psub.Pub(context.Background(), topic, ts); err != nil {
				fmt.Println("psub.Pub err:", err)
			}
			log.Info().Time("t", ts).Msg("backgroundNow")
		}
	}()
}
