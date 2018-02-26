package authentication

import (
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

// Authenticate handles authenticating users
func Authenticate(authParameters map[string]*string, identityProvider *cognitoidentityprovider.CognitoIdentityProvider) (*cognitoidentityprovider.InitiateAuthOutput, error) {
	authInput := &cognitoidentityprovider.InitiateAuthInput{
		AuthFlow:       &authFlow,
		ClientId:       &envClientID,
		AuthParameters: authParameters,
	}
	return identityProvider.InitiateAuth(authInput)
}
