package main

import (
	"blog-aggregator/internal/database"
	"net/http"
)

func (cfg *apiConfig) getAllUserFeedFollowsHandler(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollows, err := cfg.DB.GetFeedFollowsByUserId(r.Context(), user.ID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to get user feed follows")
		return
	}

	respondWithJSON(w, http.StatusOK, feedFollows)
}
