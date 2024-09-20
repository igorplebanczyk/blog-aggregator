-- name: CreatePost :one
INSERT INTO posts (id, created_at, updated_at, title, url, feed_id)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id, created_at, updated_at, title, url, feed_id;

-- name: GetPostsByUser :many
SELECT * FROM posts
JOIN feeds ON posts.feed_id = feeds.id
WHERE feeds.user_id = $1
ORDER BY posts.created_at DESC;

