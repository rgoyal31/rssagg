package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func responsewithError(w http.ResponseWriter, code int, message string) {
	if code > 499 {
		log.Println("Responding with 5XX error", message)
	}
	type errorResponse struct {
		Admin string
	}

	respondWithJSON(w, code, errorResponse{Admin: message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload any) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Println("Failed to marshal JSON REPONSE: %v", payload)
		w.WriteHeader(500)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(data)
}
