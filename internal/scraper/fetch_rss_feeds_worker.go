package scraper

import (
	"blog-aggregator/internal/database"
	"context"
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"net/http"
	"sync"
	"time"
)

func FetchRSSFeedsWorker(db *database.Queries, feedNum int32) {
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	client := http.Client{
		Timeout: time.Second * 15,
	}

	for {
		select {
		case <-ticker.C:
			feeds, err := db.GetNextFeedToFetch(context.Background(), feedNum)
			if err != nil {
				fmt.Printf("failed to fetch feeds: %v", err)
				continue
			}

			processFeeds(db, feeds, client)
		}
	}
}

func processFeeds(db *database.Queries, feeds []database.Feed, client http.Client) {
	var wg sync.WaitGroup
	for _, feed := range feeds {
		wg.Add(1)
		go func(feed database.Feed) {
			defer wg.Done()

			rss, err := FetchFeedData(client, feed.Url)
			if err != nil {
				fmt.Printf("failed to fetch feed data: %v", err)
				return
			}

			for _, item := range rss.Channel.Items {
				publishedAt, err := time.Parse(time.RFC1123Z, item.PubDate)
				if err != nil {
					fmt.Printf("failed to parse published at: %v", err)
					continue
				}

				_, err = db.CreatePost(context.Background(), database.CreatePostParams{
					ID:          uuid.New(),
					CreatedAt:   time.Now(),
					UpdatedAt:   time.Now(),
					Title:       item.Title,
					Url:         item.Link,
					Description: item.Description,
					PublishedAt: publishedAt,
					FeedID:      feed.ID,
				})
				if err != nil {
					return
				}
				fmt.Println(item.Title)
			}

			_, err = db.MarkFeedFetched(context.Background(), database.MarkFeedFetchedParams{
				ID: feed.ID,
				LastFetchedAt: sql.NullTime{
					Time: time.Now(),
				},
				UpdatedAt: time.Now(),
			})
			if err != nil {
				return
			}
		}(feed)
	}
	wg.Wait()
}
