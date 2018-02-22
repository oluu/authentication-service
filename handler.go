package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go/aws/session"
	cip "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/shinypotato/authentication-service/util"
)

// RegisterHTTPHandlers initializes routes
func RegisterHTTPHandlers() {
	http.HandleFunc("/authentication/signup", routeHandler)
}

func routeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		handleSignUp(w, r)
	}
}

func handleSignUp(w http.ResponseWriter, r *http.Request) {
	session := session.Must(session.NewSession())
	creds, err := session.Config.Credentials.Get()
	if err != nil {
		log.Fatalln("error occured while attempting to get aws credentials: ", err)
	}
	log.Println("AccessKey: ", creds.AccessKeyID, " SecretAccessKey: ", creds.SecretAccessKey, " ProviderName: ", creds.ProviderName, " SessionToken: ", creds.SessionToken)
	// decode request
	data, err := util.ReadJSON(r, &SignUpRequest{})
	if err != nil {
		log.Fatalln("error occured while attempting to decode request: ", err)
	}
	signup := data.(*SignUpRequest)
	clientID := "6ci3qn66kp1pnjonkdmk19smet"
	clientIDPointer := &clientID
	secret := "1bqm3d6u7akafnaveid24atgcpfkbb7d34ec7a31n7ob7o28s6qf"
	secretHash := generateSecretHash(*signup.Username, clientID, secret)
	region := "us-east-1"
	session.Config.Region = &region
	var userAttributes []*cip.AttributeType
	emailStr := "email"
	userAttributes = append(userAttributes, &cip.AttributeType{
		Name:  &emailStr,
		Value: signup.Username,
	})
	// map signup request to aws cognito identity provider's type
	awsSignupRequest := &cip.SignUpInput{
		Username:       signup.Username,
		Password:       signup.Password,
		ClientId:       clientIDPointer,
		SecretHash:     &secretHash,
		UserAttributes: userAttributes,
	}
	log.Println(*awsSignupRequest.Username, *awsSignupRequest.Password, *awsSignupRequest.ClientId)
	// initialize cognito identity provider client
	cip := cip.New(session)
	awsSignupResponse, err := cip.SignUp(awsSignupRequest)
	if err != nil {
		log.Fatalln("error occured while calling signup with aws sdk: ", err)
	}
	strResponse := awsSignupResponse.String()
	w.Write([]byte(strResponse))
}

func generateSecretHash(username, clientID, secret string) string {
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(username + clientID))
	return base64.StdEncoding.EncodeToString(mac.Sum(nil))
}

// SignUpRequest is the request sent when the user signs up
type SignUpRequest struct {
	Username *string `json:"username"`
	Password *string `json:"password"`
}
