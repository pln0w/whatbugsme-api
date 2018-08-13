package core

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"gopkg.in/mgo.v2/bson"
)

type Controller struct{}

// SendJSON is a controller function,
// returns JSON response of any object
func (ctrl *Controller) SendJSON(w http.ResponseWriter, v interface{}, code int) {

	w.Header().Add("Content-Type", "application/json")

	b, err := json.Marshal(v)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, `{"error": "Internal server error"}`)
	} else {
		w.WriteHeader(code)
		io.WriteString(w, string(b))
	}
}

// HandleError is a controller function,
// returns error JSON message
func (ctrl *Controller) HandleError(err error, w http.ResponseWriter, status ...int) {

	msg := map[string]string{
		"status":  "fail",
		"message": err.Error(),
	}

	returnStatus := http.StatusInternalServerError
	if len(status) > 0 {
		returnStatus = status[0]
	}

	ctrl.SendJSON(w, &msg, returnStatus)

}

// ValidEmptyParams is a controller function,
// checks for empty parameters in request
func (ctrl *Controller) ValidEmptyParams(params map[string]string, w http.ResponseWriter) error {

	// TODO: change to parse all and give one message

	for param := range params {
		if params[param] == "" {
			return errors.New("Parameter " + param + " cannot be empty")
		}
	}

	return nil
}

// IsParamCorrectHex is a controller function,
// checks for proper hex format
func (ctrl *Controller) IsParamCorrectHex(key string, param string, w http.ResponseWriter) {

	if param != "" {
		if false == bson.IsObjectIdHex(param) {
			ctrl.HandleError(errors.New("Invalid "+key+" ID"), w, http.StatusUnprocessableEntity)
		}
	} else {
		ctrl.HandleError(errors.New("Parameter "+key+" canno be empty"), w, http.StatusUnprocessableEntity)
	}
}
