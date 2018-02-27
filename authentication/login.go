package authentication

import (
	"log"

	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/oluu/authentication-service/util"
)

// Login authenticates the user in aws cognito
func Login(request *Request, identityProvider *cognitoidentityprovider.CognitoIdentityProvider) (*cognitoidentityprovider.AuthenticationResultType, error) {
	// generate secret hash
	secretHash := util.GenerateSecretHash(*request.Username, envClientID, envClientSecret)
	// setup auth parameters
	authParameters := make(map[string]*string)
	authParameters[authParamUsername] = request.Username
	authParameters[authParamPassword] = request.Password
	authParameters[authParamSecretHash] = &secretHash
	// call Authenticate and serve result
	authResponse, err := Authenticate(authParameters, identityProvider)
	if err != nil {
		log.Println("failed to authenticate: ", err)
		return nil, err
	}
	return authResponse.AuthenticationResult, nil
}
