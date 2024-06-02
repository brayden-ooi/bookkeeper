-- name: CreateUser :one
INSERT INTO users (created_at, updated_at, name)
VALUES (unixepoch(), unixepoch(), ?)
RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE id = ? LIMIT 1;

-- name: ResetCounterByUser :exec
UPDATE users SET transaction_counter = 0 WHERE id = ?;