package topic

import (
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

	ctrl.IsParamCorrectHex("organisation", params["organisation"], w)

	oID := bson.ObjectIdHex(params["organisation"])

	organisation, _ := db.FindOneBy(o.C_ORGANISATION, nil, map[string]bson.ObjectId{"_id": oID})
	if organisation != nil {

		topics, tErr := db.FindAllBy(C_TOPIC, nil, map[string]bson.ObjectId{"organisation": oID})
		if tErr != nil {
			_, fn, line, _ := runtime.Caller(1)
			log.Printf("[error] %s:%d %v", fn, line, tErr)
			ctrl.HandleError(tErr, w, http.StatusInternalServerError)
			return
		}
		if topics == nil {
			ctrl.SendJSON(w, &[]Topic{}, http.StatusOK)
		}

		ctrl.SendJSON(w, &topics, http.StatusOK)

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
		"organisation": mux.Vars(r)["organisation"],
	}

	vErr := ctrl.ValidEmptyParams(params, w)
	if vErr != nil {
		ctrl.HandleError(vErr, w, http.StatusUnprocessableEntity)
		return
	}

	ctrl.IsParamCorrectHex("organisation", params["organisation"], w)

	oID := bson.ObjectIdHex(params["organisation"])

	organisation, _ := db.FindOneBy(o.C_ORGANISATION, nil, map[string]bson.ObjectId{"_id": oID})
	if organisation != nil {

		topic := CreateNewTopic(params["content"], oID)

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
