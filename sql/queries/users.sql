-- name: CreateUser :one
INSERT INTO users (created_at, updated_at, name)
VALUES (?, ?, ?)
RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE id = ? LIMIT 1;