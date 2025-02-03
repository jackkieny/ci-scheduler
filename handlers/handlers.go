package handlers

import (
	"fmt"
	"time"

	"github.com/rs/zerolog/log"
)

func ParseTime(isoTime string) (string, error) {
	layout := "2006-01-02 15:04:05.000 -0700 MST"
	t, err := time.Parse(layout, isoTime)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to parse time")
	}

	minute := t.Minute()
	hour := t.Hour()
	day := t.Day()
	month := t.Month()
	weekday := t.Weekday()

	cron := fmt.Sprintf("%d %d %d %d %d", minute, hour, day, month, weekday)
	fmt.Println(cron)
	return cron, nil
}
