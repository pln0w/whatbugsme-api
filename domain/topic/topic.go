package topic

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Topic struct {
	ID           bson.ObjectId `json:"id" bson:"id"`
	Content      string        `json:"content" bson:"content"`
	VotesUp      int           `json:"votes_up" bson:"votes_up"`
	VotesDown    int           `json:"votes_down" bson:"votes_down"`
	Organisation bson.ObjectId `json:"organisation" bson:"organisation"`
	CreatedAt    time.Time     `json:"created_at" bson:"created_at"`
}

const (
	C_TOPIC = "topics"
)

type Topics []Topic

func CreateNewTopic(content string, organisation bson.ObjectId) Topic {
	return Topic{
		ID:           bson.NewObjectId(),
		Content:      content,
		VotesDown:    0,
		VotesUp:      0,
		Organisation: organisation,
		CreatedAt:    time.Now(),
	}
}
