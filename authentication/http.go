package authentication

import (
	"net/http"

	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/shinypotato/authentication-service/util"
)

// RegisterHTTPHandlers initializes routes
func RegisterHTTPHandlers(identityProvider *cognitoidentityprovider.CognitoIdentityProvider) {
	http.HandleFunc("/authentication/signup", routeHandler(identityProvider))
}

func routeHandler(identityProvider *cognitoidentityprovider.CognitoIdentityProvider) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			postSignup(w, r, identityProvider)
		}
	})
}

func postSignup(w http.ResponseWriter, r *http.Request, identityProvider *cognitoidentityprovider.CognitoIdentityProvider) {
	signupRequest := new(SignupRequest)
	err := util.ReadJSON(r, signupRequest)
	if err != nil {
		util.WriteResponse(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	response, err := Signup(signupRequest, identityProvider)
	if err != nil {
		util.WriteResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}
	util.WriteResponse(w, response, http.StatusCreated)
}
