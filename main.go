package main

import (
	"log"
	"net/http"

	"github.com/shinypotato/authentication-service/authentication"
	"github.com/shinypotato/authentication-service/identity"
)

func main() {
	// register authentication http handlers
	authentication.RegisterHTTPHandlers(identity.NewIdentityProvider())
	// cool logging to let us know the service is listening
	log.Println("/authentication/signup", "[POST]")
	http.ListenAndServe(":3000", nil)
}
