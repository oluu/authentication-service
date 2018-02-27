package authentication

import (
	"net/http"

	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/oluu/authentication-service/util"
)

// RegisterHTTPHandlers initializes routes
func RegisterHTTPHandlers(identityProvider *cognitoidentityprovider.CognitoIdentityProvider) {
	http.HandleFunc("/authentication/signup", routeHandler(identityProvider, "signup"))
	http.HandleFunc("/authentication/login", routeHandler(identityProvider, "login"))
}

func routeHandler(identityProvider *cognitoidentityprovider.CognitoIdentityProvider, routeStr string) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case routeStr == "signup" && r.Method == http.MethodPost:
			postSignup(w, r, identityProvider)
		case routeStr == "login" && r.Method == http.MethodPut:
			putLogin(w, r, identityProvider)
		default:

		}
	})
}

func postSignup(w http.ResponseWriter, r *http.Request, identityProvider *cognitoidentityprovider.CognitoIdentityProvider) {
	signupRequest := new(Request)
	err := util.ReadJSON(r, signupRequest)
	if err != nil {
		util.WriteResponse(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	err = Signup(signupRequest, identityProvider)
	if err != nil {
		util.WriteResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func putLogin(w http.ResponseWriter, r *http.Request, identityProvider *cognitoidentityprovider.CognitoIdentityProvider) {
	loginRequest := new(Request)
	err := util.ReadJSON(r, loginRequest)
	if err != nil {
		util.WriteResponse(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	response, err := Login(loginRequest, identityProvider)
	if err != nil {
		util.WriteResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}
	util.WriteResponse(w, response, http.StatusOK)
}
