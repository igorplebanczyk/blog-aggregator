-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, name, apikey)
VALUES ($1, $2, $3, $4, encode(sha256(random()::text::bytea), 'hex'))
RETURNING id, created_at, updated_at, name, apikey;

-- name: GetUserByApiKey :one
SELECT id, created_at, updated_at, name, apikey
FROM users
WHERE apikey = $1;