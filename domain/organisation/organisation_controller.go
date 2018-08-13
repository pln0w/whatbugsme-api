package organisation

import (
	"errors"
	"net/http"
	"whatbugsme/application/core"
	"whatbugsme/infrastructure/db"
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
