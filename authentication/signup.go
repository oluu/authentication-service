package authentication

import (
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/shinypotato/authentication-service/util"
)

var (
	emailStr        = "email"
	envClientID     = util.GetRequiredStringEnv("AWS_COGNITO_CLIENT_ID")
	envClientSecret = util.GetRequiredStringEnv("AWS_COGNITO_CLIENT_SECRET")
)

// Signup creates the user in aws cognito and authenticates them
func Signup(signupRequest *SignupRequest, identityProvider *cognitoidentityprovider.CognitoIdentityProvider) (*cognitoidentityprovider.SignUpOutput, error) {
	// get data required for cognito request
	secretHash := util.GenerateSecretHash(*signupRequest.Username, envClientID, envClientSecret)
	userAttributes := []*cognitoidentityprovider.AttributeType{
		{Name: &emailStr, Value: signupRequest.Username},
	}
	// map our data to cognito expected request type
	cognitoSignupRequest := &cognitoidentityprovider.SignUpInput{
		Username:       signupRequest.Username,
		Password:       signupRequest.Password,
		ClientId:       &envClientID,
		SecretHash:     &secretHash,
		UserAttributes: userAttributes,
	}
	// make request to cognito
	return identityProvider.SignUp(cognitoSignupRequest)
}
