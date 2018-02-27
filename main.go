package main

import (
	"log"
	"net/http"

	"github.com/oluu/authentication-service/authentication"
	"github.com/oluu/authentication-service/identity"
)

func main() {
	// register authentication http handlers
	authentication.RegisterHTTPHandlers(identity.NewIdentityProvider())
	// cool logging to let us know the service is listening
	log.Println("/authentication/signup", "[POST]")
	log.Println("/authentication/login", "[PUT]")
	http.ListenAndServe(":3000", nil)
}
