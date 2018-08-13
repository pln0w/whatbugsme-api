package main

import (
	"whatbugsme/infrastructure/env"
	"whatbugsme/infrastructure/server"

	"fmt"
)

// Main function
func main() {
	// Init server application
	server.Run()

	// Prints the version and the address of our api to the console
	fmt.Println(env.Get().APITitle + " version " + env.Get().APIVersion + " [" + env.Get().Environment + "]")
}
