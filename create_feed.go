package main

import (
	"blog-aggregator/internal/database"
	"encoding/json"
	"github.com/google/uuid"
	"net/http"
	"net/url"
	"time"
)

func (cfg *apiConfig) createFeedHandler(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}

	type FeedResponse struct {
		ID        uuid.UUID `json:"id"`
		CreatedAt string    `json:"created_at"`
		UpdatedAt string    `json:"updated_at"`
		Name      string    `json:"name"`
		URL       string    `json:"url"`
		UserID    uuid.UUID `json:"user_id"`
	}

	type FeedFollowResponse struct {
		ID     uuid.UUID `json:"id"`
		UserID uuid.UUID `json:"user_id"`
		FeedID uuid.UUID `json:"feed_id"`
	}

	type response struct {
		FeedResponse       FeedResponse
		FeedFollowResponse FeedFollowResponse
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Failed to decode request body")
		return
	}

	if params.Name == "" || params.URL == "" {
		respondWithError(w, http.StatusBadRequest, "Name and URL are required")
		return
	}

	if _, err = url.Parse(params.URL); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid URL")
		return
	}

	feed, err := cfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      params.Name,
		Url:       params.URL,
		UserID:    user.ID,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to create feed: "+err.Error())
		return
	}

	feedFollow, err := cfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:     uuid.New(),
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to create feed")
		return
	}

	respondWithJSON(w, http.StatusCreated, response{
		FeedResponse: FeedResponse{
			ID:        feed.ID,
			CreatedAt: feed.CreatedAt.Format(time.RFC3339),
			UpdatedAt: feed.UpdatedAt.Format(time.RFC3339),
			Name:      feed.Name,
			URL:       feed.Url,
			UserID:    feed.UserID,
		},
		FeedFollowResponse: FeedFollowResponse{
			ID:     feedFollow.ID,
			UserID: feedFollow.UserID,
			FeedID: feedFollow.FeedID,
		},
	})
}
