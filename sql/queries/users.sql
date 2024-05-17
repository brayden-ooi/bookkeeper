-- name: CreateUser :one
INSERT INTO users (created_at, updated_at, name)
VALUES (?, ?, ?)
RETURNING *;
