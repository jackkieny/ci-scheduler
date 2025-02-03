package db

import (
	"context"
	"time"

	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func FindScheduledPosts(client *mongo.Client) *mongo.Cursor {

	log.Info().Msg("Finding scheduled posts...")

	collection := client.Database("community_insights").Collection("posts")
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()

	filter := bson.D{{Key: "status", Value: 1}}

	count, err := collection.CountDocuments(ctx, filter)
	if err != nil {
		log.Error().Err(err).Str("func", "FindScheduledPosts").Msg("Error counting scheduled posts")
	} else {
		log.Info().Msgf("Found %d scheduled posts", count)
	}

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		log.Error().Err(err).Str("func", "FindScheduledPosts").Msg("Error finding scheduled posts")
	}

	if cursor == nil {
		log.Panic().Str("func", "FindScheduledPosts").Msg("Cursor is nil")
	}

	return cursor
}

func FindUnscheduledPosts(client *mongo.Client) *mongo.Cursor {

	log.Info().Msg("Finding unscheduled posts...")

	collection := client.Database("community_insights").Collection("posts")
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()

	filter := bson.D{{Key: "status", Value: 0}}

	count, err := collection.CountDocuments(ctx, filter)
	if err != nil {
		log.Error().Err(err).Str("func", "FindUnscheduledPosts").Msg("Error counting scheduled posts")
	} else {
		log.Info().Msgf("Found %d scheduled posts", count)
	}

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		log.Error().Err(err).Str("func", "FindUnscheduledPosts").Msg("Error finding unscheduled posts")
	}

	if cursor == nil {
		log.Panic().Str("func", "FindUnscheduledPosts").Msg("Cursor is nil")
	}

	return cursor
}
