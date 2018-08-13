package user

import (
	"errors"
	"log"
	"net/http"
	"runtime"
	"whatbugsme/application/core"
	o "whatbugsme/domain/organisation"

	"whatbugsme/infrastructure/db"

	"gopkg.in/mgo.v2/bson"
)

type UserController struct {
	core.Controller
}

// NewUserController returns pointer to controller
func NewUserController() *UserController {
	return &UserController{
		Controller: core.Controller{},
	}
}

// SignUp is a controller action,
// returns newly created user object in JSON if already does not exist in database
func (ctrl *UserController) SignUp(w http.ResponseWriter, r *http.Request) {

	params := map[string]string{
		"username":     r.FormValue("username"),
		"email":        r.FormValue("email"),
		"password":     r.FormValue("password"),
		"organisation": r.FormValue("organisation"),
	}

	vErr := ctrl.ValidEmptyParams(params, w)
	if vErr != nil {
		ctrl.HandleError(vErr, w, http.StatusUnprocessableEntity)
		return
	}

	ctrl.IsParamCorrectHex("organisation", params["organisation"], w)

	oID := bson.ObjectIdHex(params["organisation"])

	organisation, _ := db.FindOneBy(o.C_ORGANISATION, nil, map[string]bson.ObjectId{"id": oID})
	if organisation != nil {

		pass := GetPasswordHash(params["email"], params["password"])

		user, _ := db.FindOneBy(C_USER,
			map[string]string{"email": params["email"], "password": pass},
			map[string]bson.ObjectId{"organisation": oID},
		)
		if user != nil {
			ctrl.HandleError(errors.New("User exists"), w, http.StatusConflict)
			return
		}

		newUser := CreateNewUser(params["username"], params["email"], params["password"], oID)

		iErr := db.Insert(C_USER, newUser)
		if iErr != nil {
			_, fn, line, _ := runtime.Caller(1)
			log.Printf("[error] %s:%d %v", fn, line, iErr)
			ctrl.HandleError(iErr, w, http.StatusInternalServerError)
			return
		}

		ctrl.SendJSON(w, &newUser, http.StatusOK)

	} else {
		ctrl.HandleError(errors.New("Organisation not found"), w, http.StatusNotFound)
		return
	}
}

// Login is a controller action,
// returns JSON response with user token if found in database
func (ctrl *UserController) Login(w http.ResponseWriter, r *http.Request) {

	params := map[string]string{
		"email":        r.FormValue("email"),
		"password":     r.FormValue("password"),
		"organisation": r.FormValue("organisation"),
	}

	vErr := ctrl.ValidEmptyParams(params, w)
	if vErr != nil {
		ctrl.HandleError(vErr, w, http.StatusUnprocessableEntity)
		return
	}

	ctrl.IsParamCorrectHex("organisation", params["organisation"], w)

	oID := bson.ObjectIdHex(params["organisation"])

	organisation, _ := db.FindOneBy(o.C_ORGANISATION, nil, map[string]bson.ObjectId{"id": oID})
	if organisation != nil {

		pass := GetPasswordHash(params["email"], params["password"])

		fUser, _ := db.FindOneBy(C_USER,
			map[string]string{"email": params["email"], "password": pass},
			map[string]bson.ObjectId{"organisation": oID},
		)
		if fUser == nil {
			ctrl.HandleError(errors.New("Login failed"), w, http.StatusUnauthorized)
			return
		}

		var user User

		bsonBytes, _ := bson.Marshal(fUser)
		bson.Unmarshal(bsonBytes, &user)

		res := map[string]string{
			"status":   "success",
			"token":    user.Token,
			"username": user.Username,
		}
		ctrl.SendJSON(w, &res, http.StatusOK)

	} else {
		ctrl.HandleError(errors.New("Organisation not found"), w, http.StatusNotFound)
		return
	}
}
