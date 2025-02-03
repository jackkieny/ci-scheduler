package cron

import (
	"time"

	"github.com/robfig/cron/v3"
	"github.com/rs/zerolog/log"
)

func Init() *cron.Cron {
	start := time.Now()

	c := cron.New(cron.WithLocation(time.UTC))
	c.Start()

	elapsed := time.Since(start)
	log.Info().Msgf("Cron initialized! (took %.3fs)\n", elapsed.Seconds())

	return c
}
