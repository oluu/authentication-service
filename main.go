package main

import (
	"log"
	"net/http"
)

func main() {
	RegisterHTTPHandlers()
	log.Println("Listening on port 3000")
	log.Println("/authentication/signup", "[GET]")
	http.ListenAndServe(":3000", nil)
}
