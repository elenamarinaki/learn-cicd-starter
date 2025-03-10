package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Printf("Responding with 5XX error: %s", msg)
	}
	type errorResponse struct {
		Error string `json:"error"`
	}
	respondWithJSON(w, code, errorResponse{
		Error: msg,
	})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	dat, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshalling JSON: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		_, writeErr := w.Write([]byte(`{"error": "Internal Server Error"}`))
		if writeErr != nil {
			log.Printf("Error writing error response: %s", writeErr)
		}
		return
	}

	w.WriteHeader(code)
	if _, err := w.Write(dat); err != nil {
		log.Printf("Error writing JSON response: %s", err)
	}
}
