package router

import (
	"net/http"
	"whatbugsme/domain/organisation"
	"whatbugsme/domain/topic"
	"whatbugsme/domain/user"
	"whatbugsme/domain/vote"

	"whatbugsme/infrastructure/db"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
	Protected   bool
}

type Routes []Route

var (
	topicCtrl = topic.NewTopicController()
	userCtrl  = user.NewUserController()
	voteCtrl  = vote.NewVoteController()
	orgCtrl   = organisation.NewOrganisationController()
	routes    = Routes{

		// Auth
		Route{
			"RegisterOrganisation", "POST",
			"/organisation", orgCtrl.RegisterOrganisation, false,
		},
		Route{
			"SignUp", "POST",
			"/sign-up", userCtrl.SignUp, false,
		},
		Route{
			"Login", "POST",
			"/login", userCtrl.Login, false,
		},

		// Topics
		Route{
			"IndexTopics", "GET",
			"/organisation/{organisation}/topics", topicCtrl.Index, true,
		},
		Route{
			"CreateTopic", "POST",
			"/organisation/{organisation}/topics", topicCtrl.Create, true,
		},

		// Voting
		Route{
			"TopicVotes", "GET",
			"/organisation/{organisation}/topic/{topic}/votes", voteCtrl.GetTopicVotes, true,
		},
		Route{
			"VoteOnTopic", "POST",
			"/organisation/{organisation}/topic/{topic}/votes", voteCtrl.Create, true,
		},
	}
)

// CorsMiddleware adds proper headers
func CorsMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "X-Auth-Token")

		h.ServeHTTP(w, r)
	})
}

// AuthMiddleware implements setting headers
// it also takes token from request
// checks whether user with this token exists
// and pass request or return 401
func AuthMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "X-Auth-Token")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
		} else {

			// Get token from request headers
			at := r.Header.Get("X-Auth-Token")

			// Find user by this token
			user, _ := db.FindOneBy(user.C_USER, map[string]string{"token": at}, nil)
			if user == nil {
				w.WriteHeader(http.StatusUnauthorized)
			}

			// Next
			h.ServeHTTP(w, r)
		}
	})
}

// Router creates new Mux Router instance
// and registers handlers and middleware for each route
func Router() *mux.Router {

	// Create router object
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {

		var handler http.Handler

		// Set handler
		handler = CorsMiddleware(route.HandlerFunc)
		if route.Protected {
			handler = AuthMiddleware(handler)
		}

		// Add route
		router.
			Methods(route.Method, "OPTIONS").
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}
