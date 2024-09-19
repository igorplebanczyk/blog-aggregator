package main

import (
	"blog-aggregator/internal/database"
	"net/http"
)

func (cfg *apiConfig) getUserHandler(w http.ResponseWriter, r *http.Request, user database.User) {
	if user == (database.User{}) {
		respondWithError(w, http.StatusNotFound, "User not found")
		return
	}

	respondWithJSON(w, http.StatusOK, user)
}
