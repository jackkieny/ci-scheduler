package main

import (
	"fmt"
	"os"
	"time"

	"github.com/jackkieny/ci-scheduler/broker"
	"github.com/jackkieny/ci-scheduler/cron"
	"github.com/jackkieny/ci-scheduler/db"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.ANSIC, TimeLocation: time.UTC})

	client := db.Init()
	broker := broker.Init()
	cronClient := cron.Init()

	res := fmt.Sprintf("%v %v", broker, cronClient)
	_ = res

	scheduledCursor := db.FindScheduledPosts(client)
	unscheduledCursor := db.FindUnscheduledPosts(client)

	_ = scheduledCursor
	_ = unscheduledCursor

	// select {}
}
