package vote

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Vote struct {
	ID        bson.ObjectId `json:"id" bson:"id"`
	User      string        `json:"user_token" bson:"user_token"`
	Topic     bson.ObjectId `json:"topic" bson:"topic"`
	VoteType  VoteType      `json:"vote_type" bson:"vote_type"`
	CreatedAt time.Time     `json:"created_at" bson:"created_at"`
}

const (
	C_VOTE               = "votes"
	voteUnknown VoteType = 0
	voteDown    VoteType = -1
	voteUp      VoteType = 1
)

type Votes []Vote
type VoteType int8

// GetVoteType takes the number and return VoteType
func GetVoteType(vt int8) VoteType {

	// Check for invalid value
	if vt < int8(voteDown) || vt > int8(voteUp) {
		return voteUnknown
	}

	// Map allowed vote types
	voteTypes := map[int8]VoteType{
		-1: voteDown,
		1:  voteUp,
	}

	return voteTypes[vt]
}

func CreateNewVote(userToken string, topic bson.ObjectId, voteType int8) Vote {
	return Vote{
		ID:        bson.NewObjectId(),
		User:      userToken,
		Topic:     topic,
		VoteType:  GetVoteType(voteType),
		CreatedAt: time.Now(),
	}
}
