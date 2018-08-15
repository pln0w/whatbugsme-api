package organisation

import (
	"errors"
	"net/http"
	"whatbugsme/application/core"
	"whatbugsme/infrastructure/db"

	"gopkg.in/mgo.v2/bson"
)

type OrganisationController struct {
	core.Controller
}

// NewOrganisationController returns pointer to controller
func NewOrganisationController() *OrganisationController {
	return &OrganisationController{
		Controller: core.Controller{},
	}
}

// RegisterOrganisation is a controller action,
// returns JSON with newly created organisation
func (ctrl *OrganisationController) RegisterOrganisation(w http.ResponseWriter, r *http.Request) {

	params := map[string]string{
		"name": r.FormValue("name"),
	}

	vErr := ctrl.ValidEmptyParams(params, w)
	if vErr != nil {
		ctrl.HandleError(vErr, w, http.StatusUnprocessableEntity)
		return
	}

	n, eErr := db.ExistBy(C_ORGANISATION, params)
	if eErr != nil {
		ctrl.HandleError(eErr, w, http.StatusInternalServerError)
		return
	}

	if n > 0 {
		ctrl.HandleError(errors.New("Organisation already exist"), w, http.StatusConflict)
		return
	}

	organisation := CreateNewOrganisation(params["name"])

	iErr := db.Insert(C_ORGANISATION, organisation)
	if iErr != nil {
		ctrl.HandleError(iErr, w, http.StatusInternalServerError)
		return
	}

	ctrl.SendJSON(w, &organisation, http.StatusOK)
}

// SearchOrganisation is a controller action,
// returns JSON with found organisation
func (ctrl *OrganisationController) SearchOrganisation(w http.ResponseWriter, r *http.Request) {

	params := map[string]string{
		"name": r.FormValue("name"),
	}

	vErr := ctrl.ValidEmptyParams(params, w)
	if vErr != nil {
		ctrl.HandleError(vErr, w, http.StatusUnprocessableEntity)
		return
	}

	fOrg, _ := db.FindOneBy(C_ORGANISATION, map[string]string{"name": params["name"]}, nil)
	if fOrg == nil {
		ctrl.HandleError(errors.New("Organisation not found"), w, http.StatusNotFound)
		return
	}
	// Proper casting bsons to structs
	var organisation Organisation
	bsonBytes, _ := bson.Marshal(fOrg)
	bson.Unmarshal(bsonBytes, &organisation)

	ctrl.SendJSON(w, &organisation, http.StatusOK)
}
