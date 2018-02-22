package util

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

// ReadJSON takes in a request and decodes the request into the specified interface
func ReadJSON(r *http.Request, data interface{}) error {
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(data)
	return err
}

// WriteResponse handles writing status code and bytes
func WriteResponse(w http.ResponseWriter, response interface{}, statusCode int) {
	jsonResponse, _ := json.Marshal(response)
	w.WriteHeader(statusCode)
	w.Write(jsonResponse)
}

// GenerateSecretHash takes in username, clientID, and secret and hashes it
func GenerateSecretHash(username, clientID, secret string) string {
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(fmt.Sprintf("%s%s", username, clientID)))
	return base64.StdEncoding.EncodeToString(mac.Sum(nil))
}

// GetRequiredStringEnv will attempt to get an environment variable of type string, if it's nil it will panic
func GetRequiredStringEnv(key string) string {
	env := os.Getenv(key)
	if env == "" {
		log.Fatalf("%s is a required environment variable", key)
	}
	return env
}
