package cron

import (
	"context"

	"github.com/jackkieny/ci-scheduler/handlers"
	"github.com/jackkieny/ci-scheduler/types"
	"github.com/robfig/cron/v3"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateJob(c *cron.Cron, cursor *mongo.Cursor) {
	for cursor.Next(context.TODO()) {
		var post types.Post
		err := cursor.Decode(&post)
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to decode cursor")
		}

		log.Info().Msgf("Creating job for post %s", post.ID)

		time, err := handlers.ParseTime(post.Datetime.String())
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to parse time")
		}

		jobId, err := c.AddFunc(time, func() {
			log.Info().Msgf("Running job for post %s", post.ID)
		})
		if err != nil {
			log.Fatal().Err(err).Msgf("Failed to add job for post %s", post.ID)
		}

		log.Info().Msgf("Job for post %s created! %d", post.ID, jobId)
	}
}
