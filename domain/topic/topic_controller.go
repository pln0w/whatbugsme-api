package topic

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"runtime"
	"whatbugsme/application/core"
	o "whatbugsme/domain/organisation"
	"whatbugsme/infrastructure/db"

	"gopkg.in/mgo.v2/bson"

	"github.com/gorilla/mux"
)

type TopicController struct {
	core.Controller
}

// NewTopicController returns pointer to controller
func NewTopicController() *TopicController {
	return &TopicController{
		Controller: core.Controller{},
	}
}

// Index is a controller action,
// returns all topics for given organisation
func (ctrl *TopicController) Index(w http.ResponseWriter, r *http.Request) {

	params := map[string]string{
		"organisation": mux.Vars(r)["organisation"],
	}

	vErr := ctrl.ValidEmptyParams(params, w)
	if vErr != nil {
		ctrl.HandleError(vErr, w, http.StatusUnprocessableEntity)
		return
	}

	oE := ctrl.IsParamCorrectHex("organisation", params["organisation"])
	if oE != nil {
		ctrl.HandleError(oE, w, http.StatusUnprocessableEntity)
		return
	}

	oID := bson.ObjectIdHex(params["organisation"])

	organisation, _ := db.FindOneBy(o.C_ORGANISATION, nil, map[string]bson.ObjectId{"id": oID})
	if organisation != nil {

		topics, tErr := db.FindAllBy(C_TOPIC, nil, map[string]bson.ObjectId{"organisation": oID}, "-created_at")

		var data []byte
		var res Topics

		if topics == nil || tErr != nil {
			data, _ = json.Marshal(make([]Topic, 0))
		} else {
			data, _ = json.Marshal(topics)
		}

		_ = json.Unmarshal(data, &res)

		ctrl.SendJSON(w, &res, http.StatusOK)

	} else {
		ctrl.HandleError(errors.New("Organisation not found"), w, http.StatusNotFound)
		return
	}
}

// Create is a controller action,
// returns JSON with newly created topic
func (ctrl *TopicController) Create(w http.ResponseWriter, r *http.Request) {

	params := map[string]string{
		"content":      r.FormValue("content"),
		"guid":         r.FormValue("guid"),
		"organisation": mux.Vars(r)["organisation"],
	}

	vErr := ctrl.ValidEmptyParams(params, w)
	if vErr != nil {
		ctrl.HandleError(vErr, w, http.StatusUnprocessableEntity)
		return
	}

	oE := ctrl.IsParamCorrectHex("organisation", params["organisation"])
	if oE != nil {
		ctrl.HandleError(oE, w, http.StatusUnprocessableEntity)
		return
	}

	oID := bson.ObjectIdHex(params["organisation"])

	organisation, _ := db.FindOneBy(o.C_ORGANISATION, nil, map[string]bson.ObjectId{"id": oID})
	if organisation != nil {

		tGuidExst, _ := db.FindOneBy(C_TOPIC, map[string]string{"guid": params["guid"]}, nil)
		if tGuidExst != nil {
			ctrl.HandleError(errors.New("Topic with given GUID already exists"), w, http.StatusConflict)
			return
		}

		topic := CreateNewTopic(params["content"], oID, params["guid"])

		iErr := db.Insert(C_TOPIC, topic)
		if iErr != nil {
			_, fn, line, _ := runtime.Caller(1)
			log.Printf("[error] %s:%d %v", fn, line, iErr)
			ctrl.HandleError(iErr, w, http.StatusInternalServerError)
			return
		}

		ctrl.SendJSON(w, &topic, http.StatusOK)

	} else {
		ctrl.HandleError(errors.New("Organisation not found"), w, http.StatusNotFound)
		return
	}
}
