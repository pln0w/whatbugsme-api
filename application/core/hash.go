package core

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"

	"gopkg.in/mgo.v2/bson"
)

type HashData struct {
	ID       bson.ObjectId
	Org      bson.ObjectId
	Username string
	Topic    bson.ObjectId
}

// Hash is a function encoding any object to md5
func Hash(v interface{}) string {
	hasher := md5.New()
	hasher.Write([]byte(fmt.Sprintf("%#v", v)))
	return hex.EncodeToString(hasher.Sum(nil))
}
