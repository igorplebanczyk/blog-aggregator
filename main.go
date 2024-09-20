package main

import (
	"blog-aggregator/internal/database"
	"blog-aggregator/internal/scraper"
	"context"
	"database/sql"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"net/http"
	"os"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	err := godotenv.Load()
	if err != nil {
		return
	}

	dbURL := os.Getenv("DB_CONN")
	db, err := sql.Open("postgres", dbURL)

	apiCfg := &apiConfig{
		DB: database.New(db),
	}

	go func() {
		user, err := apiCfg.DB.GetUserByApiKey(context.Background(), "a966658c2a39cfc117ea1ba500b63bcd6aaabd8522a817a59cb2a23a7ee83d89")
		if err != nil {
			return
		}

		scraper.FetchRSSFeedsWorker(apiCfg.DB, user, 10)
	}()

	port := os.Getenv("PORT")
	mux := http.NewServeMux()
	server := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	mux.HandleFunc("GET /v1/healthz", readinessHandler)
	mux.HandleFunc("GET /v1/err", errorHandler)

	mux.HandleFunc("POST /v1/users", apiCfg.createUserHandler)
	mux.HandleFunc("GET /v1/users", apiCfg.middlewareAuth(apiCfg.getUserHandler))

	mux.HandleFunc("POST /v1/feeds", apiCfg.middlewareAuth(apiCfg.createFeedHandler))
	mux.HandleFunc("GET /v1/feeds", apiCfg.getAllFeedsHandler)

	mux.HandleFunc("POST /v1/feed_follows", apiCfg.middlewareAuth(apiCfg.createFeedFollowHandler))
	mux.HandleFunc("DELETE /v1/feed_follows/{feed_id}", apiCfg.middlewareAuth(apiCfg.deleteFeedFollowHandler))
	mux.HandleFunc("GET /v1/feed_follows", apiCfg.middlewareAuth(apiCfg.getAllUserFeedFollowsHandler))

	err = server.ListenAndServe()
	if err != nil {
		return
	}
}
