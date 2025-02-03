package types

import "time"

type Post struct {
	ID       string    `bson:"_id"`
	Datetime time.Time `bson:"datetime"`
}
