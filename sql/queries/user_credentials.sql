-- name: CreateUserCredentials :exec
INSERT INTO user_credentials (user_id, email, password, is_verified)
VALUES (?, ?, ?, 0);

-- name: GetUserCredentials :one
SELECT * FROM user_credentials
WHERE user_id = ? LIMIT 1;