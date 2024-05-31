-- name: CreateAccount :one
INSERT INTO accounts (id, name, type, user_id) 
VALUES (?, ?, ?, ?)
RETURNING *;

-- name: ListAccountsByUser :many
SELECT * FROM accounts WHERE user_id = ?;

-- name: GetAccountByUserAndID :one
SELECT * FROM accounts WHERE user_id = ? AND id = ?;

-- name: DeleteAccountByUserAndID :exec
DELETE FROM accounts WHERE user_id = ? AND id = ?;