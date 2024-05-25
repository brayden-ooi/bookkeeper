-- name: CreateAccount :one
INSERT INTO accounts (name, type, owner_id) 
VALUES (?, ?, ?)
RETURNING *;

-- name: ListAccountsByUser :many
SELECT * FROM accounts WHERE owner_id = ?;