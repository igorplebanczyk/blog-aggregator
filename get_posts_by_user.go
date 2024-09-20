package main

import (
	"blog-aggregator/internal/database"
	"net/http"
	"strconv"
)

func (cfg *apiConfig) getPostsByUserHandler(w http.ResponseWriter, r *http.Request, user database.User) {
	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		http.Error(w, "invalid limit", http.StatusBadRequest)
		return
	}
	if limit == 0 {
		limit = 10
	}

	posts, err := cfg.DB.GetPostsByUser(r.Context(), database.GetPostsByUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	})
	if err != nil {
		http.Error(w, "failed to get posts", http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusOK, posts)
}
