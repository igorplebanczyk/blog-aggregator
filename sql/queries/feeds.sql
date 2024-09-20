-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id, last_fetched_at)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING id, created_at, updated_at, name, url, user_id;

-- name: GetAllFeeds :many
SELECT * FROM feeds;

-- name: GetNextFeedToFetch :many
SELECT * FROM feeds
WHERE user_id = $1
ORDER BY last_fetched_at ASC NULLS FIRST
LIMIT $2;

-- name: MarkFeedFetched :one
UPDATE feeds
SET last_fetched_at = $2 AND updated_at = $3
WHERE id = $1
RETURNING id, created_at, updated_at, name, url, user_id, last_fetched_at;