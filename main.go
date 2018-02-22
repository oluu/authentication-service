package main

import (
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/shinypotato/authentication-service/authentication"
	"github.com/shinypotato/authentication-service/util"
)

var envSessionRegion = util.GetRequiredStringEnv("AWS_SESSION_REGION")

func main() {
	// initialize and configure aws session and cognito identity provider
	awsSession := session.Must(session.NewSession())
	awsSession.Config.Region = &envSessionRegion
	cognitoIdentityProvider := cognitoidentityprovider.New(awsSession)
	// register authentication http handlers
	authentication.RegisterHTTPHandlers(cognitoIdentityProvider)
	// cool logging to let us know the service is listening
	log.Println("/authentication/signup", "[POST]")
	http.ListenAndServe(":3000", nil)
}
