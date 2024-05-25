package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondeWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Println("Reading with 5xx error :", msg)
	}
	type errResponse struct {
		Error string `json:"error"`
	}
	respondeWithJSON(w, code, errResponse{
		Error: msg,
	})
}

func respondeWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	dat, err := json.Marshal(payload)
	if err != nil {
		log.Printf("failed to marshell JSON response: %v", payload)
		w.WriteHeader(500)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(dat)
}