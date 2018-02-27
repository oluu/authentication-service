package identity

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/oluu/authentication-service/util"
)

var envSessionRegion = util.GetRequiredStringEnv("AWS_SESSION_REGION")

// NewIdentityProvider returns a new cognito identity provider
func NewIdentityProvider() *cognitoidentityprovider.CognitoIdentityProvider {
	// initialize and configure aws session and cognito identity provider
	session := session.Must(session.NewSession())
	session.Config.Region = &envSessionRegion
	return cognitoidentityprovider.New(session)
}
