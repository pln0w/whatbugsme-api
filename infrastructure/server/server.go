package server

import (
	"fmt"
	"net/http"
	"whatbugsme/infrastructure/env"
	"whatbugsme/infrastructure/router"

	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func init() {

	// Load variables, default from .env
	if errEnv := godotenv.Load(); errEnv != nil {
		panic("Error loading .env file")
	}
}

// Run function in Server package implements HTTP listener
// based on environment variables
func Run() {

	// Listen on the given port on localhost
	router := router.Router()

	// Setting CORS options
	c := cors.New(cors.Options{
		AllowedHeaders:     []string{"X-Auth-Token, x-auth-token"},
		AllowedOrigins:     []string{"*"},
		AllowedMethods:     []string{"POST, GET, OPTIONS"},
		OptionsPassthrough: true,
	})

	handler := c.Handler(router)

	// Welcome message print
	fmt.Println("Server start listening on port " + env.Get().APIPort)

	// Run server at given port
	if err := http.ListenAndServe(":"+env.Get().APIPort, handler); err != nil {
		panic("Could not run server: " + err.Error())
	}
}
