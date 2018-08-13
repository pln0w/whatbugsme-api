package user

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"time"

	"gopkg.in/mgo.v2/bson"
)

type User struct {
	ID           bson.ObjectId `json:"id" bson:"_id"`
	Username     string        `json:"username" bson:"username"`
	Email        string        `json:"email" bson:"email"`
	Password     string        `json:"password" bson:"password"`
	Organisation bson.ObjectId `json:"organisation" bson:"organisation"`
	Token        string        `json:"token" bson:"token"`
	CreatedAt    time.Time     `json:"created_at" bson:"created_at"`
}

const (
	C_USER = "users"
)

type Users []User

func GetPasswordHash(email string, pass string) string {
	hasher := md5.New()
	hasher.Write([]byte(email + pass))
	return hex.EncodeToString(hasher.Sum(nil))
}

func CreateNewUser(
	username string,
	email string,
	password string,
	organisation bson.ObjectId,
) User {

	newUser := User{
		ID:           bson.NewObjectId(),
		Username:     username,
		Email:        email,
		Password:     GetPasswordHash(email, password),
		Organisation: organisation,
		CreatedAt:    time.Now(),
	}

	// Update user token based on its properties
	tokenHasher := md5.New()
	tokenHasher.Write([]byte(fmt.Sprintf("%#v", newUser)))

	newUser.Token = hex.EncodeToString(tokenHasher.Sum(nil))

	return newUser
}
