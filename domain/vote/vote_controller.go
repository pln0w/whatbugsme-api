package vote

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"runtime"
	"strconv"
	"whatbugsme/application/core"
	t "whatbugsme/domain/topic"
	u "whatbugsme/domain/user"
	"whatbugsme/infrastructure/db"

	"gopkg.in/mgo.v2/bson"

	"github.com/gorilla/mux"
)

// VoteController structure
type VoteController struct {
	core.Controller
}

// NewVoteController returns pointer to controller
func NewVoteController() *VoteController {
	return &VoteController{
		Controller: core.Controller{},
	}
}

// GetTopicVotes is a controller action,
// returns JSON list of votes for given topic
func (ctrl *VoteController) GetTopicVotes(w http.ResponseWriter, r *http.Request) {

	params := map[string]string{
		"topic":        mux.Vars(r)["topic"],
		"organisation": mux.Vars(r)["organisation"],
	}

	ctrl.ValidEmptyParams(params, w)

	ctrl.IsParamCorrectHex("topic", params["topic"], w)
	ctrl.IsParamCorrectHex("organisation", params["organisation"], w)

	tID, oID := bson.ObjectIdHex(params["topic"]), bson.ObjectIdHex(params["organisation"])

	topic, _ := db.FindOneBy(t.C_TOPIC, nil, map[string]bson.ObjectId{"id": tID, "organisation": oID})
	if topic == nil {
		ctrl.HandleError(errors.New("Topic not found"), w, http.StatusNotFound)
		return
	}

	votes, faErr := db.FindAllBy(C_VOTE, nil, map[string]bson.ObjectId{"topic": tID}, "")
	if faErr != nil {
		_, fn, line, _ := runtime.Caller(1)
		log.Printf("[error] %s:%d %v", fn, line, faErr)
		ctrl.HandleError(faErr, w, http.StatusInternalServerError)
		return
	}

	var res Votes

	data, _ := json.Marshal(votes)
	_ = json.Unmarshal(data, &res)

	ctrl.SendJSON(w, &res, http.StatusOK)
}

// Create is a controller action,
// returns JSON with newly created vote for given topic
func (ctrl *VoteController) Create(w http.ResponseWriter, r *http.Request) {

	params := map[string]string{
		"topic":        mux.Vars(r)["topic"],
		"organisation": mux.Vars(r)["organisation"],
		"voteType":     r.FormValue("vote_type"),
	}

	vErr := ctrl.ValidEmptyParams(params, w)
	if vErr != nil {
		ctrl.HandleError(vErr, w, http.StatusUnprocessableEntity)
		return
	}

	ctrl.IsParamCorrectHex("organisation", params["organisation"], w)
	ctrl.IsParamCorrectHex("topic", params["topic"], w)

	voteType, castErr := strconv.Atoi(params["voteType"])
	if castErr != nil {
		ctrl.HandleError(errors.New("Invalid vote type"), w, http.StatusUnprocessableEntity)
		return
	}

	tID, oID := bson.ObjectIdHex(params["topic"]), bson.ObjectIdHex(params["organisation"])

	fTopic, _ := db.FindOneBy(t.C_TOPIC, nil, map[string]bson.ObjectId{"id": tID, "organisation": oID})
	if fTopic == nil {
		ctrl.HandleError(errors.New("Topic not found"), w, http.StatusNotFound)
		return
	}

	// Retrieve user auth token from request
	token := r.Header.Get("X-Auth-Token")

	fUser, _ := db.FindOneBy(u.C_USER, map[string]string{"token": token}, nil)
	if fUser == nil {
		ctrl.HandleError(errors.New("User not found"), w, http.StatusNotFound)
		return
	}

	// Proper casting bsons to structs
	var user u.User
	uBsonBytes, _ := bson.Marshal(fUser)
	bson.Unmarshal(uBsonBytes, &user)

	var topic t.Topic
	tBsonBytes, _ := bson.Marshal(fTopic)
	bson.Unmarshal(tBsonBytes, &topic)

	// Encode vote association with user via special token
	hashToken := core.Hash(&core.HashData{
		ID:       user.ID,
		Org:      user.Organisation,
		Username: user.Username,
		Topic:    topic.ID,
	})

	existingVote, _ := db.FindOneBy(C_VOTE, map[string]string{"user_token": hashToken}, nil)

	if existingVote != nil {
		ctrl.HandleError(errors.New("User already voted"), w, http.StatusConflict)
		return
	}

	vote := CreateNewVote(hashToken, topic.ID, int8(voteType))

	iErr := db.Insert(C_VOTE, vote)
	if iErr != nil {
		ctrl.HandleError(iErr, w, http.StatusInternalServerError)
		return
	}

	field := "votes_up"
	if voteType == -1 {
		field = "votes_down"
	}

	incErr := db.IncrementFieldWhere(t.C_TOPIC, field, 1, nil, map[string]bson.ObjectId{"id": tID})
	if incErr != nil {
		_, fn, line, _ := runtime.Caller(1)
		log.Printf("[error] %s:%d %v", fn, line, incErr)
		ctrl.HandleError(incErr, w, http.StatusInternalServerError)
		return
	}

	updatedTopic, _ := db.FindOneBy(t.C_TOPIC, nil, map[string]bson.ObjectId{"id": topic.ID})
	if updatedTopic == nil {
		ctrl.HandleError(errors.New("Topic not found"), w, http.StatusNotFound)
		return
	}

	var res t.Topic

	data, _ := json.Marshal(updatedTopic)
	_ = json.Unmarshal(data, &res)

	ctrl.SendJSON(w, &res, http.StatusOK)
}
