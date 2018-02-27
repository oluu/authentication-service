package authentication

import (
	"log"

	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/oluu/authentication-service/util"
)

// Signup creates the user in aws cognito and authenticates them
func Signup(request *Request, identityProvider *cognitoidentityprovider.CognitoIdentityProvider) error {
	// get data required for cognito request
	secretHash := util.GenerateSecretHash(*request.Username, envClientID, envClientSecret)
	userAttributes := []*cognitoidentityprovider.AttributeType{
		{Name: &emailStr, Value: request.Username},
	}
	// map our data to cognito expected request type
	cognitoRequest := &cognitoidentityprovider.SignUpInput{
		Username:       request.Username,
		Password:       request.Password,
		ClientId:       &envClientID,
		SecretHash:     &secretHash,
		UserAttributes: userAttributes,
	}
	// make request to cognito to create user
	_, err := identityProvider.SignUp(cognitoRequest)
	if err != nil {
		log.Println("failed to signup: ", err)
		return err
	}
	return nil
}
