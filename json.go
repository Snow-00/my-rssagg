package main

import (
	"net/http"
	"log"
	"encoding/json"
)

func respondWithError(w http.ResponseWriter, code int, msg string) {
	// error code 500 above means there's a problem in the server side, error around 400 means problem from the client side
	if code > 499 {
		log.Println("Responding with 5xx error:", msg)
	}
	type errResponse struct {
		Error string `json:"error"`
	}

	respondWithJSON(w, code, errResponse{
		Error: msg,
	})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	dat, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to marshal JSON response: %v", payload)
		w.WriteHeader(500)
		return 
	}
	w.Header().Add("Content-Type", "applicatoin/json")
	w.WriteHeader(code)
	w.Write(dat)
}