package main

import (
	"net/http"

	"github.com/aws/aws-sdk-go/service/cognitoidentity"
)

// RegisterHTTPHandlers initializes routes
func RegisterHTTPHandlers() {
	http.HandleFunc("/authenticate", routeHandler)
}

func routeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		handleGET(w, r)
	}
}

func handleGET(w http.ResponseWriter, r *http.Request) {
	cognitoIdentity := cognitoidentity.New()
}
