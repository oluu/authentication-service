package authentication

import (
	"log"

	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/shinypotato/authentication-service/util"
)

// Signup creates the user in aws cognito and authenticates them
func Signup(signupRequest *SignupRequest, identityProvider *cognitoidentityprovider.CognitoIdentityProvider) (*cognitoidentityprovider.InitiateAuthOutput, error) {
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
	// make request to cognito to create user
	_, err := identityProvider.SignUp(cognitoSignupRequest)
	if err != nil {
		log.Println("failed to signup: ", err)
		return nil, err
	}
	// make request to cognito to go ahead and verify user
	adminConfirmRequest := &cognitoidentityprovider.AdminConfirmSignUpInput{
		UserPoolId: &envUserPoolID,
		Username:   signupRequest.Username,
	}
	_, err = identityProvider.AdminConfirmSignUp(adminConfirmRequest)
	if err != nil {
		log.Println("failed to confirm signup: ", err)
		return nil, err
	}
	// initialize authentication parameters
	authParameters := make(map[string]*string)
	authParameters[authParamUsername] = signupRequest.Username
	authParameters[authParamPassword] = signupRequest.Password
	authParameters[authSecretHash] = &secretHash
	// attempt to authenticate user after signup
	authResponse, err := Authenticate(authParameters, identityProvider)
	if err != nil {
		log.Println("failed to authenticate: ", err)
		return nil, err
	}
	// return authentication response
	return authResponse, nil
}
