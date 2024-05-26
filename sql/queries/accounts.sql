-- name: CreateAccount :one
INSERT INTO accounts (id, name, type, owner_id) 
VALUES (?, ?, ?, ?)
RETURNING *;

-- name: ListAccountsByUser :many
SELECT * FROM accounts WHERE owner_id = ?;

-- name: GetAccountByUserAndID :one
SELECT * FROM accounts WHERE owner_id = ? AND id = ?;