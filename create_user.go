package main

import (
	"blog-aggregator/internal/database"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"net/http"
	"time"
)

func (cfg *apiConfig) createUserHandler(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}

	type response struct {
		ID        uuid.UUID `json:"id"`
		CreatedAt string    `json:"created_at"`
		UpdatedAt string    `json:"updated_at"`
		Name      string    `json:"name"`
		Apikey    string    `json:"api_key"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Failed to decode request body: %v", err))
		return
	}

	user, err := cfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      params.Name,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Failed to create user: %v", err))
		return
	}

	respondWithJSON(w, http.StatusCreated, response{
		ID:        user.ID,
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
		UpdatedAt: user.UpdatedAt.Format(time.RFC3339),
		Name:      user.Name,
		Apikey:    user.Apikey,
	})
}
