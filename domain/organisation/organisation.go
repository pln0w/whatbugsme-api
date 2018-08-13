package organisation

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Organisation struct {
	ID        bson.ObjectId `json:"id" bson:"id"`
	Name      string        `json:"name" bson:"name"`
	CreatedAt time.Time     `json:"created_at" bson:"created_at"`
}

const (
	C_ORGANISATION = "organisations"
)

type Organisations []Organisation

func CreateNewOrganisation(name string) Organisation {
	return Organisation{
		ID:        bson.NewObjectId(),
		Name:      name,
		CreatedAt: time.Now(),
	}
}
