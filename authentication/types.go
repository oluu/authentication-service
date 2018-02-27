package authentication

import "github.com/oluu/authentication-service/util"

const (
	authParamUsername = "USERNAME"
	authParamPassword = "PASSWORD"
	authSecretHash    = "SECRET_HASH"
)

var (
	emailStr        = "email"
	authFlow        = "USER_PASSWORD_AUTH"
	envUserPoolID   = util.GetRequiredStringEnv("AWS_USER_POOL_ID")
	envClientID     = util.GetRequiredStringEnv("AWS_COGNITO_CLIENT_ID")
	envClientSecret = util.GetRequiredStringEnv("AWS_COGNITO_CLIENT_SECRET")
)

// SignupRequest is the object that is consumed from POST /authentication/signup
type SignupRequest struct {
	Username *string `json:"username"`
	Password *string `json:"password"`
}

// SignupResponse is the response from POST /authenticate/signup
type SignupResponse struct {
}
