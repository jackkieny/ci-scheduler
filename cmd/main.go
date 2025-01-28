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
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.ANSIC})

    client := db.Init()
    broker := broker.Init()
    cron := cron.Init()

    fmt.Print(client, broker, cron)

}
