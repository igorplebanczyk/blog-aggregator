package main

import (
	"blog-aggregator/internal/database"
	"github.com/google/uuid"
	"net/http"
)

func (cfg *apiConfig) deleteFeedFollowHandler(w http.ResponseWriter, r *http.Request, user database.User) {
	feedID, err := uuid.Parse(r.PathValue("feed_id"))
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid feed ID")
		return
	}

	feedFollowID, err := cfg.DB.GetFeedFollowByFeedAndUserId(r.Context(), database.GetFeedFollowByFeedAndUserIdParams{
		FeedID: feedID,
		UserID: user.ID,
	})
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Feed follow not found")
		return
	}

	_, err = cfg.DB.DeleteFeedFollow(r.Context(), feedFollowID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to delete feed follow")
		return
	}

	respondWithJSON(w, http.StatusNoContent, nil)
}
