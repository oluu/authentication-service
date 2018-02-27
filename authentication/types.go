package authentication

import "github.com/oluu/authentication-service/util"

const (
	authParamUsername   = "USERNAME"
	authParamPassword   = "PASSWORD"
	authParamSecretHash = "SECRET_HASH"
)

var (
	emailStr        = "email"
	authFlow        = "USER_PASSWORD_AUTH"
	envUserPoolID   = util.GetRequiredStringEnv("AWS_USER_POOL_ID")
	envClientID     = util.GetRequiredStringEnv("AWS_COGNITO_CLIENT_ID")
	envClientSecret = util.GetRequiredStringEnv("AWS_COGNITO_CLIENT_SECRET")
)

// Request is the object that is consumed from POST /authentication/signup
type Request struct {
	Username *string `json:"username"`
	Password *string `json:"password"`
}
