-- name: CreateFeedFollow :one
INSERT INTO feed_follows (id, user_id, feed_id)
VALUES ($1, $2, $3)
RETURNING id, user_id, feed_id;

-- name: DeleteFeedFollow :one
DELETE FROM feed_follows
WHERE id = $1
RETURNING id, user_id, feed_id;

-- name: GetFeedFollowByFeedAndUserId :one
SELECT id
FROM feed_follows
WHERE feed_id = $1 AND user_id = $2;

-- name: GetFeedFollowsByUserId :many
SELECT id, user_id, feed_id
FROM feed_follows
WHERE user_id = $1;