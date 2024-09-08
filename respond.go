package main

import (
	"encoding/json"
	"net/http"
)

func respondWithJSON(w http.ResponseWriter, code int, payload any) {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")

	res, err := json.Marshal(payload)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	_, err = w.Write(res)
	if err != nil {
		return
	}
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")

	res, err := json.Marshal(map[string]string{"error": message})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = w.Write(res)
	if err != nil {
		return
	}
}
