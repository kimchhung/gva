package mdatetime

import (
	"context"
	"fmt"
	"time"

	"github.com/gva/app/common/service"
	"github.com/gva/internal/bootstrap"
	"github.com/rs/zerolog"
)

func BackgroundNow(log *zerolog.Logger, boot *bootstrap.Bootstrap, pubsub_s *service.PubsubService) {
	go func() {
		<-boot.Done()

		ticker := time.NewTicker(1 * time.Second)
		topic := "now"

		log.Info().Msg("background initialed")

		for {
			<-ticker.C
			ts := time.Now()
			if err := pubsub_s.Local().Publish(context.Background(), topic, &ts); err != nil {
				fmt.Println("psub.Pub err:", err)
			}
		}
	}()
}
