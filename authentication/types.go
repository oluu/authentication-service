package authentication

// SignupRequest is the object that is consumed from POST /authentication/signup
type SignupRequest struct {
	Username *string `json:"username"`
	Password *string `json:"password"`
}

// SignupResponse is the response from POST /authenticate/signup
type SignupResponse struct {
}
