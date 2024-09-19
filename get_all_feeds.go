package main

import (
	"net/http"
)

func (cfg *apiConfig) getAllFeedsHandler(w http.ResponseWriter, r *http.Request) {
	feeds, err := cfg.DB.GetAllFeeds(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to get feeds")
		return
	}

	respondWithJSON(w, http.StatusOK, feeds)
}
